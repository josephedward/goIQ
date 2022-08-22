package core

import (
	"github.com/joho/godotenv"
	"os"
)

type AwsEnv struct {
	Url          string
	Username     string
	Password     string
	Key_id     string
	Access_key string
}

func LoadEnv() (login AwsEnv, err error) {

	//load env variables
	err = godotenv.Load("./.env")
	
	//set all needed vendor credentials
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	url := os.Getenv("URL")
	key_id := os.Getenv("KEY_ID")
	access_key := os.Getenv("ACCESS_KEY")

	//return credentials
	return AwsEnv{url, username, password, key_id, access_key}, err
}

