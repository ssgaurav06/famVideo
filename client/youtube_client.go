package client

import (
	"fam/storage"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type YoutubeClient interface {
	GetClient()
}
type youtubeClient struct {
	videoDataStore storage.VideoDataStorage
}

func NewYoutubeClient(vds storage.VideoDataStorage) YoutubeClient {
	return &youtubeClient{videoDataStore: vds}
}

var (
	query          = flag.String("query", "Cricket", "Search term")
	maxResults     = flag.Int64("max-results", 25, "Max YouTube results")
	publishedAfter = time.Now().AddDate(0, 0, -1).UTC().Format(time.RFC3339)
)

func (yc youtubeClient) GetClient() {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("DEVELOPER_KEY")},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	parts := []string{"id", "snippet"}

	for {
		// Make the API call to YouTube.
		call := service.Search.List(parts).
			Type("video").
			Order("date").
			PublishedAfter(publishedAfter).
			Q(*query).
			MaxResults(*maxResults)
		response, err := call.Do()
		if err != nil {
			fmt.Println("Youtube API call failed: ", err)
			return
		}
		// Clear existing DB
		if err := yc.videoDataStore.ClearDB(); err != nil {
			fmt.Println("DB clear failed")
			return
		}

		//Insert video data into DB
		fmt.Println("Insert Data into DB")
		for _, item := range response.Items {
			if err := yc.videoDataStore.Insert(item); err != nil {
				fmt.Println("Item insertion failed")
				return
			}
		}
		time.Sleep(10 * time.Second)
	}
}
