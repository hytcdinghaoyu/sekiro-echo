package main

import (
	"log"
	_ "sekiro_echo/app/idip/routers"

	"github.com/astaxie/beego"
	"github.com/micro/go-web"
)

func main() {

	service := web.NewService(
		web.Name("go.micro.api.idip"),
	)

	service.Init()

	app := beego.BeeApp
	// Register Handler
	service.Handle("/", app.Handlers)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
