package models


type Employee struct {
	Id int `gorm:"primaryKey;column:Id"`
	FirstName string `gorm:"column:FirstName"`
	LastName string `gorm:"column:LastName"`
	Phone string `gorm:"column:Phone"`
	Email string `gorm:"column:Email"`
}

func (Employee) TableName() string { return "Employees" }

func (e *Employee) GetId() int {return e.Id}
