package main

import (
	"net/http"
	"fmt"
	"strconv"

	"github.com/cloudfoundry-community/go-cfenv"
)

func main() {
	appEnv, _ := cfenv.Current()
	port := strconv.Itoa(appEnv.Port)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}

func hello(res http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(res, "Hello world!")
}