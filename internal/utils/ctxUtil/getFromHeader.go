package ctxutil

import "github.com/gin-gonic/gin"

// Get a passed key from the header of a gin context
func GetFromHeader(ctx *gin.Context, key string) string {
	header := ctx.GetHeader(key)
	return header
}

// Get a user_uuid key from the header of a gin context
func GetUserUuidFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("user_uuid")
	return header
}

// get user_units key from header of gin context
func GetUserUnitsFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("user_unit_class")
	return header
}