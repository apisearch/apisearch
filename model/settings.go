package model

type Settings struct {
	UserId     int
	FeedUrl    string `json:"feedUrl"`
	FeedFormat string `json:"feedFormat"`
	Frequency  int    `json:"frequency"`
}
