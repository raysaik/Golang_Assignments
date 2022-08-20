package main

import (
	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
)

func sendMessageToTeams() {
	_ = sendTheMessage()
}

func sendTheMessage() error {
	// init the client
	mstClient := goteamsnotify.NewClient()

	// setup webhook url, fill up this value
	webhookUrl := ""

	// setup message card
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = "Hello world"
	msgCard.Text = "This is a test message " +
		"<br> * this list itself  <br> * **bold** <br> * *italic* <br> * ***bolditalic***"
	msgCard.ThemeColor = "#DF813D"

	// send
	return mstClient.Send(webhookUrl, msgCard)
}
