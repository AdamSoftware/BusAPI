package models

type Stops struct {
	Id int `gorm:"primaryKey;column:Id"`
	StopName string `gorm:"column:StopName"`
	Latitude float64 `gorm:"column:Latitude"`
	Longitude float64 `gorm:"column:Longitude"`
	StopTypeId int `gorm:"column:StopTypeId"`
}

func (Stops) TableName() string { return "Stops" }

func (s *Stops) GetId() int { return s.Id }
