package model

import "encoding/json"

type Student struct {
	ID         int64
	Name       string
	Email      string
	GPA        float64
	Credit     int
	Department int64
	Advisor    int64
}

func (s *Student) ToJSON() string {
	student, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return ""
	}
	return string(student)
}
