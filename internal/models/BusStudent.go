package models

type BusStudent struct {
	Id int `gorm:"primaryKey;column:Id"`
	BusId int `gorm:"column:BusId"`
	StudentId int `gorm:"column:StudentId"`
}

func (BusStudent) TableName() string { return "BusStudents" }

func (b *BusStudent) GetId() int {return b.Id}
