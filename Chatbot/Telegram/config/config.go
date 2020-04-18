package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	linuxEnvPath   = "/home/tpp18/go/src/github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/telebot_secret.env"
	windowsEnvPath = "C:/Go/src/github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/telebot_secret.env"
)

func getEnvPath() string {
	log.Println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		return windowsEnvPath
	}
	fmt.Println(windowsEnvPath)
	return linuxEnvPath
}

func loadEnvFile(fileName string) error {
	err := godotenv.Load(fileName)
	return err
}

//GetEnvKey look up value of a key in env file
func GetEnvKey(key string) (string, error) {
	log.Println(runtime.GOOS)
	log.Println(getEnvPath())
	loadEnvFile(getEnvPath())
	value, exist := os.LookupEnv(key)
	if !exist {
		err := errors.New("No value for " + string(key) + "found")
		return value, err
	}
	return value, nil
}
