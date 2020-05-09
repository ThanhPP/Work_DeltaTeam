package config

import (
	"log"
	"runtime"

	"github.com/spf13/viper"
)

var (
	config     *viper.Viper
	rbConfig   *viper.Viper
	nameConfig *viper.Viper
	ggsConfig  *viper.Viper
)

//Init read the config file
func Init() {
	teleConfigInit()
	rbConfigInit()
	nameConfigInit()
	ggsConfigInit()
}

//------------------------------------------------------------------------------------------------------------

//TelebotConfig for init telegram bot
type TelebotConfig struct {
	apiKey       string
	updateOffset int
	timeOut      int
}

func teleConfigInit() {
	config = viper.New()
	config.SetConfigName("telebot_secret")
	config.SetConfigType("env")

	config.AddConfigPath(".")
	config.AddConfigPath("../../config_file/")
	config.AddConfigPath("../config_file/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

//GetConfig return the Viper to read config from file
func GetConfig() *viper.Viper {
	return config
}

//GetTeleConfigObj return object contains telegram bot configs
func GetTeleConfigObj() *TelebotConfig {
	tlConfig := GetConfig()

	tlCfObj := &TelebotConfig{
		apiKey:       tlConfig.GetString("TELEGRAMBOTAPIKEY"),
		updateOffset: tlConfig.GetInt("TELEOFFSET"),
		timeOut:      tlConfig.GetInt("TELETIMEOUT"),
	}

	return tlCfObj
}

//------------------------------------------------------------------------------------------------------------

//RBConfig contains config for rebrandly handler
type RBConfig struct {
	apiKey         string
	domainID       string
	slashTagColumn string
}

func rbConfigInit() {
	rbConfig = viper.New()
	rbConfig.SetConfigName("rb_secret")
	rbConfig.SetConfigType("env")

	rbConfig.AddConfigPath(".")
	rbConfig.AddConfigPath("../../config_file/")
	rbConfig.AddConfigPath("../config_file/")

	if err := rbConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

//GetRBConfig return config for rebrandly
func GetRBConfig() *viper.Viper {
	return rbConfig
}

//GetRBConfigObj return an object contains config for rebrandly
func GetRBConfigObj() *RBConfig {
	rbConfig = GetRBConfig()
	rbCf := &RBConfig{
		apiKey:         rbConfig.GetString("REBRANDLYAPIKEY"),
		domainID:       rbConfig.GetString("REBRANDLYDOMAINID"),
		slashTagColumn: rbConfig.GetString("SLASHTAGCOLUMN"),
	}

	return rbCf
}

//ChangeAPIKey change the API key of rb and write back to file
func (rbcf *RBConfig) ChangeAPIKey(newAPIKey string) error {
	rbConfig = GetRBConfig()

	backupAPIKey := rbcf.apiKey

	rbcf.apiKey = newAPIKey

	rbConfig.Set("REBRANDLYAPIKEY", rbcf.apiKey)

	if err := rbConfig.WriteConfig(); err != nil {
		rbcf.apiKey = backupAPIKey
		rbConfig.Set("REBRANDLYAPIKEY", rbcf.apiKey)
		return err
	}

	return nil
}

//------------------------------------------------------------------------------------------------------------

//NameDotConfig name.com config object
type NameDotConfig struct {
	domain                string
	username              string
	apiKey                string
	storeLinkColumn       string
	tempForwardLinkColumn string
}

func nameConfigInit() {
	nameConfig = viper.New()
	nameConfig.SetConfigName("namedotcom_secret")
	nameConfig.SetConfigType("env")

	nameConfig.AddConfigPath(".")
	nameConfig.AddConfigPath("../../config_file/")
	nameConfig.AddConfigPath("../config_file/")

	if err := nameConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func getNameConfig() *viper.Viper {
	return nameConfig
}

//GetNameConfigObj an object contains config for rebrandly
func GetNameConfigObj() *NameDotConfig {
	nameCf := getNameConfig()

	nameComCfObj := &NameDotConfig{
		domain:                nameCf.GetString("NAMEDOMAIN"),
		username:              nameCf.GetString("NAMEUSRNAME"),
		apiKey:                nameCf.GetString("NAMEAPIKEY"),
		storeLinkColumn:       nameCf.GetString("STORELINKCOLUMN"),
		tempForwardLinkColumn: nameCf.GetString("TEMPFORWARDLINKCOLUMN"),
	}

	return nameComCfObj
}

//------------------------------------------------------------------------------------------------------------

//GGSConfig google sheet config
type GGSConfig struct {
	Path          string
	SpreadSheetID string
	TableName     string
}

func ggsConfigInit() {
	ggsConfig = viper.New()
	ggsConfig.SetConfigName("googlesheet_secret")
	ggsConfig.SetConfigType("env")

	ggsConfig.AddConfigPath(".")
	ggsConfig.AddConfigPath("../../config_file/")
	ggsConfig.AddConfigPath("../config_file/")

	if err := ggsConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func getggsConfig() *viper.Viper {
	return ggsConfig
}

//GetGGSConfigObj object contains google sheet config
func GetGGSConfigObj() *GGSConfig {
	ggsCf := getggsConfig()

	//choose path
	var pathOS string
	if runtime.GOOS == "windows" {
		pathOS = ggsCf.GetString("WINDOWSGGSSCRPATH")
	} else {
		pathOS = ggsCf.GetString("LINUXGGSSCRPATH")
	}

	ggsCfObj := &GGSConfig{
		Path:          pathOS,
		SpreadSheetID: ggsCf.GetString("SPREADSHEETID"),
		TableName:     ggsCf.GetString("TABLENAME"),
	}

	return ggsCfObj
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
