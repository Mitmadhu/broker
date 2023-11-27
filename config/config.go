package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Endpoints map[string]string `json:"endpoints"`
	Port      int               `json:"port"`
}

var Configs Config

const configPath = "/config.json"

func Init() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	// load configs else panic
	file, err := os.Open(wd + configPath)
	if err != nil {
		panic(fmt.Sprintf("error while opening file : %v", err.Error()))
	}

	defer file.Close()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("error while reading file :  %v", err.Error()))
	}

	err = json.Unmarshal(byteData, &Configs)

	if err != nil {
		panic(fmt.Sprintf("error while unmarshalling byte data :  %v", err.Error()))
	}

}
