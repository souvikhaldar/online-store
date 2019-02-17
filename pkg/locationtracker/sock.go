package locationtracker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type longLatStruct struct {
	Long  float64 `json:"longitude"`
	Lat   float64 `json:"latitude"`
	Agent string  `json:"agent"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *longLatStruct)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func writer(coord *longLatStruct) {
	broadcast <- coord
}

func LongLatHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("--Running in LongLatHandler--")
	var coordinates longLatStruct
	if err := json.NewDecoder(r.Body).Decode(&coordinates); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", 500)
		return
	}
	defer r.Body.Close()
	go writer(&coordinates)
}

var reqAgent string

func WsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("--Running in WsHandler--")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	queryVal := r.URL.Query()
	reqAgent = queryVal.Get("agent")
	// register client
	clients[ws] = true
}

func Echo() {
	for {
		val := <-broadcast
		latlong := fmt.Sprintf("%f %f %s", val.Lat, val.Long, val.Agent)
		// send to every client that is currently connected
		for client := range clients {
			if val.Agent == reqAgent {
				err := client.WriteMessage(websocket.TextMessage, []byte(latlong))
				if err != nil {
					log.Printf("Websocket error: %s", err)
					client.Close()
					delete(clients, client)
				}
			} else {
				err := client.WriteMessage(websocket.TextMessage, []byte("Invalid Agent"))
				if err != nil {
					log.Printf("Websocket error: %s", err)
					client.Close()
					delete(clients, client)
				}
			}

		}
	}
}
