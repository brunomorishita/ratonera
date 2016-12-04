package connection

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
