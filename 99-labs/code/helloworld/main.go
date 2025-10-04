package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var hostname string
var version string

func init() {
	h, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = h

	version = runtime.Version()
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from %s running Go version %s!\n", hostname, version)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
