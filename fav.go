package main

import (
	"net/url"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		panic("a problem occurred")
	}

	screenName := os.Args[1]

	api := InitApi("config.json")

	v := url.Values{}
	v.Set("screen_name", screenName)
	v.Add("count", "20")
	v.Add("trim_user", "true")
	v.Add("include_rts", "false")
	list, err := api.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}
	var ids []int64
	for _, tw := range list {
		ids = append(ids, tw.Id)
	}

	for _, id := range ids {
		// fav tweets
		_, err := api.Favorite(id)
		if err != nil {
			panic(err)
		}
	}
}
