package response

type TableData struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

func NewTableData(name string, value int) *TableData {
	td := new(TableData)
	td.Name = name
	td.Value = value
	return td
}
