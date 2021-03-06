package model

import "time"

type TwitterTrend struct {
	ID          int       `json:"id,omitempty"`
	Rank        int       `json:"rank,omitempty"`
	Name        string    `json:"name,omitempty"`
	Url         string    `json:"url,omitempty"`
	TweetVolume int       `json:"tweet_volume,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
