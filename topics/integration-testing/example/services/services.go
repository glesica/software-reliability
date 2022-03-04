package services

// --------
// Services
// --------

func EncryptMessage(message, secret string) (string, error) {
	return fakeEncryptMessage(message, secret)
}

func fakeEncryptMessage(message, secret string) (string, error) {
	return "encrypted: " + message, nil
}

func realEncryptMessage(message, secret string) (string, error) {
	return "", nil
}

func DecryptMessage(message, secret string) (string, error) {
	return fakeDecryptMessage(message, secret)
}

func fakeDecryptMessage(message, secret string) (string, error) {
	return "decrypted: " + message, nil
}

func StoreMessage(message, id string) error {
	return fakeStoreMessage(message, id)
}

func fakeStoreMessage(message, id string) error {
	return nil
}

func FetchMessage(id string) (string, error) {
	return fakeFetchMessage(id)
}

func fakeFetchMessage(id string) (string, error) {
	return "this is a big secret", nil
}
