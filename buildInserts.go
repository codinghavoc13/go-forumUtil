package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func BuildInsertStmt(output *string) {
	*output = *output + LINE_BREAK + INSERT_ROOMS + NEW_LINE
	*output = *output + LINE_BREAK + INSERT_USERS + NEW_LINE
}

func BuildInsertPostStmts(output *string, responses *string, count int) {
	//INSERT INTO notification_demo.posts(
	// date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id)
	// VALUES (?, ?, ?, ?, ?, ?, ?);

	*output = *output + LINE_BREAK + INSERT_POSTS + NEW_LINE
	*responses = *responses + LINE_BREAK + INSERT_RESPONSES + NEW_LINE

	var date_last_updated time.Time
	postDate := time.Now()

	// overallCountMax := count * responseCount
	// var overallCounter = 0
	for i := 1; i < count; i++ {
		number_responses := 0
		user_id := rand.Intn(19) + 1
		room_id := rand.Intn(7) + 1
		post_text := PARAGRAPHS[rand.Intn(len(PARAGRAPHS))]
		post_title := TITLES[i]

		responseCount := rand.Intn(6) + 2
		for j := 1; j < responseCount; j++ {
			date_last_updated = BuildInsertResponseStmt(i, user_id, responses, postDate)
			*responses = *responses + ",\n"
			number_responses++
		}
		date_last_updated_formatted := "'" + date_last_updated.String()[:27] + "'"
		postDate_formatted := "'" + postDate.String()[:27] + "'"
		var title_formatted = "'" + post_title + "'"
		var text_formatted = "'" + post_text + "'"

		*output = *output + fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v)",
			date_last_updated_formatted, number_responses, postDate_formatted, text_formatted, title_formatted, user_id, room_id)

		if i < count-1 {
			*output = *output + ",\n"
		} else {
			*output = *output + "\n"
		}
		// fmt.Println(postDate.Sub(date_last_updated))
		postDate = postDate.AddDate(0, 0, 1)
	}
	*output = strings.TrimRight(*output, ",\n")
	*output = *output + ";"
	*responses = strings.TrimRight(*responses, ",\n")
	*responses = *responses + ";"
}

func BuildInsertResponseStmt(post_id int, orig_user_id int, responses *string, postdate time.Time) time.Time {
	// INSERT INTO notification_demo.responses(
	// 	post_id, response_date, response_text, user_id)
	// 	VALUES (?, ?, ?, ?);

	responseDate := postdate.
		Add(time.Duration((rand.Intn(4) + 1) * int(time.Hour))).
		Add(time.Duration(rand.Intn(60) * int(time.Minute)))
	responseDate_formatted := "'" + responseDate.String()[:27] + "'"
	response_text := RESPONSES[rand.Intn(len(RESPONSES))]
	var response_text_formatted = "'" + response_text + "'"
	user_id := rand.Intn(19) + 1
	for user_id == orig_user_id {
		user_id = rand.Intn(19) + 1
	}

	*responses = *responses + fmt.Sprintf("(%v, %v, %v, %v)",
		post_id, responseDate_formatted, response_text_formatted, user_id)

	// returnDate = postdate.Add(time.Duration(rand.Intn(60) * int(time.Minute)))
	return responseDate
}
