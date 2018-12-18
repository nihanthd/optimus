package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"
	"github.com/nihanthd/optimus/log"
	"github.com/patrickmn/go-cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gopkg.in/bluesuncorp/validator.v5"
)

type Handler struct {
	log       *zap.Logger
	cache     *cache.Cache
	statsd    *statsd.Client
	validator *validator.Validate
}

type HandlerParams struct {
	fx.In

	Log       *zap.Logger
	Cache     *cache.Cache        `optional:"true"`
	Statsd    *statsd.Client      `optional:"true"`
	Validator *validator.Validate `optional:"true"`
}

type (
	Config struct {
		Port int32 `yaml:"port"`
	}
)

func (h *Handler) GetHealthz(c echo.Context) error {
	c.Response().Header().Set("Server-Status", "OK")
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) GetMetrics(c echo.Context) error {
	promhttp.Handler().ServeHTTP(c.Response().Writer, c.Request())
	return nil
}

func NewServer(lc fx.Lifecycle, logger *zap.Logger, config *Config, echo *echo.Echo) (*http.Server, error) {
	logger.Sugar().Info("Executing NewEcho on Port=", config.Port)

	echo.HideBanner = true
	echo.Logger = log.NewEchoLogger(logger)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", config.Port),
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Sugar().Info("Starting HTTP server.")
			go echo.StartServer(server)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Sugar().Info("Stopping HTTP server.")
			return echo.Shutdown(ctx)
		},
	})

	return server, nil
}

func NewHandler(p HandlerParams) *Handler {
	return &Handler{
		log:       p.Log,
		statsd:    p.Statsd,
		cache:     p.Cache,
		validator: p.Validator,
	}
}
