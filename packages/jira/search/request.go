package search

import "encoding/json"

type Request struct {
	Jql        string   `json:"jql"`
	StartAt    int      `json:"startAt"`
	MaxResults int      `json:"maxResults"`
	Fields     []string `json:"fields"`
}

func NewRequest(jql string) *Request {
	req := new(Request)
	req.Jql = jql
	req.StartAt = 0
	req.MaxResults = 10

	fields := make([]string, 8)
	fields[0] = "summary"
	fields[1] = "status"
	fields[2] = "components"
	fields[3] = "labels"
	fields[4] = "issuetype"
	fields[5] = "priority"
	fields[6] = "updated"
	fields[7] = "created"
	req.Fields = fields
	return req
}

func (r *Request) Next() {
	r.StartAt += r.MaxResults
}

func NewRequestFromJson(data []byte) (req *Request, err error) {
	req = new(Request)
	err = json.Unmarshal(data, req)
	return
}
