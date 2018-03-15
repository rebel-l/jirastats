package search

type Request struct {
	Jql string `json:"jql"`
	StartAt int `json:"startAt"`
	MaxResults int `json:"maxResults"`
	Fields []string `json:"fields"`
}

func NewRequest(jql string) *Request {
	req := new(Request)
	req.Jql = jql
	req.StartAt = 0
	req.MaxResults = 10

	fields := make([]string, 2)
	fields[0] = "summary"
	fields[1] = "status"
	req.Fields = fields
	return  req
}