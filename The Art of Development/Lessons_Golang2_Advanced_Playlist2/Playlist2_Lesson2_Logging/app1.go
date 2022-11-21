package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}

func main() {
	router := httprouter.New()
	router.GET("/:name", IndexHandler)

	listener, err := net.Listen("tcp", "127.0.0.1:1234") // 127.0.0.1 - это адрес не текущей машины,
	// а адрес loopback интерфейса "loo"
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second, // Цифры беруться империческим путем
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))

}
