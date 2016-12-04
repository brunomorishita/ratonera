package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// "github.com/brunomorishita/ratonera/server/connection"
	"./connection"
)

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

type Configuration struct {
	mongo string
	port  int
}

func main() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	conn := connection.NewConnection(configuration.mongo)
	// serveSingle("/", "index.html")
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/raton", conn.HandleWebsocket)
	http.HandleFunc("/getuserinfo", conn.GetUserInfo)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configuration.port), nil))
}
