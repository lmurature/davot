package app

import (
	"context"
	"fmt"
	"github.com/cvcio/twitter"
	"log"
	"net/http"
	"time"
	"translatorBot/cmd/api/clients"
	"translatorBot/cmd/api/domain"
	"translatorBot/cmd/api/providers"
)

const (
	davooXeneizeUserID = 836969699665670145
	davooDeutschUserID = 1488237896402309122
)

func StartDabot() {
	setProperLastProcessedTweet()
	for {
		process()
		time.Sleep(1 * time.Minute)
	}
}

func process() {
	davosTweets, err := providers.TwitterProvider.SearchTweetsFromUser(context.Background(), davooXeneizeUserID)
	if err != nil {
		log.Println(err)
		return
	}
	defer clients.CacheClient.SetLastProcessedTweet(davosTweets[0].ID)

	cachedProcessedTw, err := clients.CacheClient.GetLastProcessedTweet()
	if err != nil {
		log.Println(err)
		if err.Status() != http.StatusNotFound {
			return
		}
	}

	indexLimit := len(davosTweets) - 1
	if cachedProcessedTw != nil {
		for i, tw := range davosTweets {
			if tw.ID == *cachedProcessedTw {
				indexLimit = i
			}
		}
	}
	quoteTweets(davosTweets[:indexLimit])
}

func quoteTweets(tws []twitter.Tweet) {
	for i := range tws {
		reverseIndex := len(tws) - 1 - i
		translation, err := providers.TranslationProvider.Translate(context.Background(), domain.TranslationRequest{
			Quote:  tws[reverseIndex].Text,
			Source: "es",
			Target: "de",
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		log.Printf("About to tweet quote reply to '%s' with translation '%s'\n", tws[reverseIndex].Text, translation.TranslatedText)
		req := domain.NewTweetRequest{
			Text:         translation.TranslatedText,
			QuoteTweetID: tws[reverseIndex].ID,
		}
		err = providers.TwitterProvider.NewTweet(context.Background(), req)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func setProperLastProcessedTweet() {
	botTweets, err := providers.TwitterProvider.SearchTweetsFromUser(context.Background(), davooDeutschUserID)
	if err != nil {
		log.Println(err)
		return
	}
	if len(botTweets) > 0 && len(botTweets[0].ReferencedTweets) > 0 {
		referenceID := botTweets[0].ReferencedTweets[0].ID
		log.Printf("About to set proper last processed tweet: %s\n", referenceID)
		_ = clients.CacheClient.SetLastProcessedTweet(referenceID)
	}
}
