package main

import (
	"log"
	"net/http"
)

func encrypt(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	// 1. Get the message and secret from params
	message, exists := req.URL.Query()["message"]
	if !exists {
		log.Print("missing message query param on encrypt")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	secret, exists := req.URL.Query()["secret"]
	if !exists {
		log.Print("missing secret query param on encrypt")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// 2. Encrypt the message with the secret
	encryptedMessage := XorMessage(message[0], secret[0])

	// 3. Return the encrypted message to the caller
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(encryptedMessage))
}

func decrypt(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	// 1. Get the encrypted message and secret from params
	message, exists := req.URL.Query()["message"]
	if !exists {
		log.Print("missing message query param on decrypt")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	secret, exists := req.URL.Query()["secret"]
	if !exists {
		log.Print("missing secret query param on decrypt")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// 2. Decrypt the message with the secret
	decryptedMessage := XorMessage(message[0], secret[0])

	// 3. Return the message to the caller
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(decryptedMessage))
}

func main() {
	http.HandleFunc("/encrypt", encrypt)
	http.HandleFunc("/decrypt", decrypt)

	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		panic("failed to start encrypt server")
	}
}
