package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"sekiro_echo/handler"
	. "sekiro_echo/conf"
	"flag"
)

var (
	confFilePath string
)

const (
	DefaultConfFilePath = "conf/conf.toml"
)

func init() {
	flag.StringVar(&confFilePath, "c", DefaultConfFilePath, "配置文件路径")
	flag.Parse()

}

func main() {
	e := echo.New()

	// init conf
	if err := InitConfig(confFilePath); err != nil {
		log.Panic(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/matches", handler.FetchMatches)

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(Conf.Jwt.Secret),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	

	// Routes
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	e.POST("/posts", handler.CreatePost)
	e.GET("/feed", handler.FetchPost)

	e.Logger.Fatal(e.Start(":1323"))
}