package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"random-password-generator/src/handler"
	"random-password-generator/src/param"
)

func main() {
	port := os.Getenv(param.Port)
	if port == "" {
		log.Fatal("Port not set")
		os.Exit(1)
	}
	privateKeyPath := os.Getenv(param.PrivateKey)
	if privateKeyPath == "" {
		log.Fatal("Private key path not set")
		os.Exit(1)
	}
	publicKeyPath := os.Getenv(param.PublicKey)
	if publicKeyPath == "" {
		log.Fatal("Public key path not set")
		os.Exit(1)
	}

	http.HandleFunc("/password", handler.Password)
	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServeTLS(
			fmt.Sprintf(":%s", port),
			publicKeyPath,
			privateKeyPath,
			nil,
		)
	}()
	log.Fatal(<-errChan)
}
