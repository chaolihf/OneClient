package http

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

var server *http.Server

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-PINGOTHER,Content-Type")
	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	fmt.Fprintln(w, "Hello, HTTPS!")
}

func StartSecurityServer() {
	http.HandleFunc("/", handler)
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12, // 最低支持TLS 1.2
	}
	server = &http.Server{
		Addr:      "localhost:443",
		TLSConfig: tlsConfig,
	}
	log.Println("Starting HTTPS server on port 443...")
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer() {
	http.HandleFunc("/", handler)
	server = &http.Server{
		Addr: "localhost:8080",
	}
	log.Println("Starting HTTP server on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func StopServer() {
	server.Close()
}
