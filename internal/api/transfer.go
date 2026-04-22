package api

import (
	"net/http"
	db "simplebank/internal/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type transferRequest struct {
	FromAccountID uuid.UUID  `json:"from_account_id" binding:"required,uuid"`
	ToAccountID   uuid.UUID  `json:"to_account_id" binding:"required,uuid"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, valid := server.validAccount(ctx, req.FromAccountID)
	if !valid {
		return
	}

	_, valid = server.validAccount(ctx, req.ToAccountID)
	if !valid {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}



func (server *Server) validAccount(ctx *gin.Context, accountID uuid.UUID) (db.Account, bool) {
	
	account, err := server.store.GetAccountById(ctx, accountID)
	if err != nil {	
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}
	return account, true
}