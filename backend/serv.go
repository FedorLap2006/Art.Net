package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/graarh/golang-socketio/transport"

	"github.com/FedorLap2006/Art.Net/utils"
	// ws "golang.org/x/net/websocket"
	sio "github.com/graarh/golang-socketio"

	// "C:/Users/Федор/Desktop/Github/Art.Net/utils"

	wbf "github.com/FedorLap2006/WebFrame"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// db params
const (
	UserDB  = "develop"
	PassDB  = ""
	HostDB  = "127.0.0.1"
	PortDB  = "3306"
	LabelDB = "Art_Net"
)

func main() {
	// utils.Database = sql.Open("pg")
	var err error
	sioServer := sio.NewServer(transport.GetDefaultWebsocketTransport())
	utils.SIOServer = sioServer

	utils.Database, _ = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", UserDB, PassDB, HostDB, PortDB, LabelDB))
	err = utils.Database.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connect to database successfully")
	}

	router := mux.NewRouter()
	router.HandleFunc("/api", wbf.HandleHTTP(func(c *wbf.Context) {
		c.RespWriter.Write([]byte("hello world WBF!"))

	}))
	router.Handle("/api/ws", sioServer)
	http.ListenAndServe(":8080", router)
}
