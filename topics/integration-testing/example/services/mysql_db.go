package services

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const DBUserVar = "SUPERSECRET_DB_USER"
const DBPasswordVar = "SUPERSECRET_DB_PASSWORD"
const DBHostVar = "SUPERSECRET_DB_HOST"

var dsn string

func init() {
	if !useFakeDB {
		dbUser := os.Getenv(DBUserVar)
		if dbUser == "" {
			panic(DBUserVar + " must be set")
		}

		dbPassword := os.Getenv(DBPasswordVar)
		if dbPassword == "" {
			panic(DBPasswordVar + " must be set")
		}

		dbHost := os.Getenv(DBHostVar)
		if dbHost == "" {
			dbHost = "localhost"
		}

		dsn = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":3306)" + "/supersecret_db"

		// Create the DB skeleton if necessary
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Print(err.Error())
			panic("failed to create tables")
		} else {
			_, err := db.Exec(tablesQuery)
			if err != nil {
				log.Print(err.Error())
				panic("failed to create database tables")
			}
		}

	}
}

const tablesQuery = `
	CREATE TABLE IF NOT EXISTS messages (
		message_id INT PRIMARY KEY AUTO_INCREMENT,
		message TEXT
	);
`

const storeQuery = `
	INSERT INTO messages SET message = ?
	RETURNING message_id
`

const fetchQuery = `
	SELECT message FROM messages
	WHERE message_id = ?
`

func realStoreMessage(message string) (string, error) {
	log.Print("using real db store")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Print("could not connect to message db to store message")
		log.Print(err.Error())
		return "", err
	}

	row := db.QueryRow(storeQuery, message)

	var messageID int
	err = row.Scan(&messageID)
	if err != nil {
		log.Print("bad response from database on store message")
		log.Print(err.Error())
		return "", err
	}

	return strconv.Itoa(messageID), nil
}

func realFetchMessage(messageID string) (string, error) {
	log.Print("using real db fetch")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Print("could not connect to message db to fetch message")
		log.Print(err.Error())
		return "", err
	}

	log.Print("using message id: " + messageID)

	intMessageId, err := strconv.Atoi(messageID)
	log.Printf("using int message id: %d", intMessageId)
	if err != nil {
		log.Printf("invalid message id '%s'", messageID)
		return "", err
	}

	row := db.QueryRow(fetchQuery, intMessageId)

	var message string
	err = row.Scan(&message)
	if err != nil {
		log.Print("bad response from database on fetch message")
		log.Print(err.Error())
		return "", err
	}

	return message, nil
}
