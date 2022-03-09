package main

// XorMessage "encrypts" the message using the given secret.
// Note: this is a really dumb way to encrypt data, don't use it
// for anything even mildly important!
func XorMessage(message, secret string) string {
	messageBytes := []byte(message)
	secretBytes := []byte(secret)

	secretLength := len(secretBytes)
	nextSecretIndex := 0
	for i := 0; i < len(messageBytes); i++ {
		messageBytes[i] = messageBytes[i] ^ secretBytes[nextSecretIndex]
		nextSecretIndex = (nextSecretIndex + 1) % secretLength
	}

	return string(messageBytes)
}
