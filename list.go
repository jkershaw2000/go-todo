package main

import "fmt"

func listTasks(listName string) {
	jsonLists := openTodoLists()
	listExsits := false
	for _, list := range jsonLists {
		if list.Name == listName {
			for _, item := range list.Items {
				fmt.Println(item.Description)
			}
			listExsits = true
			break
		}
	}
	if !listExsits {
		fmt.Println("List not found.")
	}
}
