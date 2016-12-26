package main

import (
	"os"
	"fmt"
	"log"
	// "flag"
	// "bufio"
	// "io"
	"io/ioutil"
	"encoding/json"
)

func main() {
	operationType := os.Args[1]

	fmt.Println(operationType)

	config := loadConfiguration(" ")

	log.Println(config);

	config.addAlias("new", "path")

	log.Println(config)

	config.removeAlias("new")

	log.Println(config)

	config.saveConfiguration("")

}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Shit our error is bad", err)
		panic(err)
	}
}

type configStruct struct {
	Alias map[string]string `json:"alias"`
}

func loadConfiguration(location string) configStruct {
	configBytes, err := ioutil.ReadFile("./config.json")
	handleErr(err)
	var config configStruct
	err2 := json.Unmarshal(configBytes, &config)
	handleErr(err2)
	log.Println(config)
	return config
}

func (config *configStruct) addAlias(alias string, path string) {
	config.Alias[alias] = path
}

func (config *configStruct) removeAlias(alias string) {
	delete(config.Alias, alias)
}

func (config *configStruct) saveConfiguration(saveLocation string) {
	configBytes, err := json.MarshalIndent(config, "", "    ")
	handleErr(err)
	writeErr := ioutil.WriteFile(saveLocation + "./config.json", configBytes, 0755)
	handleErr(writeErr)
}