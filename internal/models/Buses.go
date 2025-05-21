package models

type Bus struct {
	Id int `gorm:"primaryKey;column:Id"`
	BusNumber string `gorm:"column:BusNumber"`
	RouteId int `gorm:"column:RouteId"`
	EmployeeId int `gorm:"column:EmployeeId"`
	Capacity int `gorm:"column:Capacity"`
}

func (Bus) TableName() string { return "Buses" }


func (b *Bus) GetId() int {return b.Id}
