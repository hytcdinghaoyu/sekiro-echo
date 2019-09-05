package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"sekiro-echo/handler"

	"github.com/micro/go-web"
)

func main() {
	// Create service 这里需要注意使用的web.NewService 而不是micro.NewService 后文会有解释
	service := web.NewService(
		web.Name("go.micro.api.score"),
	)

	service.Init()

	// Create RESTful handler (using Gin)
	router := gin.Default()
	router.GET("/score/match", handler.FetchMatches)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
