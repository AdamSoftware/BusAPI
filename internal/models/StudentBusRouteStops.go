package models

type StudentBusRouteStops struct {
	Id int `gorm:"primaryKey;column:Id"`
	BusId int `gorm:"column:BusId"`
	RouteId int `gorm:"column:RouteId"`
	PickupStopId int `gorm:"column:PickupStopId"`
	DropoffStopId int `gorm:"column:DropoffStopId"`
	HasShownUp bool `gorm:"column:HasShownUp"`
}

func (StudentBusRouteStops) TableName() string { return "StudentBusRouteStops" }

func (s *StudentBusRouteStops) GetId() int { return s.Id }
