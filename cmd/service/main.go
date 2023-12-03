package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/golovpeter/ozon-trainee-task/internal/common"
	"github.com/golovpeter/ozon-trainee-task/internal/config"
	"github.com/golovpeter/ozon-trainee-task/internal/handler/urls"
	urlRepository "github.com/golovpeter/ozon-trainee-task/internal/repository/urls"
	urlsService "github.com/golovpeter/ozon-trainee-task/internal/service/urls"
	"github.com/golovpeter/ozon-trainee-task/protos/gen/go/url_shortener"
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Fatalln("cant listen port:", err)
	}

	//urlRepo := urlRepository.NewRepositoryPostgres(dbConn)
	urlRepo := urlRepository.NewRepositoryInMemory()
	urlService := urlsService.NewService(urlRepo)
	urlsHandler := urls.NewHandler(logger, urlService)

	server := grpc.NewServer()

	url_shortener.RegisterUrlShortenerServer(server, urlsHandler)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-signalChan

		err = dbConn.Close()
		if err != nil {
			return
		}

		server.GracefulStop()
	}()

	logger.Infof("grpc server starting at :%d", cfg.Server.Port)
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
