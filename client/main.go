package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/http2"
)

func main() {
	wd, _ := os.Getwd()

	certFile, _ := ioutil.ReadFile(path.Join(wd, "/ca/certificate.pem"))
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(certFile)

	client := new(http.Client)
	client.Transport = &http2.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:            caCertPool,
			InsecureSkipVerify: true,
		},
	}

	rsp, err := client.Get("https://localhost:18080")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
