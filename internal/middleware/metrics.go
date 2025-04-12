package middleware

import (
	"avito-pvz/internal/metrics"
	"github.com/labstack/echo/v4"
	"time"
)

// MetricsMiddleware collect info about number of requests and its duration
func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)

		method := c.Request().Method

		duration := time.Since(start).Seconds()
		metrics.IncrementRequestCount(method)
		metrics.ObserveRequestDuration(method, duration)

		return err
	}
}
