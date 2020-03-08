package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/http2"
)

func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor < 2 {
			w.WriteHeader(http.StatusUpgradeRequired)
			w.Write([]byte(fmt.Sprintf("HTTP/2 is required, but got %d", r.ProtoMajor)))
			return
		}

		in, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error: %+v", err)))
			return
		}
		sb := strings.Builder{}
		sb.Write(in)
		req := sb.String()

		for i := 0; i < 10; i++ {
			if _, err := w.Write([]byte(fmt.Sprintf("%d: Hello %s\n", i, req))); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("failed to write %d, error: %+v", i, err)))
				return
			}
			w.(http.Flusher).Flush()
			<-time.After(200 * time.Millisecond)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	})
	go func() {
		http.ListenAndServeTLS(":8080", "server.crt", "server.key", mux)
	}()
}

func runClient() {
	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	client := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Post("https://127.0.0.1:8080/stream", "text/plain", strings.NewReader("hello"))
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()

	bufferedReader := bufio.NewReader(resp.Body)
	result := strings.Builder{}
	buffer := make([]byte, 4*8)

	// Reads the response
	for {
		length, err := bufferedReader.Read(buffer)
		if length > 0 {
			fmt.Println(length, "bytes received")
			fmt.Println(string(buffer[:length]))
			result.WriteString(string(buffer[:length]))
		}

		if err != nil {
			if err == io.EOF {
				// Last chunk received
			}
			break
		}
	}
	fmt.Printf("Proto: %d, result: %s\n", resp.ProtoMajor, result.String())
	return
}

func main() {
	runServer()
	runClient()
}
