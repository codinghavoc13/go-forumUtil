package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BuildInsertStmt(output *string) {
	*output = *output + LINE_BREAK + INSERT_ROOMS + NEW_LINE
	*output = *output + LINE_BREAK + INSERT_USERS + NEW_LINE
}

func BuildInsertPostStmts(output *string, responses *string, title string) {
	//INSERT INTO notification_demo.posts(
	// date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id)
	// VALUES (?, ?, ?, ?, ?, ?, ?);

	/*
		Need to build out all of the fields for the post except for the date_last_updated
		Then get a random number between 2 and 8, call buildInsertResponseStmt
		Then call the buildInsertResponseStmt which will return a date
	*/

	var date_last_updated time.Time
	number_responses := 0
	postDate := time.Now()
	post_text := PARAGRAPHS[rand.Intn(len(PARAGRAPHS))]
	post_title := title
	user_id := rand.Intn(20)
	room_id := rand.Intn(6)

	responseCount := rand.Intn(6) + 2
	for i := 0; i < responseCount; i++ {
		date_last_updated = BuildInsertResponseStmt(responses)
		number_responses++
	}
	date_last_updated_formatted := "'" + date_last_updated.String()[:27] + "'"
	postDate_formatted := "'" + postDate.String()[:27] + "'"
	var title_formatted = "'" + post_title + "'"
	var text_formatted = "'" + post_text + "'"

	*output = *output + fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v)",
		date_last_updated_formatted, number_responses, postDate_formatted, text_formatted, title_formatted, user_id, room_id)
}

func BuildInsertResponseStmt(responses *string) time.Time {
	// INSERT INTO notification_demo.responses(
	// 	post_id, response_date, response_text, user_id)
	// 	VALUES (?, ?, ?, ?);
	return time.Now()
}
