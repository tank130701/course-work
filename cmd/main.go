package main

import (
	"context"
	"errors"
	"github.com/tank130701/course-work/todo-app/back-end/internal/app"
	"github.com/tank130701/course-work/todo-app/back-end/internal/config"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/tank130701/course-work/todo-app/back-end/internal/service"
	"github.com/tank130701/course-work/todo-app/back-end/internal/transport/http/handler"
)

// @title Todo App API
// @version 1.0
// @description API App for TodoCategories Application

// @host localhost:5000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	//logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.DebugLevel)
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	connData, err := getConnectionData(cfg, cfg.ConnectionType)
	if err != nil {
		logrus.Fatalf("error: %s", err)
		return
	}

	logrus.Infof("Selected connection: %s", cfg.ConnectionType)

	db, err := initDbConnection(connData, cfg.ConnectionType)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	var repos *repository.Repository
	if cfg.ConnectionType == "postgres" {
		repos = repository.NewPostgresRepository(db)
	} else if cfg.ConnectionType == "mysql" {
		repos = repository.NewPostgresRepository(db) // TODO: Fix this
	}

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(app.App)
	go func() {
		if err := srv.Run(cfg.HttpPort, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func getConnectionData(cfg *config.Config, connectionType string) (config.DBConfig, error) {
	switch connectionType {
	case "postgres":
		return cfg.PostgreSQL, nil
	case "mysql":
		return cfg.MySQL, nil
	default:
		return config.DBConfig{}, errors.New("invalid connection type")
	}
}
