package main

import "fmt"

func switchList(listName string) {
	globalConfig.CurrentList = listName
	fmt.Print(globalConfig.CurrentList)
	saveConfig(globalConfig)
}
