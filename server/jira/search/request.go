package search

type Request struct {
	Jql string `json:"jql"`
	StartAt int `json:"startAt"`
	MaxResults int `json:"maxResults"`
}
