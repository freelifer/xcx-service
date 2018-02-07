package config

import (
	"github.com/Unknwon/goconfig"
	"strings"
)

var c *goconfig.ConfigFile

func init() {
	c, _ = goconfig.LoadConfigFile("conf/app.conf")
}

func MustValue(section, key string, defaultVal string) string {
	if c != nil {
		return c.MustValue(section, key, defaultVal)
	}
	return defaultVal
}

func MustBool(section, key string, defaultVal bool) bool {
	if c != nil {
		return c.MustBool(section, key, defaultVal)
	}
	return defaultVal
}

func MustInt(section, key string, defaultVal int) int {
	if c != nil {
		return c.MustInt(section, key, defaultVal)
	}
	return defaultVal
}

func String(key string) string {
	return DefaultString(key, "")
}

func DefaultString(key, defaultVal string) string {
	pos := strings.Index(key, "::")
	if pos > -1 {
		section := key[:pos]
		next := key[pos+2:]
		return MustValue(section, next, defaultVal)
	}

	return defaultVal
}

func Bool(key string) bool {
	return DefaultBool(key, false)
}

func DefaultBool(key string, defaultVal bool) bool {
	pos := strings.Index(key, "::")
	if pos > -1 {
		section := key[:pos]
		next := key[pos+2:]
		return MustBool(section, next, defaultVal)
	}

	return defaultVal
}

func Int(key string) int {
	return DefaultInt(key, 0)
}

func DefaultInt(key string, defaultVal int) int {
	pos := strings.Index(key, "::")
	if pos > -1 {
		section := key[:pos]
		next := key[pos+2:]
		return MustInt(section, next, defaultVal)
	}

	return defaultVal
}
