package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
	url := "https://api.twitter.com/2/tweets/search/stream"
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

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataChan := make(chan TweetData)

	go func() {
		err := streamTwitterData(dataChan)
		if err != nil {
			log.Println(err)
		}
	}()

	for tweet := range dataChan {
		fmt.Printf("Tweet: %+v\n", tweet)
	}
}
