package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
	"net/http"
	"online_song_library/internal/db"
	"online_song_library/internal/rest"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Online Song Library API
// @version 1.0
// @description Swagger API for Golang Online Song Library
// @terms http://swagger.io/terms/
// @BasePath /api/v1
func main() {
	logrus.SetOutput(os.Stdout)

	DB, err := db.Init()
	if err != nil {
		logrus.Error(err)
	}

	mux := rest.CreateRestService(DB)

	srv := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}

	logrus.Info("Server is running on port 8080")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Fatal(err)
		}
	}()

	<-stop
	logrus.Info("Server stopping...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Println("Server down")
}
