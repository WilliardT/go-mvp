package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	users_transport_http "github.com/WilliardT/go-mvp/internal/features/users/transport/http/transport"
	core_http_server "github.com/WilliardT/go-mvp/internal/core/transport/http/server"
	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	defer cancel()

	logger, err := core_logger.NewConfigMust()

	if err != nil {
		fmt.Println("failed to init application logger:", err)

		os.Exit(1)
	}
	
	defer logger.Close()

	logger.Debug("Starting application")

	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)

	usersRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterAPIRouters(usersRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
	)

	httpServer.RegisterAPIRouters(*apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}