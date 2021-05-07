package main

import "fmt"

func add(listName string, itemDescriptions []string, forceCreate bool) {
	jsonLists := openTodoLists()
	newItems := createItems(itemDescriptions)

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
