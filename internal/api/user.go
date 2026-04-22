package api

import (
	"net/http"
	"simplebank/internal/common"
	db "simplebank/internal/db/sqlc"
	"simplebank/internal/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		FullName:          user.FullName,
		Email:             user.Email,
		CreatedAt:         user.CreatedAt.Time,
	}
}

// Http Post Create User
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		ID: uuid.New(),
		FullName:       req.FullName,
		Email:          req.Email,
		PasswordHash: hashedPassword,	
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//rsp := newUserResponse(user)
	//ctx.JSON(http.StatusOK, rsp)
	data := newUserResponse(user)
	common.Success(ctx, http.StatusOK, "Create user success", data)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}


type getListUsersRequest struct {
    PageId     int    `form:"page_id"`
    PageSize int    `form:"page_size"`
}


func (server *Server) listUsers(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	common.Success(ctx, http.StatusOK, "Get user success", users)
}