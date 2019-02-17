package locationtracker

import (
	"fmt"
	"log"
	"net/http"
)

//var upgrader = websocket.Upgrader{}

func Track(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)

	}
	fmt.Fprintf(w, "recv: %s", message)
	err = c.WriteMessage(mt, message)
	if err != nil {
		log.Println("write:", err)

	}
}
