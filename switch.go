package main

import "fmt"

func switchList(listName string) {
	jsonLists := openTodoLists()
	found := false
	for _, list := range jsonLists {
		if list.Name == listName {
			found = true
			break
		}
	}
	if found {
		globalConfig.CurrentList = listName
		saveConfig(globalConfig)
		fmt.Printf("Current list switched to %s\n", listName)
	} else {
		fmt.Printf("%s doesn't exsits. Not change made", listName)
	}
}
