package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"supersecret/services"
	"time"
)

//go:embed templates/index.html
var indexTemplateContent string

//go:embed templates/view.html
var viewTemplateContent string

type ViewParams struct {
	ID string
}

//go:embed templates/read.html
var readTemplateContent string

type ReadParams struct {
	ID      string
	Message string
}

// -------------
// Page Handlers
// -------------

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	if req.Method == http.MethodGet {
		templ, err := template.New("index-page").Parse(indexTemplateContent)
		if err != nil {
			log.Println("Failed to parse index-page template")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		var buffer bytes.Buffer
		err = templ.Execute(&buffer, nil)
		if err != nil {
			log.Println("Failed to execute index-page template")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		w.Write(buffer.Bytes())
	}

	if req.Method == http.MethodPost {
		req.ParseForm()

		// 1. Send message and secret to encryption service
		message := req.Form.Get("message")
		secret := req.Form.Get("secret")
		if secret == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			return
		}

		encryptedMessage, err := services.EncryptMessage(message, secret)
		if err != nil {
			log.Println("failed to encrypt message")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		// 2. Generate ID for message
		messageID := "fa8723fh2dh7d"

		// 3. Store encrypted message and ID in message DB
		err = services.StoreMessage(encryptedMessage, messageID)
		if err != nil {
			log.Println("failed to encrypt message")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		// 4. Redirect user to message view
		showView(w, messageID)
	}
}

func view(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	if req.Method == http.MethodGet {
		// On get:
		// 1. Get the message ID from the URL
		messageID, exists := req.URL.Query()["id"]
		if !exists {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			return
		}

		// 2. Populate it in the secret field
		showView(w, messageID[0])
	}

	if req.Method == http.MethodPost {
		req.ParseForm()

		// On post:
		// 1. Look up the message in the message DB
		messageID := req.Form.Get("id")
		if messageID == "" {
			log.Println("id not found in view request")
			log.Printf("form contained %s\n", req.Form)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			return
		}

		secret := req.Form.Get("secret")
		if secret == "" {
			log.Println("secret not found in view request")
			log.Printf("form contained %s\n", req.Form)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			return
		}

		encryptedMessage, err := services.FetchMessage(messageID)
		if err != nil {
			log.Println("Failed to parse view-page template")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		// 2. Send message and secret to encryption service
		message, err := services.DecryptMessage(encryptedMessage, secret)
		if err != nil {
			log.Println("Failed to parse view-page template")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		// 3. Send message back to user
		templ, err := template.New("read-page").Parse(readTemplateContent)
		if err != nil {
			log.Println("Failed to parse read-page template")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		var buffer bytes.Buffer
		err = templ.Execute(&buffer, ReadParams{ID: messageID, Message: message})
		if err != nil {
			log.Println("Failed to execute read-page template")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		w.Write(buffer.Bytes())
	}
}

func showView(w http.ResponseWriter, messageID string) {
	templ, err := template.New("view-page").Parse(viewTemplateContent)
	if err != nil {
		log.Println("Failed to parse view-page template")
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	var buffer bytes.Buffer
	err = templ.Execute(&buffer, ViewParams{ID: messageID})
	if err != nil {
		log.Println("Failed to execute view-page template")
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.Write(buffer.Bytes())
}

func main() {
	http.HandleFunc("/favicon.png", func(w http.ResponseWriter, req *http.Request) {
		image, err := os.Open("images/favicon.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(http.StatusText(http.StatusNotFound)))
			return
		}

		http.ServeContent(w, req, "images/favicon.png", time.UnixMicro(0), image)
	})

	http.HandleFunc("/", index)
	http.HandleFunc("/view", view)

	http.ListenAndServe(":8090", nil)
}
