package models

type StopType struct {
	Id int `gorm:"primaryKey;column:Id"`
	TypeName string `gorm:"column:TypeName"`
}

func (StopType) TableName() string { return "StopTypes" }

func (s *StopType) GetId() int { return s.Id }
