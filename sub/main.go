package main

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PROJECT_ID := os.Getenv("PROJECT_ID")
	SUBSCRIPTION_ID := os.Getenv("SUBSCRIPTION_ID")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(SUBSCRIPTION_ID)
	var mu sync.Mutex
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()

		log.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack() // Acknowledge that we've consumed the message.
	})
	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}
}
