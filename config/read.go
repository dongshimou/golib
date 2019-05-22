package config

import (
	"encoding/json"
	"fmt"
	"github.com/koding/multiconfig"
	"os"
)

const(
	SERVER_ENV="SERVER_ENV"
)

func serverEnv()string{
	return os.Getenv(SERVER_ENV)
}

func Read(obj interface{})error{
	return readConfig(obj)
}

func readConfig(obj interface{}) error {

	e:=serverEnv()
	var configPath string
	if e!=""{
		configPath=e+".json"
	}else{
		configPath="local.json"
	}

	if err:=readConfigFile(obj,configPath);err!=nil{
		return err
	}
	ind, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		return err
	}
	fmt.Println("-----------------------------")
	fmt.Printf("Load Config from %s :\n", configPath)
	fmt.Println(string(ind))
	fmt.Println("-----------------------------")
	return nil
}

func readConfigFile(obj interface{}, filePath string) error {
	m := multiconfig.New()
	m.Loader = &multiconfig.JSONLoader{Path: filePath}
	err := m.Load(obj)
	if err != nil {
		return err
	}
	return m.Validate(obj)
}