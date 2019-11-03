package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetLatestVideo(channelName, channelType string) string {
	if channelType == "user" {
		uploadsId := getUserUploadsID(channelName)
		requestURL := API_ENDPOINT_PLAYLIST + uploadsId + "&maxResults=2" + "&key=" + API_KEY

		resp, err := http.Get(requestURL)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		var video NameBody

		json.Unmarshal([]byte(string(body)), &video)

		return video.Items[0].Snippet.ResourceID.VideoID
	}
	requestURL := API_ENDPOINT_ID + channelName + "&" + MAX_RESULTS + "&" + ORDER_BY + "&" + TYPE + "&key=" + API_KEY

	resp, err := http.Get(requestURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var video ChannelBody

	json.Unmarshal([]byte(string(body)), &video)

	return video.Items[0].ID.VideoID
}

func getUserUploadsID(channelName string) string {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	requestURL := API_ENDPOINT_NAME + channelName + "&key=" + API_KEY

	resp, err := http.Get(requestURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user UserInformation

	json.Unmarshal([]byte(string(body)), &user)

	return user.Items[0].ContentDetails.RelatedPlaylists.Uploads
}
