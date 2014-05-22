package main

import (
	"log"
	"net/http"
	"github.com/lytics/httppprof"
)

func main() {
	httppprof.Register(http.DefaultServeMux)
	log.Println(http.ListenAndServe("localhost:6060", nil))
}
