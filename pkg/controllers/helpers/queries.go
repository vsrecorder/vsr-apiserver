package helpers

import "github.com/gin-gonic/gin"

func GetStartDate(ctx *gin.Context) string {
	return ctx.Query("start_date")
}

func GetEndDate(ctx *gin.Context) string {
	return ctx.Query("end_date")
}

func GetPage(ctx *gin.Context) string {
	return ctx.Query("page")
}
