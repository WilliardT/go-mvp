package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/WilliardT/go-mvp/internal/core/config"
	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_pgx_pool "github.com/WilliardT/go-mvp/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/WilliardT/go-mvp/internal/core/transport/http/middleware"
	core_http_server "github.com/WilliardT/go-mvp/internal/core/transport/http/server"
	products_postgres_repository "github.com/WilliardT/go-mvp/internal/features/products/repository/postgres"
	products_service "github.com/WilliardT/go-mvp/internal/features/products/service"
	products_transport_http "github.com/WilliardT/go-mvp/internal/features/products/transport/http"
	statistics_postgres_repository "github.com/WilliardT/go-mvp/internal/features/statistics/repository/postgres"
	statistics_service "github.com/WilliardT/go-mvp/internal/features/statistics/service"
	statistics_transport_http "github.com/WilliardT/go-mvp/internal/features/statistics/transport/http"
	users_postgres_repository "github.com/WilliardT/go-mvp/internal/features/users/repository/postgres"
	users_service "github.com/WilliardT/go-mvp/internal/features/users/service"
	users_transport_http "github.com/WilliardT/go-mvp/internal/features/users/transport/http"
	"go.uber.org/zap"

	_ "github.com/WilliardT/go-mvp/docs"
)


// @title       Go MVP API
// @version     1.0
// @description GO MVP application REST-API schema
// @host        localhost:5050
// @BasePath    /api/v1
func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	defer cancel()

	loggerConfig := core_logger.NewConfigMust()

	logger, err := core_logger.NewLogger(*loggerConfig)
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("Initializing postgres connection pool...")

	pool, err := core_pgx_pool.NewPgxConnectionPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)

	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}

	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))

	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "products"))

	productsRepository := products_postgres_repository.NewProductsRepository(pool)
	productsService := products_service.NewProductsService(productsRepository)
	productsTransportHTTP := products_transport_http.NewProductsHTTPHandler(productsService)

	logger.Debug("initializing feature", zap.String("feature", "statistics"))

	statisticsRepository := statistics_postgres_repository.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	logger.Debug("initializing HTTP server...")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.ReqestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(productsTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(statisticsTransportHTTP.Routes()...)
	
	httpServer.RegisterAPIRouters(apiVersionRouter)

	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
