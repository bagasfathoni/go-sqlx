package model

import "encoding/json"

type Department struct {
	ID     int64
	Name   string
	School string
}

func (d *Department) ToJSON() string {
	depar, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		return ""
	}
	return string(depar)
}
