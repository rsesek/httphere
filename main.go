package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

var (
	port = flag.String("port", "0", "Port to run the server. 0 for a random port.")
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		print(err)
		os.Exit(1)
	}

	server := http.FileServer(http.Dir(pwd))

	addr := net.JoinHostPort("", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		print(err)
		os.Exit(2)
	}

	fmt.Printf("Listening on %s\n", listener.Addr().String())
	fmt.Printf("Serving conents of %s\n", pwd)
	http.Serve(listener, server)
}
