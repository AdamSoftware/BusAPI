package models

type Student struct {
  Id int `gorm:"primaryKey;column:Id"`
	SchoolId int `gorm:"column:SchoolId"`
	FirstName string `gorm:"column:FirstName"`
	LastName string `gorm:"column:LastName"`
	Photo string `grom:"column:Photo"`
}

func (Student) TableName() string {
	return "Students"
}

func (s *Student) GetId() int {
  return s.Id
}
