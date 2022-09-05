package config

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	yamlpath := "./config.example.yaml"

	want := Config{Version: "Test", Blockchain: struct {
		Token string "yaml:\"token\""
	}{Token: "abcdefghijklmnoprstuvwyz"}, Port: "1000"}

	got := ParseConfig(yamlpath)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}

}
