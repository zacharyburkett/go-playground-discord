package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type configuration struct {
	BotToken string `json:"bot_token"`
}

var conf *configuration

func init() {
	conf = &configuration{}
}

func loadConf() error {
	f, err := os.Open("conf.json")
	if err != nil {
		return err
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(data, &conf)

	return err
}
