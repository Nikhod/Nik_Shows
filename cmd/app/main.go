package main

import (
	"Nik/internal/configs"
	"Nik/internal/handlers"
	"Nik/internal/repositories"
	"Nik/internal/services"
	"Nik/loggers"
	"Nik/pkg/database"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	initLogger, err := loggers.InitLogger()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Println(err)
		}
	}(initLogger)

	// Get configs
	config, err := configs.InitConfigs()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Connect To Database
	db, err := database.ConnectToDB(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Main Stream Working
	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service, initLogger)

	// Init Mux
	mux := InitMux(handler)
	log.Println("Mux Connected Successfully!")
	srv := http.Server{
		Addr:    config.Server.Host + config.Server.Port,
		Handler: mux,
	}

	log.Println("Start!")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
