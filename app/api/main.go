package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/micro/go-web"
)

type Say struct{}

func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func main() {
	// Create service 这里需要注意使用的web.NewService 而不是micro.NewService 后文会有解释
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
	)

	service.Init()

	// Create RESTful handler (using Gin)
	say := new(Say)
	router := gin.Default()
	router.GET("/greeter", say.Anything)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
