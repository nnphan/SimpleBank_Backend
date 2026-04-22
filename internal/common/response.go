package common

import "github.com/gin-gonic/gin"

type APIResponse struct {
    Code    int         `json:"code"`    // HTTP status code
    Message string      `json:"message"` // Message cho client
    Data    interface{} `json:"data"`    // Kết quả trả về
}

func Success(
    ctx *gin.Context,
    statusCode int,
    message string,
    data interface{},
) {
    ctx.JSON(statusCode, APIResponse{
        Code:    statusCode,
        Message: message,
        Data:    data,
    })
}

func Error(
    ctx *gin.Context,
    statusCode int,
    message string,
) {
    ctx.JSON(statusCode, APIResponse{
        Code:    statusCode,
        Message: message,
        Data:    nil,
    })
}

