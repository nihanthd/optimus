package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/nihanthd/optimus/exporter"

	"github.com/nihanthd/optimus/pi"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/go-resty/resty"

	"go.uber.org/fx"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nihanthd/optimus/env"
	"github.com/nihanthd/optimus/log"
	statsd2 "github.com/nihanthd/optimus/metrics/statsd"
	middleware2 "github.com/nihanthd/optimus/middleware"
	"github.com/nihanthd/optimus/pins"
	"github.com/nihanthd/optimus/pins/handler"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const serveDesc = `Run echo server.
Checkout /api/docs for available endpoints.
Example:
optimus server -c config.yaml`

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run echo server",
		Long:  serveDesc,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

func NewClient() *resty.Client {
	return resty.NewWithClient(&http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
			IdleConnTimeout:       60 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			DisableKeepAlives:     true,
		},
	})
}

// Configure router
func NewRouter(
	logger *zap.Logger,
	statsd *statsd.Client,
	config *statsd2.Config,
	base *handler.Handler,
	pinsHandler *handler.PinsHandler,
	metricsHandler *exporter.MetricsHandler,
) (*echo.Echo, error) {
	logger.Info("Creating new Echo Router")
	// Echo router
	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	e.HideBanner = true
	e.HidePort = true

	e.Use()

	public := e.Group("/")
	public.Use(
		middleware2.LogEachRequest(logger),
		middleware2.NewMetricsReporter(statsd, config.Enabled),
		middleware.Gzip(),
	)
	public.GET("healthz", base.GetHealthz)

	authorized := e.Group("/")

	authorized.Use(
		middleware2.LogEachRequest(logger),
		middleware2.NewMetricsReporter(statsd, config.Enabled),
		middleware.Gzip(),
	)

	authorized.GET("port/:id/status", pinsHandler.GetPinStatus)
	authorized.POST("port/:id", pinsHandler.TogglePin)
	authorized.GET("metrics", metricsHandler.GetMetrics)
	return e, nil
}

// Bootstrap the application run
func Start(server *http.Server, e *echo.Echo) {
	// Noop
}

func run() {
	// Intialize random seed
	rand.Seed(time.Now().UTC().UnixNano())

	// Initialize config module
	configModule := fx.Provide(
		func() *env.Config {
			return &config.Env
		},
		func() *log.Config {
			return &config.Log
		},
		func() *statsd2.Config {
			return &config.Statsd
		},
		func() *handler.Config {
			return &config.Http
		},
		func() *pi.Config {
			return &config.Pi
		},
	)

	// Initialize logger module
	fxLogger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(2 * time.Second)
		os.Exit(1)
	}

	// Initialize application
	app := fx.New(
		configModule,
		fx.Logger(log.NewFXLogger(fxLogger)),
		fx.Provide(
			log.NewFXLogger,
			log.NewConfigLogger,
			statsd2.NewClient,
			NewClient,
			NewRouter,
			handler.NewServer,
		),
		pi.Module,
		pins.Module,
		exporter.Module,
		fx.Invoke(Start),
	)

	// Run application
	startCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		fxLogger.Sugar().Fatal(err)
	}

	<-app.Done()
}
