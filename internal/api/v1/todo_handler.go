package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/db"
)

func listTodoHandler(c *gin.Context) {
	todos, err := db.MustClient().Todo.Query().All(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}
