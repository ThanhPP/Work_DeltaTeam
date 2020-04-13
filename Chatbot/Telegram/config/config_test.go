package config

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	value, err := GetEnvKey("TEST")
	if err != nil {
		t.Error(err)
	}
	if value != "TESTTEST" {
		t.Errorf("Cannot find the value %+v", value)
	}
}
