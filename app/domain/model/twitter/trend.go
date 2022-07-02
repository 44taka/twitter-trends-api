package model

import "time"

type TwitterTrend struct {
	ID          int       `json:"id"`
	Rank        int       `json:"rank"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	TweetVolume int       `json:"tweet_volume"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
