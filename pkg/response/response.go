package response

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	Code     int         `json:"code,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message,omitempty"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
	Total    int64       `json:"total,omitempty"`
}

func Pagination(ctx *gin.Context, code int, data interface{}, page, pageSize int, total int64) {
	ctx.JSON(code, &JsonResponse{
		Code:     code,
		Data:     data,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})
}

func Success(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, &JsonResponse{
		Code: code,
		Data: data,
	})
}

func Error(ctx *gin.Context, code int, err interface{}) {
	var msg string
	switch err.(type) {
	case error:
		msg = err.(error).Error()
	case string:
		msg = err.(string)
	default:
		msg = "internal error."
	}
	ctx.AbortWithStatusJSON(code, &JsonResponse{
		Code:    code,
		Message: msg,
	})
}
