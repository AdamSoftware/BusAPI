package models

type User struct {
  UserId int `json:"id"`
  EmployeeId int `json:"employee_id"`
  EmployeeRoles int `json:"employee_roles"` 
  Username string `json:"username"`
  Password string `json:"password"`
}
