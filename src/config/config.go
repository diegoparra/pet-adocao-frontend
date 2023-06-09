package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.fatal(err)
	}

	apiurl = os.getenv("api_url")

	hashkey = []byte(os.getenv("hash_key"))

	blockkey = []byte(os.getenv("block_key"))

}
