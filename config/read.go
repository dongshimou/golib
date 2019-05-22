package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/koding/multiconfig"
	"os"
)

const (
	SERVER_ENV = "SERVER_ENV"
)

func serverEnv(env string) string {
	return os.Getenv(env)
}

func Read(obj interface{}, args ...interface{}) error {
	env := SERVER_ENV
	if len(args) >= 1 {
		return errors.New("too much args")
	} else if len(args) == 1 {
		arg, ok := args[0].(string)
		if !ok {
			return errors.New("not a string")
		}
		env = arg
	}
	return readConfig(obj, env)
}

func readConfig(obj interface{}, env string) error {

	e := serverEnv(env)
	var configPath string
	if e != "" {
		configPath = e + ".json"
	} else {
		configPath = "local.json"
	}

	if err := readConfigFile(obj, configPath); err != nil {
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
