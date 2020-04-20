package namedotcom

import (
	"testing"
)

var (
	GetDataFromEnv = getDataFromEnv
)

func TestGetDataFromEnv(t *testing.T) {
	nameDomain, _, _ := GetDataFromEnv()
	if nameDomain != "deltapodvn.com" {
		t.Error("GET NOTHING")
	}
}

func TestForwardLinks(t *testing.T) {
	_, _, errorForwardCount := ForwardLinks("3934:3935")
	if errorForwardCount != 2 {
		t.Errorf("TestForwardLinks create %+v instead of 2", errorForwardCount)
	}
}
