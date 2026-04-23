package response

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Status  int         `json:"status"`         // HTTP status code
	Message string      `json:"message"`        // Message for http status code
	Data    interface{} `json:"data,omitempty"` // omitempty: if data is nil, it will be omitted in the response
}

func SuccessResponse(context *gin.Context, code int, data interface{}){
	context.JSON(code, ResponseData{
		Status:  code,
		Message: msg[code],
		Data:    data,
	})	
}

func ErrorResponse(context *gin.Context, code int, message string){
	context.JSON(code, ResponseData{
		Status:  code,
		Message: message,
		Data:    nil,
	})	
}