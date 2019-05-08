package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"sekiro_echo/handler"

	"sekiro_echo/conf"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/micro/go-web"
)

func main() {
	// Create service 这里需要注意使用的web.NewService 而不是micro.NewService 后文会有解释
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	// init conf
	if err := conf.InitConfig("../conf/conf.toml"); err != nil {
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
	private := router.Group("/user/private")
	{
		private.Use(jwt.Auth(conf.Conf.Jwt.Secret))
		private.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello from private"})
		})

	}

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
