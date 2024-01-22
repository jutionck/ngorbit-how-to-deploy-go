package common

import (
	"net/http"

	"enigmacamp.com/blog-apps/shared/shared_model"
	"github.com/gin-gonic/gin"
)

func SendCreateResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, &shared_model.SingleResponse{
		Status: shared_model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, &shared_model.SingleResponse{
		Status: shared_model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendPagedResponse(ctx *gin.Context, data []interface{}, paging shared_model.Paging, message string) {
	ctx.JSON(http.StatusOK, &shared_model.PagedResponse{
		Status: shared_model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data:   data,
		Paging: paging,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &shared_model.Status{
		Code:    code,
		Message: message,
	})
}
