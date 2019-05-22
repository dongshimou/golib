package logger

import (
	prefixed "github.com/dongshimou/logrus-prefixed-formatter"
	"github.com/sirupsen/logrus"
)

var (

log logrus.New()

)

func init(){
}

func New(server string){

	formatter:=new(prefixed.TextFormatter)
	formatter.FullTimestamp=true
	formatter.TimestampFormat="15:04:05"


	log.Formatter = formatter
	log.Level = logrus.DebugLevel
	log.SetReportCaller(true)

	log=log.WithFields(logrus.Fields{"prefix":server})
}