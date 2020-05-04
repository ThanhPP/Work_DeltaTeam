package config

import (
	"testing"
)

// func TestGetEnv(t *testing.T) {
// 	value, err := GetEnvKey("TEST")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if value != "TESTTEST" {
// 		t.Errorf("Cannot find the value %+v", value)
// 	}
// }

func TestGetConfig(t *testing.T) {
	Init()
	config := GetConfig()
	val := config.GetString("TEST")
	if val != "TESTTEST" {
		t.Error("Can not read Test config")
	}
}
