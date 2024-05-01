// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// )

// const BEARER_TOKEN = "AAAAAAAAAAAAAAAAAAAAAKb6tQEAAAAA%2FKjz%2B9VxdlhhA8geqTiSxL5MV%2B4%3DlHATvPk2FpNp1FnqsHfQJ1SuKHESBG4caNWJAwVhlKAtgJQNEh"

// func main() {
// 	app := fiber.New()

// 	// Define the endpoint that calls Twitter's streaming API
// 	app.Get("/stream-twitter", func(c *fiber.Ctx) error {
// 		// Set up the request to Twitter's streaming API
// 		url := "https://api.twitter.com/2/tweets/search/stream"
// 		method := "GET"

// 		client := &http.Client{}
// 		req, err := http.NewRequest(method, url, nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		bearerToken := BEARER_TOKEN
// 		req.Header.Add("Authorization", "Bearer "+bearerToken)

// 		// Make the request to Twitter's API
// 		res, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		defer res.Body.Close()

// 		// Read the response body
// 		body, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}

// 		// Log the response body
// 		log.Println(string(body))

// 		// Send the response back to the client
// 		return c.SendString(string(body))
// 	})

// 	app.Post("/add-rules", func(c *fiber.Ctx) error {
// 		url := "https://api.twitter.com/2/tweets/search/stream/rules"
// 		method := "POST"

// 		// Define the rules you want to add
// 		rules := []byte(`{
// 			"add": [
// 				{"value": "dog has:images", "tag": "dog pictures"},
// 				{"value": "cat has:images -grumpy"}
// 			]
// 		}`)

// 		client := &http.Client{}
// 		req, err := http.NewRequest(method, url, bytes.NewBuffer(rules))
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		bearerToken := BEARER_TOKEN // Replace with your actual bearer token
// 		req.Header.Add("Authorization", "Bearer "+bearerToken)
// 		req.Header.Add("Content-Type", "application/json")

// 		res, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		defer res.Body.Close()

// 		body, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}

// 		log.Println(string(body))
// 		return c.SendString(string(body))
// 	})

// 	app.Get("/get-rules", func(c *fiber.Ctx) error {
// 		url := "https://api.twitter.com/2/tweets/search/stream/rules"
// 		method := "GET"

// 		client := &http.Client{}
// 		req, err := http.NewRequest(method, url, nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		bearerToken := BEARER_TOKEN // Replace with your actual bearer token
// 		req.Header.Add("Authorization", "Bearer "+bearerToken)

// 		res, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		defer res.Body.Close()

// 		body, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}

// 		log.Println(string(body))
// 		return c.SendString(string(body))
// 	})

// 	// Start the server on port 3000
// 	log.Fatal(app.Listen(":3000"))
// }
