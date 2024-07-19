package main

import (
	"github.com/jesses-code-adventures/jort-url/server"
	"net/http"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	defer server.Close()
	err = http.ListenAndServe(":8080", server.Mux)
	if err != nil {
		panic(err)
	}
}
