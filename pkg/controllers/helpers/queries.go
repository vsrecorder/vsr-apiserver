package helpers

import (
	"github.com/gin-gonic/gin"
)

func GetStartDate(ctx *gin.Context) (start_date string) {
	return ctx.Query("start_date")
}

func GetEndDate(ctx *gin.Context) (end_date string) {
	return ctx.Query("end_date")
}

func GetPage(ctx *gin.Context) (page string) {
	return ctx.Query("page")
}
