package services

import "os"

const FakeEncryptVar = "SUPERSECRET_FAKE_ENCRYPT"
const FakeDBVar = "SUPERSECRET_FAKE_DB"

// ----------
// Encryption
// ----------

var useFakeEncrypt = os.Getenv(FakeEncryptVar) == "true"

func EncryptMessage(message, secret string) (string, error) {
	if useFakeEncrypt {
		return fakeEncryptMessage(message, secret)
	} else {
		return realEncryptMessage(message, secret)
	}
}

func DecryptMessage(message, secret string) (string, error) {
	if useFakeEncrypt {
		return fakeDecryptMessage(message, secret)
	} else {
		return realDecryptMessage(message, secret)
	}
}

// --------
// Messages
// --------

var useFakeDB = os.Getenv(FakeDBVar) == "true"

func StoreMessage(message string) (string, error) {
	if useFakeDB {
		return fakeStoreMessage(message)
	} else {
		return realStoreMessage(message)
	}
}

func FetchMessage(messageID string) (string, error) {
	if useFakeDB {
		return fakeFetchMessage(messageID)
	} else {
		return realFetchMessage(messageID)
	}
}
