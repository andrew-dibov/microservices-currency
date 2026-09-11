package main

import (
	"database/sql"
	"errors"
	"microservices-currency/internal/clients"
	"microservices-currency/internal/configs"
	"microservices-currency/internal/loggers"
	"microservices-currency/internal/repositories"
	"microservices-currency/internal/servers"
	"microservices-currency/pkg/api/currency"
	"net"
	"os"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	appConfig := configs.NewAppConfig()
	appLogger := loggers.NewAppLogger(appConfig)

	appLogger.Info("config",
		"port", appConfig.App.Port,
		"prod", appConfig.App.Prod,
		"postgres", appConfig.PostgresDatabase.Address,
		"exchange", appConfig.ExchangeClient.Address,
	)

	/* --- --- --- */

	postgresDatabase, err := sql.Open("postgres", appConfig.PostgresDatabase.Address)
	if err != nil {
		appLogger.Error("NewPostgresDatabase returned error", "error", err)
		os.Exit(1)
	}
	defer postgresDatabase.Close()
	postgresRepository := repositories.NewPostgresRepository(postgresDatabase)

	/* --- --- --- */

	exchangeClient := clients.NewExchangeClient(&appConfig)

	/* --- --- --- */

	// UPDATES

	/* --- --- --- */

	appServer := servers.NewCurrencyServer(postgresRepository, appLogger)
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(4*1024*1024), grpc.MaxSendMsgSize(4*1024*1024),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    10 * time.Second, // вынести в конфиг
			Timeout: 1 * time.Second,
		}))

	currency.RegisterCurrencyServer(grpcServer, appServer)

	listener, err := net.Listen("tcp", appConfig.App.Port)
	if err != nil {
		appLogger.Error("ERRROR", err)
		os.Exit(1)
	}

	go func() {
		if err := grpcServer.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			appLogger.Error("ERROR")
			os.Exit(1)
		}
	}()

	/* --- --- --- */

	// SHUTDOWN
}
