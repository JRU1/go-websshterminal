package router

import (
	"time"

	"github.com/c/websshterminal.io/handler"
	"github.com/c/websshterminal.io/middlewares"
	"github.com/c/websshterminal.io/ubzer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// go:embed dist
//var dist embed.FS

// RunSshTerminal
func RunSshTerminal() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS, echo.PATCH, echo.DELETE},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 24,
		}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip} time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
		Output: ubzer.EchoLog,
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middlewares.RequestLog())
	e.Use(middleware.BodyDumpWithConfig(middlewares.DefaultBodyDumpConfig))

	e.Static("/", "dist")
	e.Static("/ssh/node", "dist")
	e.Static("/ssh/dial", "dist")

	//distFS, err := fs.Sub(dist, "dist")
	//if err != nil {
	//	panic(err)
	//}
	//distHandler := http.FileServer(http.FS(distFS))
	//e.GET("/", echo.WrapHandler(distHandler))
	//e.GET("/css/*", echo.WrapHandler(http.StripPrefix("/css/", distHandler)))
	//e.GET("/js/*", echo.WrapHandler(http.StripPrefix("/js/", distHandler)))

	e.GET("/ssh", handler.ShellWeb)

	e.Logger.Fatal(e.Start(":6666"))
}
