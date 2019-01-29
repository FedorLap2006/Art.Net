package main

import (
	"net/http"

	wbf "github.com/FedorLap2006/WebFrame"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api", wbf.HandleHTTP(func(c *wbf.Context) {
		c.RespWriter.Write([]byte("hello world WBF!"))
	}))
	http.ListenAndServe(":8080", router)
}
