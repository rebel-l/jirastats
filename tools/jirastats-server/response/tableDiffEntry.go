package response

type TableDiffEntry struct {
	Name string `json:"name"`
	Left int `json:"left"`
	Right int `json:"right"`
	Diff int `json:"diff"`
}
