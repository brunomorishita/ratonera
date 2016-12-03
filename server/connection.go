package ratonera

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/websocket"
)

const (
	testDatabase = "goinggo"
)

// Person is ...
type Person struct {
	ID        string        `json:"id" bson:"id"`
	Gps       Gps           `json:"gps" bson:"gps"`
	Accel     Accelerometer `json:"accel" bson:"accel"`
	Timestamp time.Time     `json:"time" bson:"time"`
}

// Gps is ...
type Gps struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lgt float64 `json:"lgt" bson:"lgt"`
}

// Accelerometer is ...
type Accelerometer struct {
	Z float64 `json:"z" bson:"z"`
	Y float64 `json:"y" bson:"y"`
	X float64 `json:"x" bson:"x"`
}

// Connection is ...
type Connection struct {
	upgrader websocket.Upgrader
	session  *mgo.Session
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

// NewConnection is ...
func NewConnection() Connection {
	return Connection{upgrader: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		}},
		session: getSession()}
}

// HandleWebsocket connection.
func (conn Connection) HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	// Collection People
	dbCol := conn.session.DB(testDatabase).C("people")

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

		var record Person
		err = json.Unmarshal(message, &record)

		log.Print(record)

		// Insert Datas
		err = dbCol.Insert(record)
		if err != nil {
			log.Printf("shit happens in DB: %s", err.Error())
		}

		log.Printf("recv: %s", message)

		result := []Person{}
		err = dbCol.Find(bson.M{}).All(&result)
		j, _ := json.Marshal(result)

		err = c.WriteMessage(mt, j)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
