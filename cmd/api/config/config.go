package config

import "os"

var (
	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string
)

func InitializeConfigs() {
	consumerKey = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")
	accessToken = os.Getenv("ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
}

func GetConsumerKey() string {
	return consumerKey
}

func GetConsumerSecret() string {
	return consumerSecret
}

func GetAccessToken() string {
	return accessToken
}

func GetAccessTokenSecret() string {
	return accessTokenSecret
}
