package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cvcio/twitter"
	"io/ioutil"
	"log"
	"translatorBot/cmd/api/domain"
	"translatorBot/cmd/api/utils/apierrors"
)

type TwitterClientInterface interface {
	NewTweet(ctx context.Context, request domain.NewTweetRequest) apierrors.ApiError
	SearchTweetsFromUser(ctx context.Context, userId int64) ([]twitter.Tweet, apierrors.ApiError)
}

type twitterClient struct {
	twitter *twitter.Twitter
}

func (t *twitterClient) NewTweet(ctx context.Context, request domain.NewTweetRequest) apierrors.ApiError {
	requestJson, _ := json.Marshal(request)
	resp, err := t.twitter.GetClient().Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewReader(requestJson))
	if err != nil {
		return apierrors.NewInternalServerApiError(err.Error(), err)
	}
	log.Printf("Tweeting... Response status code: %d\n", resp.StatusCode)
	return nil
}

func (t *twitterClient) SearchTweetsFromUser(ctx context.Context, userId int64) ([]twitter.Tweet, apierrors.ApiError) {
	resp, err := t.twitter.GetClient().Get(fmt.Sprintf("https://api.twitter.com/2/users/%d/tweets?exclude=replies,retweets&max_results=10&tweet.fields=attachments,id,referenced_tweets", userId))
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(err.Error(), err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 399 {
		log.Printf("Error response from Twitter API: %s\n", string(b))
		return nil, apierrors.NewInternalServerApiError(err.Error(), err)
	}

	var responseData twitter.Data
	if err := json.Unmarshal(b, &responseData); err != nil {
		log.Println("Error unmarshal data into Twitter data struct", err)
		return nil, apierrors.NewInternalServerApiError(err.Error(), err)
	}

	var tweets []twitter.Tweet
	tweetsBytes, _ := json.Marshal(responseData.Data)
	if err := json.Unmarshal(tweetsBytes, &tweets); err != nil {
		log.Println("Error unmarshal data into Twitter tweets struct", err)
		return nil, apierrors.NewInternalServerApiError(err.Error(), err)
	}

	return tweets, nil
}
