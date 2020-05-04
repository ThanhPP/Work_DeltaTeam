package rebrandly

import (
	"testing"

	"github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
)

func TestGetDataFromEnv(t *testing.T) {
	config.Init()
	_, domainID := getDataFromEnv()
	if domainID == "" {
		t.Error("Can not get domain")
	}
}

func TestCreateShortLinkRebrandly(t *testing.T) {
	config.Init()
	inputFwdSlice := []string{"1704fm004pxrac.deltapodvn.com"}
	_, _, errCount, _ := CreateShortLinkRebrandly("4031:4031", inputFwdSlice)
	if errCount != 1 {
		t.Error(errCount)
	}
}
