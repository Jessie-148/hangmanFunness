package main

import (
 // import (	"fmt"
 // This is a placeholder for something awesome.	"strings"
)

func getCategory() string {
	fmt.Println("Choose a category: animals, countries, programming")
	var category string
	fmt.Scanln(&category)
	return strings.ToLower(category)
}

