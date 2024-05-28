package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

type TweetData struct {
	Data struct {
		ID                  string   `json:"id"`
		EditHistoryTweetIDs []string `json:"edit_history_tweet_ids"`
		Text                string   `json:"text"`
	} `json:"data"`
}

func streamTwitterData(dataChan chan<- TweetData) error {
	BEARER_TOKEN := os.Getenv("BEARER_TOKEN")
	// url := "https://api.twitter.com/2/tweets/search/stream"
	url := `https://api.twitter.com/2/tweets/search/stream?expansions=author_id,in_reply_to_user_id,geo.place_id,referenced_tweets.id,entities.mentions.username&place.fields=country,full_name,name&tweet.fields=author_id,created_at,geo,id,in_reply_to_user_id,lang,text&user.fields=description,name,username,location
	`

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+BEARER_TOKEN)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			break // Exit the loop on error
		}

		var tweet TweetData
		err = json.Unmarshal(line, &tweet)
		if err != nil {
			log.Println(err)
			continue // Skip on unmarshaling error
		}

		dataChan <- tweet
	}
	return nil
}

func PublishTweetData(dataChan <-chan TweetData) {
	PROJECT_ID := os.Getenv("PROJECT_ID")
	TOPIC_ID := os.Getenv("TOPIC_ID")

	// Setup publisher
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	topic := client.Topic(TOPIC_ID)
	defer topic.Stop()

	// Publish tweets from the channel
	for tweet := range dataChan {
		tweetBytes := []byte(tweet.Data.Text)

		result := topic.Publish(ctx, &pubsub.Message{
			Data: tweetBytes,
		})

		// Block until the result is returned and a server-generated
		// ID is returned for the published message.
		id, err := result.Get(ctx)
		if err != nil {
			log.Fatalf("Failed to publish: %v", err)
		}
		log.Printf("Published a message; msg ID: %v\n", id)
	}
}

func saveTwitterData() error {
	BEARER_TOKEN := os.Getenv("BEARER_TOKEN")
	url := "https://api.twitter.com/2/tweets/search/stream?expansions=author_id,in_reply_to_user_id,geo.place_id,referenced_tweets.id,entities.mentions.username&place.fields=country,full_name,name&tweet.fields=author_id,created_at,geo,id,in_reply_to_user_id,lang,text&user.fields=description,name,username,location"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+BEARER_TOKEN)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	file, err := os.Create("twitter_data.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	reader := bufio.NewReader(res.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break // End of stream
			}
			log.Println(err)
			break // Exit the loop on error
		}

		_, err = writer.Write(line)
		if err != nil {
			log.Println(err)
			continue // Skip on write error
		}
		writer.Flush() // Flush after each line to ensure all data is written
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	saveTwitterData()
	// dataChan := make(chan TweetData, 100)

	// // Start the logging goroutine first
	// go PublishTweetData(dataChan)

	// // Then start streaming Twitter data
	// err = streamTwitterData(dataChan)
	// if err != nil {
	// 	log.Println(err)
	// }
}
