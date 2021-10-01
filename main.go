package main

import (
	"github.com/takeshiemoto/golang-test-driven-development/di"
	"net/http"
	"os"
)

func main() {
	di.Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(di.MyGreeterHandler))
}
