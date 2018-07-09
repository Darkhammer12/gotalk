package main

import (
	"fmt"
	"gotalk/glogger"
	"net/http"
)

func main() {
	logger := glogger.GetInstance()
	logger.Println("Starting gotalk web service")

	http.HandelFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := glogger.GetInstance()
	fmt.Fprint(w, "Wlecome to the gotalk software system")

	logger.Println("Recieved an http request on root url")
}
