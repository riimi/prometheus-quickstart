package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

func MiddlewareAPIProfiler() echo.MiddlewareFunc {
	sv := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "api_durations_msec",
			Help:       "API latency distributions",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"api", "code"},
	)
	prometheus.MustRegister(sv)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now().UnixNano()
			err := next(c)
			elapsedTimeMsec := (time.Now().UnixNano() - startTime) / 1000000

			url := c.Request().URL.String()
			code := strconv.Itoa(c.Response().Status)
			sv.WithLabelValues(url, code).Observe(float64(elapsedTimeMsec))
			return err
		}
	}
}
