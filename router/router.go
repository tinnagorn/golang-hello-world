package router

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/tOnkowzl/libs/middleware"
)

var (
	buildstamp = time.Now().String()
	githash    = "developing"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	m := middleware.New("hello-world")
	m.Skipper = func(c echo.Context) bool {
		return c.Path() == "/builds" || c.Path() == "/health"
	}

	e.Use(m.Build(buildstamp, githash))
	e.Use(m.RequestID())
	e.Use(m.Recover())
	e.Use(m.LogRequestInfo())
	e.Use(m.Logger())

	return e
}

func Runs(router *echo.Echo) {
	log.Printf("starting %s", viper.GetString("app.name"))
	log.Printf("application serve at port %s", "1323")
	log.Println(router.Start(":1323"))
}

func Shutdown(router *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := router.Shutdown(context.Background()); err != nil {
		log.Fatal("shutdown server: ", err)
	}
}
