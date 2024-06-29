package api

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/lowc1012/gin-web-app-with-entgo/internal/api/v1"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/config"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
)

func StartAsync() (server *http.Server, err error) {
	setApiServerMode()

	router := gin.New()

	if config.IsDevEnv() {
		router.Use(gin.Logger())
	}

	// register a recovery middleware
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Errorw("Received panic", "error", err, "stack", string(debug.Stack()))
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// register a cors middleware
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	// registers handlers
	v1Group := router.Group("/v1")
	v1.Mount(v1Group)

	server = &http.Server{
		Addr:              fmt.Sprintf(":%d", config.Global.ApiServerPort),
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// use goroutine because http.ListenAndServe() generates blocking call
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info("API server closed")
			} else {
				log.Errorw("API server got error", "Error", err.Error())
			}
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
