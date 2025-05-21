package models

type Bus struct {
	Id int `gorm:"primaryKey;column:Id"`
	BusNumber string `gorm:"column:BusNumber"`
	RouteId int `gorm:"column:RouteId"`
	EmployeeId int `gorm:"column:EmployeeId"`
	Compacity int `gorm:"column:Compacity"`
}

func (Bus) TableName() string { return "Buses" }


func (b *Bus) GetId() int {return b.Id}
