package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/rylio/ytdl"
)

func DownloadVideo(videoID string) {
	vid, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=" + videoID)
	if err != nil {
		fmt.Println("Failed to get video info")
		return
	}
	file, _ := os.Create(vid.Title + ".mp4")
	defer file.Close()
	vid.Download(vid.Formats[0], file)
}

func UpdateChannelsDatabase(channelURL string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	jsonFile, err := os.Open("channels.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var db []Channel

	json.Unmarshal(byteValue, &db)

	var exists bool

	for _, v := range db {
		if v.ChannelURL == channelURL {
			fmt.Println("already exists:", channelURL)
			exists = true
			break
		} else {
			fmt.Println("doesnt exist:", channelURL)
			exists = false
		}
	}

	if exists == true {
		fmt.Println("channel already added")
	} else {
		db = append(db, Channel{ChannelURL: channelURL})

		result, err := json.Marshal(db)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(result))

		json.Unmarshal(result, &db)

		file, _ := json.MarshalIndent(db, "", " ")

		_ = ioutil.WriteFile("channels.json", file, 0644)
	}
}

func GetChannels() []Channel {
	jsonFile, err := os.Open("channels.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var db []Channel

	json.Unmarshal(byteValue, &db)

	return db
}
