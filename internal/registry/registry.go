package registry

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	LOGRUS_KEY   = "logrus"
	LOG_FILE_KEY = "log_file"
)

var v = viper.New()

func SetValue(key, value string) {
	v.Set(key, value)
}

func SetLogger(l *logrus.Entry) {
	v.Set(LOGRUS_KEY, l)
}

func GetLogger() *logrus.Entry {
	l, _ := v.Get(LOGRUS_KEY).(*logrus.Entry)
	return l
}

func GetValue(key string) string {
	val, _ := v.Get(key).(string)
	return val
}
