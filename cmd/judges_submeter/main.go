package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/RodrigoCMoraes/judges_submeter/internal/middlewares"
	"github.com/RodrigoCMoraes/judges_submeter/internal/submeter"
)

func main() {

	err := godotenv.Load()
	if err != nil && os.Getenv("ENV") != "PROD" {
		log.Fatal("Error loading .env file")
	}

	http.Handle("/submit", middlewares.LogHandler(submeter.Submit))

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	logrus.WithField("addr", addr).Info("starting server")

	if err := http.ListenAndServe(addr, nil); err != nil {
		logrus.WithField("event", "start server").Fatal(err)
	}
}
