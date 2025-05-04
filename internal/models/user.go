package models

type User struct {
	UserId        int    `gorm:"primaryKey;column:UserID"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	EmployeeId    int    `gorm:"column:employeeId"`
	EmployeeRoles int    `gorm:"column:employeeroles"`
}

func (User) TableName() string {
	return "Users"
}

func (u *User) GetId() int {
	return u.UserId
}
