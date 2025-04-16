package reports

import (
	"backend-api/internal/config"
	"backend-api/internal/handler"
	httpserv "backend-api/internal/server/http"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("cant initialize config: %v", err)
	}

	router := handler.NewRouter(cfg)

	server := httpserv.NewServer(cfg.RunAddress, router)

	go func() {
		// launch the server
		err = server.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("failed to run http server", err)
		}
	}()
	log.Printf("Server listening on: %s", fmt.Sprintf("address : %s", cfg.RunAddress))

	// if we catch the OS signal, then we do a graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// creating context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), cfg.CtxTimeout)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		log.Println("error in shutting down server")
	}

	log.Println("server stopped successfully")
}
