package services

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const EncryptHostVar = "SUPERSECRET_ENCRYPT_HOST"

var encryptHost string

func init() {
	encryptHost = os.Getenv(EncryptHostVar)
	if encryptHost == "" {
		encryptHost = "localhost"
	}
}

func fakeEncryptMessage(message, secret string) (string, error) {
	log.Print("using fake encrypt")
	return "encrypted: " + message, nil
}

func realEncryptMessage(message, secret string) (string, error) {
	log.Print("using real encrypt")
	endpoint, err := url.Parse("http://" + encryptHost + ":8091/encrypt")
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("message", message)
	params.Add("secret", secret)
	endpoint.RawQuery = params.Encode()

	res, err := http.Get(endpoint.String())
	if err != nil {
		log.Print("encrypt message request failed")
		log.Print(err.Error())
		return "", err
	}

	if res.StatusCode != 200 {
		log.Print("bad response from encrypt service on encrypt")
		log.Print(res.Status)
		return "", fmt.Errorf("bad response on encrypt: %d", res.StatusCode)
	}

	log.Printf("content length is: %d", res.ContentLength)

	defer res.Body.Close()
	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print("could not read encrypt message response")
		log.Print(err.Error())
		return "", err
	}

	return string(buffer), nil
}

func fakeDecryptMessage(message, secret string) (string, error) {
	log.Print("using fake decrypt")
	return "decrypted: " + message, nil
}

func realDecryptMessage(message, secret string) (string, error) {
	log.Print("using real decrypt")
	endpoint, err := url.Parse("http://" + encryptHost + ":8091/decrypt")
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("message", message)
	params.Add("secret", secret)
	endpoint.RawQuery = params.Encode()

	res, err := http.Get(endpoint.String())
	if err != nil {
		log.Print("decrypt message request failed")
		log.Print(err.Error())
		return "", err
	}

	if res.StatusCode != 200 {
		log.Print("bad response from encrypt service on decrypt")
		log.Print(res.Status)
		return "", fmt.Errorf("bad response on decrypt: %d", res.StatusCode)
	}

	defer res.Body.Close()
	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print("could not read decrypt message response")
		log.Print(err.Error())
		return "", err
	}

	return string(buffer), nil
}
