package response

import "time"

type DateSelect struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

func NewDateSelect(data time.Time) *DateSelect {
	ds := new(DateSelect)
	if data.IsZero() {
		ds.Name = "actual"
	} else {
		ds.Name = data.Format(dateFormatDisplay)
		ds.Value = data.Format(dateFormatInternal)
	}
	return ds
}
