// +build ignore

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/brunomorishita/ratonera/server"
)

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

var addr = flag.String("addr", "192.168.1.6:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	conn := ratonera.NewConnection()
	// serveSingle("/", "index.html")
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/raton", conn.HandleWebsocket)
	http.HandleFunc("/getuserinfo", conn.GetUserInfo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
