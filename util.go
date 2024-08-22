package main

import (
	"os"
	"strings"
)

/*
Reads a long list of short lipsum generated content and breaks it down into
a list of strings
*/
func getDataFromFile(filename string) []string {
	data, _ := os.ReadFile(filename)
	var splitData = strings.Split(string(data), "\n")
	return splitData
}
