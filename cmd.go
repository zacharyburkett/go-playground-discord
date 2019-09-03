package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	prefix = "```Go"
	suffix = "```"
)

type runResponse struct {
	Errors string                   `json:"Errors"`
	Events []map[string]interface{} `json:"Events"`
}

func exec(code, channelID string) {
	r, err := http.PostForm("https://play.golang.org/compile", url.Values{
		"version": {"2"},
		"body":    {code},
		"withVet": {"true"},
	})
	if err != nil {
		discord.ChannelMessageSend(channelID, err.Error())
		return
	}
	defer r.Body.Close()

	output := &runResponse{}

	data, _ := ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(data, &output); err != nil {
		discord.ChannelMessageSend(channelID, err.Error())
		return
	}

	if output.Errors != "" {
		discord.ChannelMessageSend(channelID, "`"+output.Errors+"`")
		return
	}
	discord.ChannelMessageSend(channelID, "```\n"+output.Events[0]["Message"].(string)+"\n```")
}

func inputHandler(discord *discordgo.Session, mc *discordgo.MessageCreate) {
	msg := mc.Message
	if !strings.HasPrefix(msg.Content, prefix) ||
		!strings.HasSuffix(msg.Content, suffix) {
		return
	}

	code := strings.TrimSuffix(strings.TrimPrefix(msg.Content, prefix), suffix)
	go exec(code, msg.ChannelID)
}
