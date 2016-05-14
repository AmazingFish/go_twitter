package main

import (
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
	"io/ioutil"
)

type Config struct {
	Consumer_key     string `json:"consumer_key"`
	Consumer_secret  string `json:"consumer_secret"`
	Access_token     string `json:"access_token"`
	Access_token_key string `json:"access_token_key"`
}

func InitApi(configFileName string) *anaconda.TwitterApi {
	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}

	var config Config
	json.Unmarshal(file, &config)

	anaconda.SetConsumerKey(config.Consumer_key)
	anaconda.SetConsumerSecret(config.Consumer_secret)
	api := anaconda.NewTwitterApi(config.Access_token, config.Access_token_key)
	return api
}
