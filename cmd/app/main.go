package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/KimMachineGun/automemlimit/memlimit"
	"github.com/joho/godotenv"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/config"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
	"github.com/pbnjay/memory"
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	// Load configuration (.env)
	err := godotenv.Load()
	if err != nil {
		log.Info("Error loading .env file, will use default configuration")
	}

	if err = config.Init(); err != nil {
		log.Fatalw("Failed to initialize configuration", "error", err.Error())
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

func main() {
	if err := rootCmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
