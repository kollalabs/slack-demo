package main

import (
	"context"
	"log"
	"os"

	"github.com/kollalabs/sdk-go/kc"
	"github.com/slack-go/slack"
)

func main() {
	// Get api key from environment variable
	apiKey := os.Getenv("KC_API_KEY")
	ctx := context.Background()

	// Create a new client
	kolla, err := kc.New(apiKey)
	if err != nil {
		log.Fatalf("unable to load kolla connect client: %s\n", err)
	}

	creds, err := kolla.Credentials(ctx, "slack", "CONSUMER_ID") // Use consumer ID set in consumer token
	if err != nil {
		log.Fatalf("unable to load consumer credentials: %s\n", err)
	}

	slackapi := slack.New(creds.Token)
	_, _, err = slackapi.PostMessage("general", slack.MsgOptionText("Hello world! (Send with Kolla managed token)", false))
	if err != nil {
		log.Fatalf("unable to post slack message: %s\n", err)
	}
}
