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

// func LoadEnv() (login AwsEnv, err error) {
// 	//load env variables
// 	err = godotenv.Load("./.env")

// 	//set all needed vendor credentials
// 	username := os.Getenv("USERNAME")
// 	password := os.Getenv("PASSWORD")
// 	url := os.Getenv("URL")
// 	key_id := os.Getenv("KEY_ID")
// 	access_key := os.Getenv("ACCESS_KEY")

// 	//return credentials
// 	return AwsEnv{url, username, password, key_id, access_key}, err
// }

func Env() (login AwsEnv){
	
	//set all needed vendor credentials
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	url := os.Getenv("URL")
	key_id := os.Getenv("KEY_ID")
	access_key := os.Getenv("ACCESS_KEY")

	return AwsEnv{url, username, password, key_id, access_key}
}

func LoadEnv() (login AwsEnv, err error) {
	//load env variables
	err = godotenv.Load("./.env")
	env := Env()
	return env, err
}

func LoadEnvPath(path string) (login AwsEnv, err error) {
	//load env variables
	err = godotenv.Load(path)
	env := Env()
	return env, err
}

func ArgEnv() (login AwsEnv, err error) {
	//set all needed vendor credentials from cli arg
	url := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	key_id := os.Args[4]
	access_key := os.Args[5]

	return AwsEnv{url, username, password, key_id, access_key}, err
}

