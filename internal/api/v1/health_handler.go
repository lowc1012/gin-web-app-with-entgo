package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthHandler(c *gin.Context) {
	// TODO: check dependencies like database

	c.Status(http.StatusOK)
}
