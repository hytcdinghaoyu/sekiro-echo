package main

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/consul"
	"log"

	"github.com/gin-gonic/gin"

	"sekiro-echo/handler"

	"sekiro-echo/conf"

	"sekiro-echo/lib/middleware"

	"github.com/micro/go-micro/config/encoder/toml"
	"github.com/micro/go-micro/web"
)

func main() {
	// Create service 这里需要注意使用的web.NewService 而不是micro.NewService 后文会有解释
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	//load config from consul
	e := toml.NewEncoder()
	configure := config.NewConfig()
	_ = configure.Load(consul.NewSource(
		consul.WithAddress("129.211.75.241:8500"),
		consul.WithPrefix("/sekiro/config"),
		source.WithEncoder(e),
	))
	_ = configure.Get("sekiro", "config").Scan(&conf.Conf)

	_ = service.Init()

	// Create RESTful handler (using Gin)
	router := gin.Default()

	// public group without auth
	public := router.Group("/user")
	{
		public.POST("/signup", handler.Signup)
		public.POST("/login", handler.Login)

	}

	// private group use jwt auth
	private := router.Group("/user/post")
	{
		private.Use(middleware.JWT(conf.Conf.Jwt.Secret))
		private.POST("/get", handler.FetchPost)
		private.POST("/create", handler.CreatePost)

	}

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
