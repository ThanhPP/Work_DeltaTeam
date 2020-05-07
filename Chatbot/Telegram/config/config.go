package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	config   *viper.Viper
	rbConfig *viper.Viper
)

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

	rbConfig = viper.New()
	rbConfig.SetConfigName("rb_secret")
	rbConfig.SetConfigType("env")

	rbConfig.AddConfigPath(".")
	rbConfig.AddConfigPath("../../config/")
	rbConfig.AddConfigPath("config/")

	if err := rbConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

//GetConfig return the Viper to read config from file
func GetConfig() *viper.Viper {
	return config
}

//GetRBConfig return config for rebrandly
func GetRBConfig() *viper.Viper {
	return rbConfig
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
