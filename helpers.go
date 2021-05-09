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

// createItems makes the item struct to be inserted into the list
// Takes the descriptors([]string) and priorities([]int) and returns []item
func createItems(itemDescription []string, priorities []int) []item {
	var newItems []item
	for i, itemDesc := range itemDescription {

		var priority priority = priority(priorities[i])
		newItem := item{
			Description: itemDesc,
			Completed:   false,
			Priority:    priority,
		}
		newItems = append(newItems, newItem)
	}
	return newItems
}

// openTodoLists open the json todo list file and retunrns a struct containg each list indivdualy
func openTodoLists() []list {
	var jsonLists []list
	jsonContent, _ := ioutil.ReadFile(globalConfig.SaveLocation)
	json.Unmarshal([]byte(jsonContent), &jsonLists)
	return jsonLists
}

// saveTodoLists saves the todo lists to a JSON format that can be recovered post termination for future use
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

// saveConfig savest the global configuration as JSON so that can be recovered post termination for future use
func saveConfig(config config) {
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
