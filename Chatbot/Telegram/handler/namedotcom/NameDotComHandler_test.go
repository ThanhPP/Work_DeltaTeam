package namedotcom

import (
	"testing"

	"github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
)

// var (
// 	GetDataFromEnv = getDataFromEnv
// )

// func TestGetDataFromEnv(t *testing.T) {
// 	nameDomain, _, _ := GetDataFromEnv()
// 	if nameDomain != "deltapodvn.com" {
// 		t.Error("GET NOTHING")
// 	}
// }

func TestForwardLinks(t *testing.T) {
	config.Init()
	_, _, errorForwardCount := ForwardLinks("3934:3935")
	if errorForwardCount != 2 {
		t.Errorf("TestForwardLinks create %+v instead of 2", errorForwardCount)
	}
}
