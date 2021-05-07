package main

import "fmt"

func current() {
	fmt.Printf("Current list: %s\n", globalConfig.CurrentList)
}
