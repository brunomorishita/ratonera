package connection

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	testDatabase = "goinggo"
)

// UserInfo is ...
type UserInfo struct {
	ID        string        `json:"id" bson:"id"`
	Gps       Gps           `json:"gps" bson:"gps"`
	Accel     Accelerometer `json:"accel" bson:"accel"`
	Timestamp int64         `json:"time" bson:"time"`
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
	upgrader   websocket.Upgrader
	session    *mgo.Session
	collection *mgo.Collection
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
	conn := Connection{upgrader: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		}},
		session: getSession(),
	}

	conn.collection = conn.session.DB(testDatabase).C("people")

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

		var record UserInfo
		err = json.Unmarshal(message, &record)

		log.Print(record)

		// Insert Datas
		err = conn.collection.Insert(record)
		if err != nil {
			log.Printf("shit happens in DB: %s", err.Error())
		}

		log.Printf("recv: %s", message)
	}
}

// GetUserInfo is ...
func (conn Connection) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	log.Printf("echo getuser")

	query := r.URL.Query()
	//fmt.Println("Request to create datapoint is ", query)
	id := query.Get("id")
	lastUpdate := query.Get("last_update")

	log.Printf("id: %s", id)
	log.Printf("lastUpdate: %s", lastUpdate)

	lastNum, err := strconv.ParseInt(lastUpdate, 10, 64)

	// layout := "2006-01-02T15:04:05.000Z"
	// t, err := time.Parse(layout, last_update)

	result := []UserInfo{}
	// err := conn.collection.Find(bson.M{"id": id, "time": bson.M{"$gte": lastUpdate}}).All(&result)
	err = conn.collection.Find(bson.M{"time": bson.M{"$gte": lastNum}}).All(&result)
	if err != nil {
		log.Printf("shit happens in DB: %s", err.Error())
	}

	j, _ := json.Marshal(result)

	w.Write(j)
}
