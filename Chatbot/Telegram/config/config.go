package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

//Init read the config file
func Init() {
	config = viper.New()
	config.SetConfigName("telebot_secret")
	config.SetConfigType("env")

	config.AddConfigPath(".")
	config.AddConfigPath("../../config/")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

//GetConfig return the Viper to read config from file
func GetConfig() *viper.Viper {
	return config
}

// var (
// 	linuxEnvPath   = "/home/tpp18/go/src/github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/telebot_secret.env"
// 	windowsEnvPath = "C:/Go/src/github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/telebot_secret.env"
// )

// func getEnvPath() string {
// 	if runtime.GOOS == "windows" {
// 		return windowsEnvPath
// 	}
// 	return linuxEnvPath
// }

// func loadEnvFile(fileName string) error {
// 	err := godotenv.Load(fileName)
// 	return err
// }

// //GetEnvKey look up value of a key in env file
// func GetEnvKey(key string) (string, error) {
// 	loadEnvFile(getEnvPath())
// 	value, exist := os.LookupEnv(key)
// 	if !exist {
// 		err := errors.New("No value for " + string(key) + "found")
// 		return value, err
// 	}
// 	return value, nil
// }
