package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// All tokens and keys needed for authentication against twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// Return twitter client which we can use to send tweets
// Takes in a pointer to Credential needed to authenticate
func getClient(creds *Credentials) (*twitter.Client, error) {
	// Passing consumer key (API Key) and  Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Passing Access Token and Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// Retrieve the user and verify the credentials
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}
func main() {
	fmt.Println("Go-Twitter Bot")

	// load .env file from given path
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	tweet, resp, err := client.Statuses.Update("A Test Tweet", nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}
