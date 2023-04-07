package helper

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespBadRequest(ctx *gin.Context, msg string) {
	RestPonse(ctx,
		http.StatusBadRequest,
		gin.H{
			"message": msg,
		})
}

func RespStatusOk(ctx *gin.Context, data interface{}) {
	RestPonse(ctx,
		http.StatusOK,
		gin.H{
			"data": data,
		})
}

func RespStatusOkWithMessage(ctx *gin.Context, msg string) {
	RestPonse(ctx,
		http.StatusOK,
		gin.H{
			"message": msg,
		})
}

func RespNotFound(ctx *gin.Context, msg string) {
	RestPonse(ctx,
		http.StatusNotFound,
		gin.H{
			"message": msg,
		})
}

func RespNoContent(ctx *gin.Context, msg string) {
	RestPonse(ctx,
		http.StatusNoContent,
		gin.H{
			"message": msg,
		})
}

func RespInternalErr(ctx *gin.Context, msg string) {
	RestPonse(ctx,
		http.StatusInternalServerError,
		gin.H{
			"message": msg,
		})
}

func RespCreated(ctx *gin.Context, msg string) {
	RestPonse(ctx,
		http.StatusCreated,
		gin.H{
			"message": msg,
		})
}

func RestPonse(ctx *gin.Context, code int, param interface{}) {
	ctx.JSON(code, param)
}

func RespCatchSqlErr(ctx *gin.Context, err error) {
	switch err {
	case sql.ErrNoRows:
		RespNoContent(ctx, err.Error())
	default:
		RespInternalErr(ctx, err.Error())
	}
}
