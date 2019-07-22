package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Controller struct {
	Repo GreetingRepo
}

func (ctl *Controller) Metrics(c echo.Context) error {
	req := c.Request()
	res := c.Response()
	ph := promhttp.Handler()
	ph.ServeHTTP(res.Writer, req)
	return nil
}

//func (ctl *Controller) ViewGreetings(c echo.Context) error {
//	var g []*Greeting
//	var err error
//	if g, err = ctl.Repo.SelectGreetingsLimit10(); err != nil {
//		return c.JSON(http.StatusInternalServerError, nil)
//	}
//
//	return c.JSON(http.StatusOK, g)
//}
//
//func (ctl *Controller) CreateGreetings(c echo.Context) error {
//	g := &Greeting{}
//	if err := c.Bind(g); err != nil {
//		return c.JSON(http.StatusBadRequest, nil)
//	}
//
//	if g.Text == "" {
//		return c.JSON(http.StatusBadRequest, nil)
//	}
//	if err := ctl.Repo.InsertGreeting(*g); err != nil {
//		return c.JSON(http.StatusInternalServerError, nil)
//	}
//
//	return c.JSON(http.StatusOK, g)
//}

func (ctl *Controller) Sleep(c echo.Context) error {
	mean, err := strconv.ParseFloat(c.Param("mean"), 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	std, err := strconv.ParseFloat(c.Param("std"), 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	g, err := ctl.Repo.SelectGreetingsLimit10()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	time.Sleep(time.Duration(math.Abs(rand.NormFloat64()*std+mean)) * time.Millisecond)
	return c.JSON(http.StatusOK, g)
}

func (ctl *Controller) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
