package vmextension

import (
	"io/ioutil"
	"testing"
)

func Test_parseHandlerSettingsFile_Bad(t *testing.T) {
	json := `{"runtimeSettings": [  ]}`
	if _, err := parseHandlerSettingsFile([]byte(json)); err == nil {
		t.Fatal("did not fail")
	}
}

func Test_parseHandlerSettingsFile_Good(t *testing.T) {
	json, err := ioutil.ReadFile("../../testdata/Extension/config/2.settings")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := parseHandlerSettingsFile(json); err != nil {
		t.Fatal(err)
	}
}
