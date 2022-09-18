package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	wd, _ := os.Getwd()
	certFile := path.Join(wd, "/ca/certificate.pem")
	keyFile := path.Join(wd, "/ca/private.pem")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(strings.Repeat("hello ", 100) + "world, " + r.Proto))
	})

	h1s := &http.Server{
		Addr:    ":18080",
		Handler: h2c.NewHandler(handler, new(http2.Server)),
	}
	log.Fatal(h1s.ListenAndServeTLS(certFile, keyFile))
}
