package config

import (
	"testing"
)

func TestGetConfigFile(t *testing.T) {
	x := GetConfigFile("../../config.yaml")
	if x == nil {
		t.Error("config file is nil")
	}
}
