package main

func BuildCreateStmt(output *string) {
	*output = *output + LINE_BREAK + CREATE_USER_TABLE + NEW_LINE
	*output = *output + LINE_BREAK + CREATE_ROOMS_TABLE + NEW_LINE
	*output = *output + LINE_BREAK + CREATE_POSTS_TABLE + NEW_LINE
	*output = *output + LINE_BREAK + CREATE_RESPONSE_TABLE + NEW_LINE
}
