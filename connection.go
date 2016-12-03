package ratonera

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Connection is ...
type Connection struct {
	upgrader websocket.Upgrader
}

// NewConnection is ...
func NewConnection() Connection {
	return Connection{upgrader: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}}
}

// HandleWebsocket connection.
func (conn Connection) HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	c, err := conn.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
