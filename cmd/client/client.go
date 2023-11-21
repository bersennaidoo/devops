package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	caBytes, err := ioutil.ReadFile("./documentation/ca/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	ca := x509.NewCertPool()
	if !ca.AppendCertsFromPEM(caBytes) {
		log.Fatal("ca.cert not valid")
	}

	cert, err := tls.LoadX509KeyPair("documentation/cert/client.crt", "documentation/cert/client.key")
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      ca,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	resp, err := client.Get("https://go-demo.localtest.me:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Printf("Body (status %d): %s\n", resp.StatusCode, body)
}
