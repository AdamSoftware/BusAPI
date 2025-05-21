package models

type User struct {
	Id        int    `gorm:"primaryKey;column:Id"`
	Username  string `gorm:"column:Username"`
	Password  string `gorm:"column:Password"`
	EmployeeID int    `gorm:"column:EmployeeID"`
	EmployeeRoleId int    `gorm:"column:EmployeeRoleId"`
}

func (User) TableName() string {
	return "Users"
}

func (u *User) GetId() int {
	return u.Id
}
