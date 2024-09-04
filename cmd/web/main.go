package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/aikuci/go-subdivisions-id/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	zapLog := config.NewZapLog(viperConfig)
	db := config.NewDatabase(viperConfig)
	validate := config.NewValidator(viperConfig)

	file, err := os.OpenFile("storage/log/fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app := config.NewFiber(viperConfig, &config.AppOptions{LogWriter: file})

	// Ref: https://github.com/gofiber/fiber/issues/899#issuecomment-1429739373
	var serverShutdown sync.WaitGroup
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		if err := app.ShutdownWithTimeout(60 * time.Second); err != nil {
			log.Fatalf("failed to shut down server: %v", err)
		}
	}()

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      zapLog,
		Validate: validate,
		Config:   viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	log.Printf("Starting REST, listening at %d\n", webPort)
	if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	serverShutdown.Wait()
	log.Println("Running cleanup tasks...")
	// Cleanup tasks
}
