package main

import (
	"os"
	"fmt"
	"log"
	"github.com/kardianos/osext"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	// "flag"
	// "flag"
	// "bufio"
	// "io"
	"io/ioutil"
	"encoding/json"
)

func main() {
	executableFolder, err := osext.ExecutableFolder()
	handleErr(err)

	config := loadConfiguration(executableFolder)

	app := cli.NewApp()

	app.Name = "Teleporter"

	app.Usage = "Alias paths and teleport to them!"

	app.Commands = []cli.Command{
		{
			Name: "add",
			Aliases: []string{"a"},
			Usage: "Add an alias, <alis name> <optional path, default current>",
			Action: func(c *cli.Context) error {
				log.Println("We selected add:", c.Args())
				return nil
			},
		},
		{
			Name: "remove",
			Aliases: []string{"r", "rm"},
			Usage: "Remove an alias",
			Action: func(c *cli.Context) error {
				log.Println("We have removed an alias", c.Args())
				return nil
			},
		},
		{
			Name: "list",
			Aliases: []string{"l","ls"},
			Usage: "List current aliases",
			Action: func(c *cli.Context) error {
				log.Println("We will list our current aliases")
				for alias, path := range config.Alias {
					fmt.Println("Alias: ", alias, "Path: ", path)
				}
				return nil
			},
		},
		{
			Name: "teleport",
			Aliases: []string{"to", "go"},
			Usage: "Teleport to an alias location",
			Action: func(c *cli.Context) error {
				log.Println("About to teleport somewhere?")
				return nil
			},
		},
	}

	app.Run(os.Args)

	// if len(os.Args) < 2 {
	// 	log.Println("Invalid usage see --help for details")
	// 	os.Exit(1)
	// }

	// addAliasCommand := flag.NewFlagSet("add", flag.ExitOnError)
	// removeAliasCommand := flag.NewFlagSet("remove", flag.ExitOnError)
	// listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	// teleportCommand := flag.NewFlagSet("to", flag.ExitOnError)

	// availableCommands := map[string]interface{}{
	// 	"add": addAliasCommand, 
	// 	"remove" :removeAliasCommand,
	// 	"list": listCommand,
	// 	"to": teleportCommand,
	// }

	// if  flagSet, exists := availableCommands[os.Args[1]]; exists {
	// 	log.Println("here are our arguments")
	// 	log.Println(flagSet)
	// } else {
	// 	log.Println("Invalid arguments.")
	// 	os.Exit(1)
	// }
	// operationType := os.Args[1]
	// var executableFolder string
	
	// log.Println(config);

	// config.addAlias("new", "path")

	// log.Println(config)

	// config.removeAlias("new")

	// log.Println(config)

	// config.saveConfiguration(executableFolder)

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
	configBytes, err := ioutil.ReadFile(location + "/config.json")
	handleErr(err)
	var config configStruct
	err2 := json.Unmarshal(configBytes, &config)
	handleErr(err2)
	return config
}

func (config *configStruct) addAlias(alias string, path string) {
	config.Alias[alias] = path
}

func (config *configStruct) removeAlias(alias string) {
	delete(config.Alias, alias)
}

func (config *configStruct) saveConfiguration(location string) {
	configBytes, err := json.MarshalIndent(config, "", "    ")
	handleErr(err)
	writeErr := ioutil.WriteFile(location + "/config.json", configBytes, 0755)
	handleErr(writeErr)
}