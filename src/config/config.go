package config

import (
	"log"
	"os"
	"strconv"
)

var (
	// APIURL to connect to api
	APIURL = ""
	// Port where our web app is running
	Port = 0
	// HashKey is utilized to authenticate the cookie
	HashKey []byte
	// BlockKey is utilized to crypt the cookie data
	BlockKey []byte
)

// Load initialize the env variables
func Load() {
	var err error

	// if err = godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")

	HashKey = []byte(os.Getenv("HASH_KEY"))

	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
