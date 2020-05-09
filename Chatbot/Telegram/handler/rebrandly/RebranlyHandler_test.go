package rebrandly

import (
	"testing"

	"github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
)

// func TestGetDataFromEnv(t *testing.T) {
// 	config.Init()
// 	_, domainID := getDataFromEnv()
// 	if domainID == "" {
// 		t.Error("Can not get domain")
// 	}
// }

func TestCreateShortLinkRebrandly(t *testing.T) {
	config.Init()
	inputFwdSlice := []string{"1704fm004pxrac.deltapodvn.com"}
	_, _, errCount, _ := CreateShortLinkRebrandly("4031:4031", inputFwdSlice)
	if errCount != 1 {
		t.Error(errCount)
	}
}

func TestSetRebrandlyAPIKey(t *testing.T) {
	config.Init()
	err := SetRebrandlyAPIKey("d2b46b231a55436db161d83e8774f3b9")
	if err != nil {
		t.Error(err)
	}
	cf := config.GetRBConfig()
	val := cf.GetString("REBRANDLYAPIKEY")
	if val != "d2b46b231a55436db161d83e8774f3b9" {
		t.Error(val)
	}
}
