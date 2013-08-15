package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhellName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayhellName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}
func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
