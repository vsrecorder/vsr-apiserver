package helpers

import "github.com/gin-gonic/gin"

func SetUID(ctx *gin.Context, value string) {
	ctx.Set("uid", value)
}

func GetUID(ctx *gin.Context) (uid string, exists bool) {
	value, exists := ctx.Get("uid")

	switch v := value.(type) {
	case string:
		uid = v
	default:
		uid = ""
	}

	return uid, exists
}
