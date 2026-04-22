package api

import (
	"crypto/rand"
	"math/big"
	"net/http"

	db "simplebank/internal/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createAccountRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type addAccountBalanceRequest struct {
	AccountID uuid.UUID `json:"account_id" binding:"required,uuid"`
	Amount int64 `json:"amount" binding:"required"`
}


func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userId, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid user id",
		})
		return
	}

	arg := db.CreateAccountParams{
		ID : uuid.New(),
		UserID: userId,
		AccountNumber : random10DigitString(),
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {	
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) addAccountBalance(ctx *gin.Context) {
	var req addAccountBalanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddAccountBalanceParams{
		ID : req.AccountID,
		Amount: req.Amount,
	}

	account, err := server.store.AddAccountBalance(ctx, arg)
	if err != nil {	
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func random10DigitString() (string) {
    const digits = "0123456789"
    result := make([]byte, 10)

    for i := range result {
        n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
        if err != nil {
            return ""
        }
        result[i] = digits[n.Int64()]
    }

    return string(result)
}
