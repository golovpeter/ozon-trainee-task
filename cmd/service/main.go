package main

import (
	"net"
	"ozon_trainee_task/internal/common"
	"ozon_trainee_task/internal/config"
	"ozon_trainee_task/internal/handler/urls"
	urlRepository "ozon_trainee_task/internal/repository/urls"
	urlsService "ozon_trainee_task/internal/service/urls"
	"ozon_trainee_task/protos/gen/go/url_shortener"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger := logrus.New()

	cfg, err := config.Parse()
	if err != nil {
		logger.Fatalln("error to parse config file:", err)
	}

	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		logger.Fatalln("error to parse logger level:", err)
	}

	logger.SetLevel(level)

	dbConn, err := common.CreateDbClient(cfg.Database)
	if err != nil {
		logger.Fatalln("error to create database client:", err)
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatalln("cant listen port:", err)
	}

	urlRepo := urlRepository.NewRepository(dbConn)
	urlService := urlsService.NewService(urlRepo)
	urlsHandler := urls.NewHandler(logger, urlService)

	server := grpc.NewServer()

	url_shortener.RegisterUrlShortenerServer(server, urlsHandler)

	logger.Info("grpc server starting at :8081")
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
