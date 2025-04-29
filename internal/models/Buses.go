package models

type Bus struct {
  Id int `json:"id"`
  BusNumber string `json:"bus_number"`
  RouteId int `json:"route_id"` 
  EmployeeId int `json:"employee_id"`
  Compacity int `json:"compacity"`
}
