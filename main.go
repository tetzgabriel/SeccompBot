package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func main() {
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client, err := getClient(&creds)
	if err != nil {
		log.Fatal("ERROR")
	}

	tweet(client)

}

func getClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	log.Printf("User's ACCOUNT: %s\n", user.Name)
	return client, nil
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func daysUnitlSeccomp() int {
	seccompDay := date(2022, 10, 11)
	today := time.Now()

	daysUntil := seccompDay.Sub(today).Hours() / 24

	return int(daysUntil)
}

func tweet(client *twitter.Client) {

	daysUntil := daysUnitlSeccomp()
	textToTweet := fmt.Sprintf("Faltam %d dias para a SECCOMP 2022", daysUntil)

	_, _, err := client.Statuses.Update(textToTweet, nil)
	if err != nil {
		log.Printf("Error tweeting!")
		log.Fatal(err)
	}

	log.Printf("Success tweeting!")
}
