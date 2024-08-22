package main

import (
	"os"
)

var TITLES = []string{}
var PARAGRAPHS = []string{}

func main() {
	var output = ""
	var responses = ""
	TITLES = getDataFromFile(RAW_TITLES)
	PARAGRAPHS = getDataFromFile(RAW_PARAGRAPHS)
	BuildCreateStmt(&output)
	BuildInsertStmt(&output)
	output = output + LINE_BREAK + INSERT_POSTS
	for i := 0; i < 10; i++ {
		BuildInsertPostStmts(&output, &responses, TITLES[i])
		if i < 9 {
			output = output + ",\n"
		}
	}
	os.WriteFile("output/forum_build.sql", []byte(output), 0644)
	// checkingRandomNumber()
	// fmt.Println("Titles:", len(getDataFromFile(RAW_TITLES)))
	// fmt.Println("Paragraphs:", len(getDataFromFile(RAW_PARAGRAPHS)))
}
