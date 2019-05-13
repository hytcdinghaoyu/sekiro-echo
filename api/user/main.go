package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"sekiro_echo/handler"

	"sekiro_echo/conf"

	"sekiro_echo/lib/middleware"

	"github.com/micro/go-web"
)

func main() {
	// Create service 这里需要注意使用的web.NewService 而不是micro.NewService 后文会有解释
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	// init conf
	if err := conf.InitConfig("./conf/conf.toml"); err != nil {
		log.Panic(err)
	}

	service.Init()

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
