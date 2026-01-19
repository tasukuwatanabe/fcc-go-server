package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	if err := godotenv.Load(fmt.Sprintf("env/%s.env", os.Getenv("GO_ENV"))); err != nil {
		fmt.Println("Error loading env file")
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelID := os.Getenv("CHANNEL_ID")

	filename := "sample.pdf"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	params := slack.UploadFileV2Parameters{
		Channel:  channelID,
		File:     filename,
		Filename: filename,
		FileSize: int(fileInfo.Size()),
	}

	uploadedFile, err := api.UploadFileV2(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s posted to slack channel\n", uploadedFile.Title)
}
