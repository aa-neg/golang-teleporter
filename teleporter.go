package main

import (
	"os"
	"fmt"
	"flag"
	// "bufio"
	// "io"
	"io/ioutil"
	// "endcoding/json"
)

func handleErr(err error) {
	fmt.Println("here is our error", err)
	if err != nil {
		fmt.Println("Shit our error is bad", err)
		panic(err)
	}
}

func main() {

	operationType := os.Args[1]

	fmt.Println(operationType)

	helpMessage := flag.String("help", "did you require any help mate?", "more help")

	flag.Parse()

	fmt.Println(helpMessage)

	_, err := ioutil.ReadFile("./config.json")
	handleErr(err)



}