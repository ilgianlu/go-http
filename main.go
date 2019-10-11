package main

import (
	"fmt"
	"net/http"
	"github.com/ilgianlu/go-mw"
)

type cameConnect int

func (m cameConnect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("do 1")
}

func main() {
	var d cameConnect
	http.ListenAndServe(":8080", middleware.LogMiddleware(d.ServeHTTP))
}
