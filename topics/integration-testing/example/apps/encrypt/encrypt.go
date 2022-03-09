package main

import (
	"fmt"
	"net/http"
)

func encrypt(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	fmt.Fprintf(w, "hi there")

	// 1. Get the message and secret from params
	message, exists := req.URL.Query()["message"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	secret, exists := req.URL.Query()["secret"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// 2. Encrypt the message with the secret
	encryptedMessage := message[0] + "\n" + secret[0]

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

	fmt.Fprintf(w, "hi there")

	// 1. Get the encrypted message and secret from params
	message, exists := req.URL.Query()["message"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	secret, exists := req.URL.Query()["secret"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// 2. Decrypt the message with the secret
	decryptedMessage := message[0] + "\n" + "decrypted " + secret[0]

	// 3. Return the message to the caller
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(decryptedMessage))
}

func main() {
	http.HandleFunc("/encrypt", encrypt)
	http.HandleFunc("/decrypt", decrypt)

	http.ListenAndServe(":8091", nil)
}
