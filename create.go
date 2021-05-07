package main

import (
	"fmt"
)

// Create a new todo list with the provided names. If force is true, overwrite previous list(s) with same name
func create(lists []string, force bool, items []item) {
	jsonLists := openTodoLists()

	var listNames []string
	for _, exsitingList := range jsonLists {
		listNames = append(listNames, exsitingList.Name)
	}
	var listItems []item
	if len(items) >= 1 {
		listItems = items
	} else {
		listItems = nil
	}
	for _, l := range lists {
		_, lExsits := Find(listNames, l)
		if l != "" && !lExsits {
			jsonLists = append(jsonLists, list{l, listItems})
			fmt.Printf("List called %s created.\n", l)
		} else if force && lExsits {
			for _, exsisitingList := range jsonLists {
				if exsisitingList.Name == l {
					exsisitingList.Items = listItems
				}
			}
			fmt.Printf("List called %s created. Previous version has been lost.\n", l)
		} else {
			fmt.Printf("List called %s already exsits. No action was taken.\n", l)
		}
	}

	saveTodoLists(jsonLists)
}
