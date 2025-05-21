package models


type School struct {
	Id int `gorm:"primaryKey;column:Id"`
	SchoolName string `gorm:"column:SchoolName"`
}


func (s *School) GetId() int {return s.Id}
