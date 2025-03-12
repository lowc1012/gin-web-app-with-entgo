package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/KimMachineGun/automemlimit/memlimit"
	"github.com/joho/godotenv"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/api"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/config"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
	"github.com/pbnjay/memory"
	"github.com/urfave/cli/v2"
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	// Load configuration (.env)
	err := godotenv.Load()
	if err != nil {
		log.Info("Error loading .env file, will use default configuration")
	}

	if err = config.Init(); err != nil {
		log.Errorw("Failed to initialize configuration", "error", err.Error())
		os.Exit(1)
	}

	// Initialize logger
	log.Init()

	// Set GOMEMLIMIT & GOMAXPROCS for performance consideration
	if _, ok := os.LookupEnv("GOMEMLIMIT"); ok {
		mem := debug.SetMemoryLimit(-1)
		log.Infow("Set memory limit by GOMEMLIMIT environment variable",
			"memory", memByteToStr(mem))
	} else {
		sysTotalMem := memory.TotalMemory()
		if limit, err := memlimit.FromCgroup(); err == nil && limit < sysTotalMem {
			mem, _ := memlimit.SetGoMemLimit(0.9)
			log.Infow("Set memory limit by cgroup",
				"memory", memByteToStr(mem),
				"system_total_mem", memByteToStr(sysTotalMem))
		} else {
			mem := int64(float64(sysTotalMem) * 0.9)
			debug.SetMemoryLimit(mem)
			log.Infow("Set memory limit by system total memory",
				"memory", memByteToStr(mem),
				"system_total_mem", memByteToStr(sysTotalMem))
		}
	}

	undo, err := maxprocs.Set(maxprocs.Logger(log.StdInfo))
	defer undo()
	if err != nil {
		log.Warnw("Failed to set GOMAXPROCS", err.Error())
	}

}

func memByteToStr[T int64 | uint64](v T) string {
	return fmt.Sprintf("%d MB", uint64(v)/1048576)
}

var Start = &cli.Command{
	Name:   "start",
	Usage:  "Start MyApp http server",
	Action: startHTTPServer,
}

func startHTTPServer(*cli.Context) error {
	// Create context for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	apiServer, err := api.StartAsync()
	if err != nil {
		log.Errorw("Failed to start MyApp", "error", err.Error())
		return cli.Exit("Failed to start MyApp", 1)
	}
	log.Info("MyApp started successfully")

	// blocks app here
	<-ctx.Done()

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		log.Error("API server shutdown timeout", "error", err.Error())
		return cli.Exit("API server shutdown timeout", 1)
	}

	log.Infow("MyApp shutdown gratefully", "event", "shutdown")
	return nil
}
