package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, GetChannels())
}

func HandleCheckChannel(w http.ResponseWriter, r *http.Request) {
	channelURL := r.FormValue("channelURL")
	fmt.Println(channelURL)
	// channelName := strings.Split(channelURL, "/")[4]
	// channelType := strings.Split(channelURL, "/")[3]

	// if channelType == "user" {
	// 	uploadsId := GetUserUploadsID(channelName)
	// 	videoId, videoTitle := GetUserVideoData(uploadsId)
	// 	DownloadVideoAndAudio(videoId, videoTitle)
	// 	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)

	// } else if channelType == "channel" {
	// 	videoId, videoTitle := GetChannelVideoData(channelName)
	// 	DownloadVideoAndAudio(videoId, videoTitle)
	// 	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)

	// }
}

func HandleAddChannel(w http.ResponseWriter, r *http.Request) {
	channelURL := r.FormValue("channelURL")
	downloadMode := r.FormValue("mode")
	UpdateChannelsDatabase(channelURL)

	channelName := strings.Split(channelURL, "/")[4]
	channelType := strings.Split(channelURL, "/")[3]

	if channelType == "user" {
		fmt.Println("USER")
		videoId := GetLatestVideo(channelName, channelType)
		if downloadMode == "Video And Audio" {
			DownloadVideo(videoId)
		}
	} else if channelType == "channel" {
		fmt.Println("CHANNEL")
		videoId := GetLatestVideo(channelName, channelType)
		if downloadMode == "Video And Audio" {
			DownloadVideo(videoId)
		}
	}

	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
}
