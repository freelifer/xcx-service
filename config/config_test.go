package config

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_filepath(t *testing.T) {
	var err error
	var AppPath string
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		t.Error(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	var appConfigPath = filepath.Join(workPath, "conf", "app.conf")
	t.Log(appConfigPath)
	appConfigPath = filepath.Join(AppPath, "conf", "app.conf")
	t.Log(appConfigPath)
}

func Test_string(t *testing.T) {
	r := String("wx::")

	t.Log(r)
}
func Test_mustValue(t *testing.T) {
	r := MustValue("app", "db", "")

	t.Log(r)
}
