package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "it's working")
}
func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServeTLS(":8443", "documentation/cert/server.crt", "documentation/cert/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS error: ", err)
	}
	fmt.Println("Listening on PORT: 8443")
}
