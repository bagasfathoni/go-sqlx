package model

import "encoding/json"

type Student struct {
	ID         int32
	Name       string
	Email      string
	Department Department
	GPA        float64
	Credit     int
	Advisor    Advisor
}

func (s *Student) ToJSON() string {
	student, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return ""
	}
	return string(student)
}
