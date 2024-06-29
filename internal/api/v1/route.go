package v1

import "github.com/gin-gonic/gin"

func Mount(r *gin.RouterGroup) {
	r.GET("/health", healthHandler)
}
