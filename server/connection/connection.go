package connection

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

// Connection is ...
type Connection struct {
	upgrader websocket.Upgrader
	db       *Database
}

// NewConnection is ...
func NewConnection(mongoServer string) Connection {
	conn := Connection{upgrader: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		}},
		db: NewDatabase(mongoServer),
	}

	return conn
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
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		err = conn.db.ProcessMessage(message)
		if err != nil {
			continue
		}

		log.Printf("recv: %s", message)
	}
}

// GetUserInfo is ...
func (conn Connection) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	lastUpdate := query.Get("last_update")

	log.Printf("id: %s", id)
	log.Printf("lastUpdate: %s", lastUpdate)

	timestamp, _ := strconv.ParseInt(lastUpdate, 10, 64)

	result, err := conn.db.GetUserInfo(id, timestamp)
	if err != nil {
		log.Printf("fail %s:", err.Error())
	}

	j, _ := json.Marshal(result)

	w.Write(j)
}
