package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"sekiro_echo/handler"

	"github.com/micro/go-web"
)

func main() {
	// Create service 这里需要注意使用的web.NewService 而不是micro.NewService 后文会有解释
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	service.Init()

	// Create RESTful handler (using Gin)
	router := gin.Default()
	router.POST("user/signup", handler.Signup)
	router.POST("user/login", handler.Login)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
