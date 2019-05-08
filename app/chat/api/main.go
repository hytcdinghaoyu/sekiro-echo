package main

import (
	"log"
	"net/http"

	"github.com/micro/go-web"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	// New web service
	service := web.NewService(
		web.Name("go.micro.api.chat"),
	)

	// parse command line
	service.Init()

	m := melody.New()
	m.HandleDisconnect(HandleConnect)

	// Handle websocket connection
	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	// run service
	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}

// 处理用户连接
func HandleConnect(session *melody.Session) {
	log.Println("new connection ======>>>")
}
