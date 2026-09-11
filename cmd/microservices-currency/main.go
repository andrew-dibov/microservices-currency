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
	"os/signal"
	"syscall"
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

	exchangeClient, err := clients.NewExchangeClient(&appConfig)
	if err != nil {
		appLogger.Error("NewExchangeClient returned error", "error", err)
		os.Exit(1)
	}

	/* --- --- --- */

	// UPDATES

	/* --- --- --- */

	appServer := servers.NewAppServer(postgresRepository, appLogger)
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(4*1024*1024), grpc.MaxSendMsgSize(4*1024*1024),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    appConfig.App.KeepaliveTime,
			Timeout: appConfig.App.KeepaliveTimeout,
		}))

	currency.RegisterCurrencyServer(grpcServer, appServer)

	appListener, err := net.Listen("tcp", ":"+appConfig.App.Port)
	if err != nil {
		appLogger.Error("appListener returned error", err)
		os.Exit(1)
	}

	go func() {
		if err := grpcServer.Serve(appListener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			appLogger.Error("grpcServer returned error", "error", err)
			os.Exit(1)
		}
	}()

	/* --- --- --- */

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	appLogger.Info("server shutting down")

	done := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		appLogger.Info("server stopped")
	case <-time.After(appConfig.App.ShutdownTimeout):
		appLogger.Warn("server forced to stop")
		grpcServer.Stop()
	}
}
