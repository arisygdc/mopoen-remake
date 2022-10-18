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

func RespInternalErr(ctx *gin.Context, msg string) {
	RestPonse(ctx,
		http.StatusNotFound,
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
		RespNotFound(ctx, err.Error())
	case sql.ErrConnDone:
		RespInternalErr(ctx, err.Error())
	}
}
