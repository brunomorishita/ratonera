package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/brunomorishita/ratonera/server/connection"
	//"./connection"
)

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

var mongo = flag.String("mongo", "127.0.0.1", "Mongo Server")

type Configuration struct {
	mongo string
	port  int
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	conn := connection.NewConnection(*mongo)
	// serveSingle("/", "index.html")
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/raton", conn.HandleWebsocket)
	http.HandleFunc("/getuserinfo", conn.GetUserInfo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
