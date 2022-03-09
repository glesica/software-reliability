package services

import (
	"fmt"
	"log"
	"strconv"
)

var fakeMessages = []string{
	"test message",
}

func fakeStoreMessage(message string) (string, error) {
	log.Print("using fake db store")
	fakeMessages = append(fakeMessages, message)
	return fmt.Sprintf("%d", len(fakeMessages)-1), nil
}

func fakeFetchMessage(messageID string) (string, error) {
	log.Print("using fake db store")
	index, err := strconv.Atoi(messageID)
	if err != nil {
		return "", err
	}

	return fakeMessages[index], nil
}
