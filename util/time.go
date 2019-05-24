package util

import (
	"strings"
	"time"
)

func ParseTime(str string)(time.Time,error){

	if strings.Contains(str,"-"){
		if t1,err:=time.Parse(DatetimeFormat,str);err==nil{
			return t1,nil
		}else{
			if t2,err:=time.Parse(DateFormat,str);err==nil{
				return t2,nil
			}else{
				return time.Time{},err
			}
		}
	}else{
		return time.Parse(TimeFormat,str)
	}
}