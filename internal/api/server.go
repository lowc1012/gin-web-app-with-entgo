package api

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/api/v1"
	taskApp "github.com/lowc1012/gin-web-app-with-entgo/internal/application/task"
	todoApp "github.com/lowc1012/gin-web-app-with-entgo/internal/application/todo"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/config"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/persistence"
)

func StartAsync() (server *http.Server, err error) {
	setApiServerMode()

	// initial db client
	entClient := persistence.MustClient()

	router := gin.New()
	router.GET("/health", v1.HealthHandler)

	if config.IsDevEnv() {
		router.Use(gin.Logger())
	}

	// register recovery middleware
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Errorw("Received panic", "error", err, "stack", string(debug.Stack()))
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// register cors middleware
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	// Repository implementation
	taskRepo := persistence.NewEntTaskRepository(entClient)
	todoRepo := persistence.NewEntTodoRepository(entClient)

	// Application services
	taskSvc := taskApp.NewService(taskRepo)
	todoSvc := todoApp.NewService(todoRepo)

	// registers handlers
	taskHandler := v1.NewTaskHandler(taskSvc)
	todoHandler := v1.NewTodoHandler(todoSvc)

	v1Group := router.Group("/api/v1")
	taskHandler.RegisterTaskRoutes(v1Group)
	todoHandler.RegisterTodoRoutes(v1Group)

	server = &http.Server{
		Addr:              fmt.Sprintf(":%d", config.Global.ApiServerPort),
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// use goroutine because http.ListenAndServe() generates blocking call
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorw("API server error", "error", err.Error())
		}
	}()

	return server, nil
}

func setApiServerMode() {
	if config.IsProdEnv() {
		gin.SetMode(gin.ReleaseMode)
	} else if config.IsTestEnv() {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	}
}
