package config

import (
	"errors"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	linuxEnvPath   = "/home/tpp18/go/src/github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/telebot_secret.env"
	windowsEnvPath = ""
)

func getEnvPath() string {
	if runtime.GOOS == "windows" {
		return windowsEnvPath
	}

	return linuxEnvPath
}

func loadEnvFile(fileName string) error {
	err := godotenv.Load(fileName)
	return err
}

//GetEnvKey look up value of a key in env file
func GetEnvKey(key string) (string, error) {
	loadEnvFile(getEnvPath())
	value, exist := os.LookupEnv(key)
	if !exist {
		err := errors.New("No value for " + string(key) + "found")
		return value, err
	}
	return value, nil
}
