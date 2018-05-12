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
		ds.Value = data.Format(DateFormatInternal)
	} else {
		ds.Name = data.Format(DateFormatDisplay)
		ds.Value = data.Format(DateFormatInternal)
	}
	return ds
}
