package routes

import "github.com/gin-gonic/gin"

// Protected wraps a router group with middleware(s)
func Protected(group *gin.RouterGroup, middleware ...gin.HandlerFunc) *gin.RouterGroup {
	return group.Group("", middleware...)
}
