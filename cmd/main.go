package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	ginger_api "github.com/petrrnnt/ginger-api"
	"github.com/petrrnnt/ginger-api/pkg/handler"
	"github.com/petrrnnt/ginger-api/pkg/repository"
	"github.com/petrrnnt/ginger-api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error initializing database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(ginger_api.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error starting http server: %s", err.Error())
		}
	}()

	logrus.Print("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Print("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error shutting down http server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error closing database connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
