package middleware

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"
)

// Metrics around request
func NewMetricsReporter(statsd *statsd.Client, enabled bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			if enabled {
				statsd.Histogram("optimus.http.in.latency", time.Now().Sub(start).Seconds()*1000, []string{}, 100)
			}
			return nil
		}
	}
}

func LogEachRequest(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}
			log.Info("Serving path", zap.String("path", c.Request().URL.Path), zap.String("user", c.Request().Header.Get("x-email")), zap.String("responseCode", fmt.Sprintf("%d", c.Response().Status)))
			return nil
		}
	}
}
