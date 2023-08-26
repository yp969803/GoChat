package main 

import(
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
	handlers "gotut/specificChat/handlers"
)

func AddApproutes(route *mux.Router){
	log.Println("Loading Routes...")
	hub:= handlers.NewHub()
	go hub.Run()
	route.HandleFunc("/ws/{username}",func(responseWriter http.ResponseWriter, request *http.Request){
		var upgrader=websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
		}
          // Reading username from request parameter
		username:= mux.Vars(request)["username"]
		connection, err:= upgrader.Upgrade(responseWriter, request, nil)
		if err!=nil{
			log.Println(err)
			return
		}   
		handlers.CreateNewSocketUser(hub, connection, username)
	})
	log.Println("Routes are Loaded")
}