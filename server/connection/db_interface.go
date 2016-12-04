package connection

import (
	"encoding/json"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	testDatabase = "goinggo"
)

// Database is ...
type Database struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func getSession(mongoServer string) *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://" + mongoServer)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

// NewDatabase is ...
func NewDatabase(mongoServer string) *Database {
	conn := &Database{
		session: getSession(mongoServer),
	}

	conn.collection = conn.session.DB(testDatabase).C("people")

	return conn
}

// ProcessMessage is ...
func (db *Database) ProcessMessage(message []byte) error {
	var record UserInfo
	err := json.Unmarshal(message, &record)
	if err != nil {
		return err
	}

	log.Print(record)

	// Insert Datas
	err = db.collection.Insert(record)
	if err != nil {
		return err
	}

	return nil
}

// GetUserInfo get user information from id and timestamp
func (db *Database) GetUserInfo(id string, timestamp int64) ([]UserInfo, error) {
	result := []UserInfo{}
	err := db.collection.Find(bson.M{"id": id, "time": bson.M{"$gte": timestamp}}).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
