package models

type RouteStops struct {
	Id int `gorm:"primaryKey;column:Id"`
	RouteId int `gorm:"column:RouteId"`
	Stopid int `gorm:"column:StopId"`
	StopOrder int `gorm:"column:StopOrder"`
}

func (RouteStops) TableName() string { return "RouteStops" }

func (r *RouteStops) GetId() int { return r.Id }


