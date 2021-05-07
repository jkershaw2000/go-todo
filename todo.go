package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/thatisuday/commando"
)

var savePath string
var globalConfig config

// Config struct for saving global settings to be saved between uses.
type config struct {
	CurrentList, SaveLocation string
}

// List struct containg the items constutuating the todo list.
type list struct {
	Name  string
	Items []item
}

// Item struct describing each task within a todo list.
type item struct {
	Description string
	Completed   bool
	Priority    int
}

func main() {
	commando.
		SetExecutableName("todo").
		SetVersion("0.1").
		SetDescription("A simple CLI todo program supporting multiple lists.\nLets you get orgnaized by being able to view your tasks diretcly from your terminal!")

	// Register subcommands
	commando.
		Register("add").
		SetDescription("Add an item to the specified list.\nIf none specified adds to current.").
		SetShortDescription("Add an item to list").
		AddArgument("item...", "Item(s) to be added to list", "").
		AddFlag("list,l", "Specify what list the item(s) should be added to", commando.String, globalConfig.CurrentList).
		AddFlag("force,f", "Create a list of the specified name if it doesn't exsits.", commando.Bool, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			list, _ := flags["list"].GetString()
			create, _ := flags["force"].GetBool()
			items := strings.Split(args["item"].Value, ",")

			add(list, items, create)
		})

	commando.
		Register("create").
		SetDescription("Creates a new list with the given name").
		SetShortDescription("Creates a new list").
		AddArgument("name...", "Name of list to be created", "").
		AddFlag("force,f", "Force create a new list even if one with same name exsits. Will delete previous list.", commando.Bool, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			lists := strings.Split(args["name"].Value, ",")
			force, _ := flags["force"].GetBool()
			create(lists, force, nil)
		})

	commando.
		Register("clean").
		SetDescription("Removes completed items from todo list. Optionaly move to an archived list").
		SetShortDescription("Removes completed items from todo list")

	commando.
		Register("complete").
		SetDescription("Marks selected task as complete").
		SetShortDescription("Marks task as complete")

	commando.
		Register("current").
		SetDescription("Gets name of current todo list").
		SetShortDescription("Gets name of current todo list").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			current()
		})

	commando.
		Register("edit").
		SetDescription("Change task description in todo list").
		SetShortDescription("Change task description in todo list.")

	commando.
		Register("list").
		SetDescription("Show all items in the selected todo list").
		SetShortDescription("Show all items in the selected todo list").
		AddFlag("list,l", "Specify what list to show. (default: current list)", commando.String, globalConfig.CurrentList).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			listName, err := flags["list"].GetString()
			if err != nil {
				fmt.Println(err)
			}
			listTasks(listName)
		})

	commando.
		Register("remove").
		SetDescription("Removes item from todo list").
		SetShortDescription("Removes item from todo list").
		AddArgument("item...", "What item should be removed from the todo list.", "").
		AddFlag("archive,a", "Should the removed item be archived. (default: true", commando.Bool, true).
		AddFlag("list,l", "Specify what list to remove an item from. (default: current list)", commando.String, globalConfig.CurrentList)

	commando.
		Register("switch").
		SetDescription("Set specified list as the current list").
		SetShortDescription("Set specified list as the curent list").
		AddArgument("list", "What list should be the current list", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			switchList(args["list"].Value)
		})

	setup()
	commando.Parse(nil)
}

func setup() {
	savePath, _ = homedir.Dir()
	savePath += "/.todo/"
	// Makes the ~/.todo/ folder if it doesnt already exsit
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		err = os.Mkdir(savePath, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
	// If the settings file doesn't exsist, create one with default settings.
	if _, err := os.Stat(savePath + "settings.json"); os.IsNotExist(err) {
		os.Create(savePath + "settings.json")
		conf := config{"main", savePath + "lists.json"}
		bytes, err := json.Marshal(conf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(bytes)
		err = ioutil.WriteFile(savePath+"settings.json", bytes, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Get settings from file and convert to config struct
	jsonContent, _ := ioutil.ReadFile(savePath + "settings.json")
	json.Unmarshal([]byte(jsonContent), &globalConfig)

	// If list file doesn't exsits create it and make default list called "main"
	if _, err := os.Stat(globalConfig.SaveLocation); os.IsNotExist(err) {
		os.Create(globalConfig.SaveLocation)
		defaultLists := list{"main", []item{}}
		bytes, err := json.Marshal(defaultLists)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(bytes)
		err = ioutil.WriteFile(globalConfig.SaveLocation, bytes, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

}
