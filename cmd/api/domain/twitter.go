package domain

type NewTweetRequest struct {
	Text         string `json:"text"`
	QuoteTweetID string `json:"quote_tweet_id,omitempty"`
}
