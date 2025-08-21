package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shrikant9024/go-tutorials/internal/config"
	"github.com/shrikant9024/go-tutorials/internal/http/handlers/student"
)

func main() {
	//load config

	cfg := config.MustLoad()

	//db setup
	//router setup
	router := http.NewServeMux()

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to the server"))

	// })

	router.Handle("POST /api/students", student.New())

	//setup server
	server := http.Server{

		Addr:    cfg.Address,
		Handler: router,
	}
	slog.Info("server started", slog.String("address", cfg.Address))
	// fmt.Printf("Server is started %s", cfg.HttpServer.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown", slog.String("Error", err.Error()))
	}

	slog.Info("server shut down successfully")

}
