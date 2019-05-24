package logger

import (
	"github.com/dongshimou/logrus"
	prefixed "github.com/dongshimou/logrus-prefixed-formatter"
)

var (
	log    = logrus.New()
	prefix = "untitled"
)

func init() {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "15:04:05"

	log.Formatter = formatter
	log.Level = logrus.DebugLevel
	log.SetReportCaller(true)
	log.CallerSkip = 1
}

func New(server string) {
	prefix = server
	//entry=log.WithField("prefix",prefix)
}
func appendSpace(args ...interface{})[]interface{}{
	res:=[]interface{}{}
	for _,a:=range args{
		res=append(res,a)
		res=append(res," ")
	}
	return res[:len(res)-1]
}
func Debug(args ...interface{}) {
	log.WithField("prefix", prefix).Debug(appendSpace(args)...)
}
func Warn(args ...interface{}) {
	log.WithField("prefix", prefix).Warn(appendSpace(args)...)
}
func Error(args ...interface{}) {
	log.WithField("prefix", prefix).Error(appendSpace(args)...)
}
func Fatal(args ...interface{}) {
	log.WithField("prefix", prefix).Fatal(appendSpace(args)...)
}
func Warning(args ...interface{}) {
	log.WithField("prefix", prefix).Warning(appendSpace(args)...)
}
func Info(args ...interface{}) {
	log.WithField("prefix", prefix).Info(appendSpace(args)...)
}
func Get() *logrus.Entry {
	return log.WithField("prefix", prefix)
}
