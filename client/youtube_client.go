package client

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	query      = flag.String("query", "Cricket", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

const developerKey = "AIzaSyAdEZoZtIMjK0O1NM9b1vAUOcf5fMS_scs"

func GetYoutubeClient() {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	parts := []string{"id", "snippet"}
	publishedAfter := time.Now().AddDate(0, 0, -1).UTC().Format(time.RFC3339)
	fmt.Println(publishedAfter)
	// Make the API call to YouTube.
	call := service.Search.List(parts).
		Type("video").
		Order("date").
		PublishedAfter(publishedAfter).
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		fmt.Println("Youtube API call failed")
		return
	}

	videos := make(map[string]string)

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.PublishedAt
		}
	}

	printIDs("Videos", videos)
}

func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
