package mydb

import (
	"mainpkg/config"
	"testing"
)

func TestCrossing(t *testing.T) {
	confFile := config.GetConfigFile("../../config.yaml")
	dbData := Init(confFile.MongodbURI)

	result := dbData.GetOrCreateUser("admin@admin.pl")

	if result.IsZero() {
		t.Error("Can't find admin@admin.pl")
	}
}
