package main

import (
	//"encoding/json"
	"fmt"
	//"github.com/ChimeraCoder/anaconda"
	//"io/ioutil"
	"log"
	"os"
)

//// 小文字から始まる名前を持つものはパッケージ外からアクセスできないので、大文字にしている
//type Config struct {
//	Consumer_key     string `json:"consumer_key"`
//	Consumer_secret  string `json:"consumer_secret"`
//	Access_token     string `json:"access_token"`
//	Access_token_key string `json:"access_token_key"`
//}

func main() {
//	file, err := ioutil.ReadFile("config.json")
//	if err != nil {
//		panic(err)
//	}
//
//	var config Config
//	json.Unmarshal(file, &config)
//
//	anaconda.SetConsumerKey(config.Consumer_key)
//	anaconda.SetConsumerSecret(config.Consumer_secret)
//	api := anaconda.NewTwitterApi(config.Access_token, config.Access_token_key)

	api := InitApi("config.json")

	text := os.Args[1]
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tweet.Text)
}
