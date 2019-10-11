package main

import (
	"fmt"
	"net/http"
)

type cameConnect int

func (m cameConnect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("do 1")
}

func main() {
	var d cameConnect
	http.ListenAndServe(":8080", d)
}
