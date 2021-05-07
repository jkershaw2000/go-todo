package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func createItems(itemDescription []string) []item {
	var newItems []item
	for _, itemDesc := range itemDescription {

		newItem := item{
			Description: itemDesc,
			Completed:   false,
			Priority:    0,
		}
		newItems = append(newItems, newItem)
	}
	return newItems
}

func openTodoLists() []list {
	var jsonLists []list
	jsonContent, _ := ioutil.ReadFile(globalConfig.SaveLocation)
	json.Unmarshal([]byte(jsonContent), &jsonLists)
	return jsonLists
}
func saveTodoLists(jsonLists []list) {
	// Marshals struct into json to save
	bytes, err := json.Marshal(jsonLists)
	if err != nil {
		fmt.Println(err)
	}
	// Writes JSON data to the specifed save location
	err = ioutil.WriteFile(globalConfig.SaveLocation, bytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func saveConfig(config config) {
	fmt.Print(config.CurrentList)
	// Marshals struct into json to save
	bytes, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}
	// Writes JSON data to the specifed save location
	err = ioutil.WriteFile(savePath+"settings.json", bytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
