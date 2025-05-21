package models

type EmployeeRole struct {
	Id int `gorm:"primaryKey;column:Id"`
	RoleName string `grom:"column:RoleName"`
}

func (EmployeeRole) TableName() string { return "EmployeeRoles" }


func (e *EmployeeRole) GetId() int { return e.Id}
