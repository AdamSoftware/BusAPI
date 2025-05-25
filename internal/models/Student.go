package models

type Student struct {
  Id int `gorm:"primaryKey;column:Id"`
	SchoolId int `gorm:"column:SchoolId"`
	FirstName string `gorm:"column:FirstName"`
	LastName string `gorm:"column:LastName"`
	Photo string `gorm:"column:Photo"`
	Address string `gorm:"column:Address"`
	Latitude float64 `gorm:"column:Latitude"`
	Longitude float64 `gorm:"column:Longitude"`
}

func (Student) TableName() string {
	return "Students"
}

func (s *Student) GetId() int {
  return s.Id
}
