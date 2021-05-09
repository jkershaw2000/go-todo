package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func listTasks(listName string) {
	jsonLists := openTodoLists()
	listExsits := false
	table := tablewriter.NewWriter(os.Stdout)
	// Configure table format
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	// ID is just position in array. Will change each time.
	table.SetHeader([]string{"ID", "Description", "Priority", "Complete"})
	for _, list := range jsonLists {
		if list.Name == listName {
			for i, item := range list.Items {
				table.Append([]string{strconv.Itoa(i), item.Description, item.Priority.string(), strconv.FormatBool(item.Completed)})
			}
			listExsits = true
			table.Render()
			break
		}
	}
	if !listExsits {
		fmt.Println("List not found.")
	}
}
