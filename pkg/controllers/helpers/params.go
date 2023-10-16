package helpers

import "github.com/gin-gonic/gin"

func GetId(ctx *gin.Context) string {
	return ctx.Param("id")
}
