package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ----------
// Encryption
// ----------

func EncryptMessage(message, secret string) (string, error) {
	return fakeEncryptMessage(message, secret)
}

func fakeEncryptMessage(message, secret string) (string, error) {
	return "encrypted: " + message, nil
}

func realEncryptMessage(message, secret string) (string, error) {
	endpoint, err := url.Parse("http://localhost:8091/encrypt")
	if err != nil {
		return "", err
	}

	endpoint.Query().Add("message", message)
	endpoint.Query().Add("secret", secret)

	res, err := http.Get(endpoint.String())
	if err != nil {
		return "", err
	}

	var buffer []byte
	_, err = res.Body.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func DecryptMessage(message, secret string) (string, error) {
	return fakeDecryptMessage(message, secret)
}

func fakeDecryptMessage(message, secret string) (string, error) {
	return "decrypted: " + message, nil
}

func realDecryptMessage(message, secret string) (string, error) {
	endpoint, err := url.Parse("http://localhost:8091/decrypt")
	if err != nil {
		return "", err
	}

	endpoint.Query().Add("message", message)
	endpoint.Query().Add("secret", secret)

	res, err := http.Get(endpoint.String())
	if err != nil {
		return "", err
	}

	var buffer []byte
	_, err = res.Body.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

// --------
// Messages
// --------

var fakeMessages []string

func StoreMessage(message string) (string, error) {
	return fakeStoreMessage(message)
}

func fakeStoreMessage(message string) (string, error) {
	fakeMessages = append(fakeMessages, message)
	return fmt.Sprintf("%d", len(fakeMessages)-1), nil
}

func FetchMessage(id string) (string, error) {
	return fakeFetchMessage(id)
}

func fakeFetchMessage(id string) (string, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	return fakeMessages[index], nil
}
