package model

import "encoding/json"

type Advisor struct {
	ID    int64
	Name  string
	Email string
}

func (a *Advisor) ToJSON() string {
	advisor, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return ""
	}
	return string(advisor)
}
