package main

import (
	"os"
	"fmt"
	"log"
	"github.com/kardianos/osext"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
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
			Action: func(c *cli.Context) {
				log.Println(addAlias(c, config, executableFolder))
			},
		},
		{
			Name: "remove",
			Aliases: []string{"r", "rm"},
			Usage: "Remove an alias",
			Action: func(c *cli.Context) {
				log.Println(removeAlias(c, config, executableFolder))
			},
		},
		{
			Name: "list",
			Aliases: []string{"l","ls"},
			Usage: "List current aliases",
			Action: func(c *cli.Context) error {
				fmt.Println("")
				listAliases(config.Alias)
				fmt.Println("")
				return nil
			},
		},
		{
			Name: "teleport",
			Aliases: []string{"to", "go"},
			Usage: "Teleport to an alias location, <alias name>",
			Action: func(c *cli.Context) {
				log.Println(teleportTo(c, config))
			},
		},
	}

	app.Run(os.Args)

}

func teleportTo(context *cli.Context, config configStruct) string {
	args := context.Args()

	switch len(args) {
		case 0:
			return "Please specify an alias to teleport to"
        case 1:
			if filePath, exists := config.Alias[args[0]]; exists {
				fmt.Println(filePath)
				os.Exit(2)
			} else {
				return "Alias doesn't exist"
			}
        default:
			return "Invalid numbre of arguments"
	}

	return ""
}

func listAliases(aliasDict map[string]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Alias", "Path"})
	table.SetBorders(tablewriter.Border{Left: false, Top:false, Right: false, Bottom: false})
	table.SetCenterSeparator("  ")
	table.SetColumnSeparator("  ")
	table.SetRowSeparator("-")
	for alias, path := range aliasDict {
		table.Append([]string{alias, filepath.Clean(path)})
	}

	table.Render()
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Shit hapened: ", err)
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

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func addAlias(context *cli.Context, config configStruct, executableFolder string) string{
	args := context.Args()
	switch len(args) {
		case 0:
			return "Please specifty an alias"
        case 1:
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			handleErr(err)
			if path, exists := config.Alias[args[0]]; exists {
				 return ("Alias: " + args[0] + " already exists as path: " + path + " . Please remove before adding.")
			} else {
				config.Alias[args[0]] = dir
				saveConfiguration(config, executableFolder)
			}
        case 2:
            if exists, _ := pathExists(args[1]); exists {
				config.Alias[args[0]] = args[1]
				saveConfiguration(config, executableFolder)
			} else {
				return "Unable to add alias. Invalid path."
			}
        default:
			return "Invalid number of arguments. See --help."
    
	}

	return ""
}

func removeAlias(context *cli.Context, config configStruct, executableFolder string) string {
	args := context.Args()

	switch len(args) {
		case 0:
			return "Please specify an alias to remove"
        case 1:
			if _, exists := config.Alias[args[0]]; exists {
				delete(config.Alias, args[0])
				saveConfiguration(config, executableFolder)
				return "Removed alias: " + args[0]
			} else {
				return "Alias doesn't exist"
			}
        default:
			return "Invalid numbre of arguments"
	}

	return ""
}

func saveConfiguration(config configStruct,location string) {
	configBytes, err := json.MarshalIndent(config, "", "    ")
	handleErr(err)
	writeErr := ioutil.WriteFile(location + "/config.json", configBytes, 0755)
	handleErr(writeErr)
}