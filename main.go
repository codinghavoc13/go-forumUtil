package main

import "os"

var TITLES = []string{}
var PARAGRAPHS = []string{}

func main() {
	var test = true
	var output = ""
	var responses = ""
	if test {
		TITLES = getDataFromFile(RAW_TITLES)
		PARAGRAPHS = getDataFromFile(RAW_PARAGRAPHS)
		BuildCreateStmt(&output)
		BuildInsertStmt(&output)
		BuildInsertPostStmts(&output, &responses, 20)
		output = output + responses
		os.WriteFile("output/forum_build.sql", []byte(output), 0644)
	} else {
		dateTesting()
		checkingRandomNumber()
		getTwoDifferentUserIDs()
		buildingString(&output)
	}
	// fmt.Println("Titles:", len(getDataFromFile(RAW_TITLES)))
	// fmt.Println("Paragraphs:", len(getDataFromFile(RAW_PARAGRAPHS)))
}
