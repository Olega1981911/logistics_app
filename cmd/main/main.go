package main

import (
	"log/slog"
	"logistics_App/internal/config"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	//TODO: init config
	cfg := config.MustLoadConfig()
	//TODO: init logger
	log := setupLogger(cfg.Env)
	log.Info("starting server", slog.String("env", cfg.Env))
	//TODO: init DB(storage)
	/*db := database.Connect(cfg.DatabaseURL)

		repo := repository.NewCustomerRepo(db)

		service := service.NewCustomerService(repo)

		handler := handler.NewCustomerHandler(service)
	//TODO: init Router
		router := chi.NewRouter()
		handler.RegisterRoutes(router)
	//TODO: init server
		http.ListenAndServe(":8080", router)*/
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
