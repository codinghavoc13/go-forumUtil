package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dateTesting() {
	var START_YEAR = 2024
	var year, month, day, hour, minute, second int
	year = START_YEAR
	month = rand.Intn(12)
	// day = rand.Intn(31)
	day = 28
	hour = rand.Intn(23)
	minute = rand.Intn(60)
	second = rand.Intn(60)
	for i := 0; i < 20; i++ {
		// fmt.Println("-----------------------------------------------------------")
		// month = rand.Intn(8)
		// day = rand.Intn(31)
		// hour = rand.Intn(23)
		// minute = rand.Intn(60)
		// second = rand.Intn(60)
		hour = hour + 10
		date := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
		fmt.Println(date)
		year = date.Year()
		month = int(date.Month())
		hour = date.Hour()
		day = date.Day()
	}
	date := time.Date(1984, 12, 13, 19, 15, 00, 0, time.Local)
	fmt.Println(date)
}

func buildingString(output *string) {
	*output = *output + CREATE_USER_TABLE + NEW_LINE
	*output = *output + CREATE_ROOMS_TABLE + NEW_LINE
	*output = *output + CREATE_POSTS_TABLE + NEW_LINE
	*output = *output + CREATE_RESPONSE_TABLE

	var scriptsArr []string = []string{}
	scriptsArr = append(scriptsArr, CREATE_USER_TABLE)
	scriptsArr = append(scriptsArr, CREATE_ROOMS_TABLE)
	scriptsArr = append(scriptsArr, CREATE_POSTS_TABLE)
	scriptsArr = append(scriptsArr, CREATE_RESPONSE_TABLE)

	for _, item := range scriptsArr {
		*output = *output + item + NEW_LINE
		*output = *output + NEW_LINE + LINE_BREAK + NEW_LINE
	}

	fmt.Println("RESPONSES length:", len(RESPONSES))
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(len(RESPONSES)))
	}
}

// Simple check to see how distributed a random generation of numbers is
func checkingRandomNumber() {
	results := map[int]int{}
	for i := 0; i < 100; i++ {
		temp := rand.Intn(20)
		_, ok := results[temp]
		if ok {
			results[temp] = results[temp] + 1
		} else {
			results[temp] = 1
		}
	}
	fmt.Println(results)
}
