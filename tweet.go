package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	api := InitApi("config.json")

	text := os.Args[1]
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tweet.Text)
}
