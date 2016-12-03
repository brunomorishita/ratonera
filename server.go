// +build ignore

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/brunomorishita/ratonera"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	conn := ratonera.NewConnection()
	http.HandleFunc("/raton", conn.HandleWebsocket)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
