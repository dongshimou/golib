package main

import (
	"github.com/dongshimou/golib/config"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Point float64 `json:"point"`
}

func main(){

	p:=Person{}

	if err:=config.Read(&p);err!=nil{
		panic(err)
	}

}