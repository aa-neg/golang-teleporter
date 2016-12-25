package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	// "bufio"
	// "io"
	"io/ioutil"
	"encoding/json"
)

type configStruct struct {
	Alias map[string]string `json:"alias"`
}

// type configStruct struct {
// 	alias map[string]interface{}
// }

func handleErr(err error) {
	if err != nil {
		fmt.Println("Shit our error is bad", err)
		panic(err)
	}
}

func (config *configStruct) addAlias(alias string, path string) {
	config.Alias[alias] = path
}

func main() {

	operationType := os.Args[1]

	fmt.Println(operationType)

	helpMessage := flag.String("help", "did you require any help mate?", "more help")

	flag.Parse()

	fmt.Println(helpMessage)

	var config configStruct

	
	configBytes, err := ioutil.ReadFile("./config.json")
	handleErr(err)

	err2 := json.Unmarshal(configBytes, &config)
	handleErr(err2)

	fmt.Println("Here is our config")
	fmt.Println(config)

	alias := config.Alias["working"]
	fmt.Println("here is our alias")
	fmt.Println(alias)
	// fmt.Println(config["alias"])
	// fmt.Println(config["alias"]["working"])



}