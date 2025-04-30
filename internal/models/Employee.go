package models


type Employee struct {
  EmployeeId int `json:"id"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Phone string `json:"phone"`
  Email string `json:"email"`
}
