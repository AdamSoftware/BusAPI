package models


type School struct {
  SchoolId int `json:"id"`
  SchoolName string `json:"school_name"`
}


func (s *School) GetId() int {return s.SchoolId}
