package models

type User struct {
	UserId        int    `gorm:"primaryKey;column:UserID"`
	Username      string `gorm:"column:Username"`
	Password      string `gorm:"column:Password"`
	EmployeeID    int    `gorm:"column:EmployeeID"`
	EmployeeRoles int    `gorm:"column:EmployeeRoles"`
}

func (User) TableName() string {
	return "Users"
}

func (u *User) GetId() int {
	return u.UserId
}
