package providers

import (
	"github.com/cvcio/twitter"
	"github.com/go-resty/resty/v2"
	"translatorBot/cmd/api/config"
)

var (
	TranslationProvider TranslationProviderInterface
	TwitterProvider     TwitterClientInterface
)

func InitializeProviders() {
	translationClient := resty.New()
	TranslationProvider = &translationProvider{
		client: translationClient,
	}

	client, err := twitter.NewTwitterWithContext(config.GetConsumerKey(), config.GetConsumerSecret(), config.GetAccessToken(), config.GetAccessTokenSecret())
	if err != nil {
		panic(err)
	}
	TwitterProvider = &twitterClient{
		twitter: client,
	}
}
