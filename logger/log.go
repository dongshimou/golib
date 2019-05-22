package logger

import (
	prefixed "github.com/dongshimou/logrus-prefixed-formatter"
	"github.com/dongshimou/logrus"
)

var (
	log =logrus.New()
	prefix="untitled"
)

func init(){
	formatter:=new(prefixed.TextFormatter)
	formatter.FullTimestamp=true
	formatter.TimestampFormat="15:04:05"

	log.Formatter = formatter
	log.Level = logrus.DebugLevel
	log.SetReportCaller(true)
}

func New(server string){
	prefix=server
	log.CallerSkip=1
	//entry=log.WithField("prefix",prefix)
}
func Debug(args ...interface{}){
	log.WithField("prefix",prefix).Debug(args...)
}
func  Warn(args ...interface{}) {
	log.WithField("prefix",prefix).Warn(args...)
}
func Error(args ...interface{}){
	log.WithField("prefix",prefix).Error(args...)
}
func Fatal(args ...interface{}){
	log.WithField("prefix",prefix).Fatal(args...)
}
func Warning(args...interface{}){
	log.WithField("prefix",prefix).Warning(args...)
}
func Info(args...interface{}){
	log.WithField("prefix",prefix).Info(args...)
}
func Get()*logrus.Entry{
	return log.WithField("prefix",prefix)
}