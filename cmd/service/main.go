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
	"github.com/golovpeter/ozon-trainee-task/internal/interceptors/accesslog"
	urlRepository "github.com/golovpeter/ozon-trainee-task/internal/repository/urls"
	urlsService "github.com/golovpeter/ozon-trainee-task/internal/service/urls"
	"github.com/golovpeter/ozon-trainee-task/protos/gen/go/url_shortener"
	"github.com/jmoiron/sqlx"
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

	var urlRepo urlRepository.Repository
	var dbConn *sqlx.DB

	memoryMode := os.Args[1]

	switch memoryMode {
	case "memory":
		urlRepo = urlRepository.NewRepositoryInMemory()
	case "postgres":
		dbConn, err = common.CreateDbClient(cfg.Database)
		if err != nil {
			logger.Fatalln("error to create database client:", err)
		}

		urlRepo = urlRepository.NewRepositoryPostgres(dbConn)
	default:
		logrus.Fatal("invalid memory mode")
	}

	urlService := urlsService.NewService(urlRepo)
	urlsHandler := urls.NewHandler(logger, urlService)

	accessLog := accesslog.NewInterceptor(logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(accessLog.Interceptor),
	)

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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Fatalln("cant listen port:", err)
	}

	logger.Infof("grpc server starting at :%d", cfg.Server.Port)
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
