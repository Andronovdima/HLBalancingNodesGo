package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"time"
)

func main()  {
	var port string
	var timeout int64

	flag.StringVar(&port, "port", ":8080", "usage port")
	flag.Int64Var(&timeout, "timeout" , 200, "timeout set")
	flag.Parse()

	timeoutD := time.Duration(timeout)

	prometheus.MustRegister(RequestCount)

	srv := NewServer(port)
	NewMainHandler(srv.Mux, port, timeoutD)
	fmt.Println("starting server on port" + port)
	if err := http.ListenAndServe(port, srv); err != nil {
		log.Fatal(err)
	}

}

