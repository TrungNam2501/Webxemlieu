package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Webxemlieu/config"
	"Webxemlieu/database"
	"Webxemlieu/middleware"
	"Webxemlieu/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)
	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Printf("warning: SetTrustedProxies: %v", err)
	}

	r.Use(middleware.CORS(cfg.CORSAllowedOrigins))

	if err := database.Connect(cfg.PostgresDSN()); err != nil {
		log.Fatalf("database: %v", err)
	}

	routes.SetupAPI(r)
	routes.SetupSPA(r, cfg.FrontendDist)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("server: lắng nghe trên :%s (mode=%s)", cfg.Port, cfg.GinMode)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server: shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced shutdown: %v", err)
	}
	log.Println("server: stopped")
}
