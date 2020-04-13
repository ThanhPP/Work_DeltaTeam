package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func loadEnvFile(fileName string) error {
	err := godotenv.Load(fileName)
	return err
}

//GetEnvKey look up value of a key in env file
func GetEnvKey(key string) (string, error) {
	loadEnvFile("/home/tpp18/go/src/github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/telebot_secret.env")
	value, exist := os.LookupEnv(key)
	if !exist {
		err := errors.New("No value for " + string(key) + "found")
		return value, err
	}
	return value, nil
}
