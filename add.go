package main

import "fmt"

func add(listName string, priority int, itemDescription string, forceCreate bool) {
	jsonLists := openTodoLists()
	// Send as slices for convienience
	newItems := createItems([]string{itemDescription}, []int{priority})
	if listName == "zzDEFAULTzz" {
		listName = globalConfig.CurrentList
	}
	var listExsists bool
	for i, list := range jsonLists {
		if listName == list.Name {
			listExsists = true
			for _, item := range newItems {
				jsonLists[i].Items = append(jsonLists[i].Items, item)
			}
			fmt.Printf("Items(s) added to list called %s.\n", listName)
			saveTodoLists(jsonLists)
			break
		}
	}
	if !listExsists && forceCreate {
		create([]string{listName}, false, newItems)
		fmt.Printf("Item(s) added to new list called %s.\n", listName)
	} else if !listExsists && !forceCreate {
		fmt.Printf("No list called %s was found. Nothing occured.\n", listName)
	}

}
