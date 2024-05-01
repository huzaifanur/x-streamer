// app.Post("/add-rules", func(c *fiber.Ctx) error {
// 	url := "https://api.twitter.com/2/tweets/search/stream/rules"
// 	method := "POST"

// 	// Define the rules you want to add
// 	rules := []byte(`{
// 		"add": [
// 			{"value": "dog has:images", "tag": "dog pictures"},
// 			{"value": "cat has:images -grumpy"}
// 		]
// 	}`)

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(rules))
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	bearerToken := BEARER_TOKEN // Replace with your actual bearer token
// 	req.Header.Add("Authorization", "Bearer "+bearerToken)
// 	req.Header.Add("Content-Type", "application/json")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	log.Println(string(body))
// 	return c.SendString(string(body))
// })

// app.Get("/get-rules", func(c *fiber.Ctx) error {
// 	url := "https://api.twitter.com/2/tweets/search/stream/rules"
// 	method := "GET"

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	bearerToken := BEARER_TOKEN // Replace with your actual bearer token
// 	req.Header.Add("Authorization", "Bearer "+bearerToken)

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	log.Println(string(body))
// 	return c.SendString(string(body))
// })



// func main() {
//     app := fiber.New()

//     app.Get("/stream-twitter", func(c *fiber.Ctx) error {
       
//         fmt.Println("received request")
//         url := "https://api.twitter.com/2/tweets/search/stream"
//         method := "GET"

//         client := &http.Client{}
//         req, err := http.NewRequest(method, url, nil)
//         if err != nil {
//             fmt.Println(err)
//             return err
//         }
//         bearerToken := BEARER_TOKEN
//         req.Header.Add("Authorization", "Bearer "+bearerToken)

//         // Make the request to Twitter's API
//         res, err := client.Do(req)
//         if err != nil {
//             fmt.Println(err)
//             return err
//         }
//         defer res.Body.Close()

//         // Create a channel to receive data from the response body
//         dataChan := make(chan string)

//         // Start a goroutine to read from the response body and send data to the channel
//         go func() {
//             defer close(dataChan)
//             scanner := bufio.NewScanner(res.Body)
//             for scanner.Scan() {
//                 dataChan <- scanner.Text()
//             }
//             if err := scanner.Err(); err != nil {
//                 log.Println("Error reading from stream:", err)
//             }
//         }()

//         // Start another goroutine to process the data received from the channel
//         go func() {
//             for data := range dataChan {
//                 log.Println(data)
//             }
//         }()

//         // Wait for a few seconds before returning the response
//         time.Sleep(10 * time.Second)

//         return nil // Return nil or an appropriate Fiber error response
//     })

//     app.Listen(":8080")
// }