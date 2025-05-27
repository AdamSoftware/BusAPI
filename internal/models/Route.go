package models

import "time"

type Route struct {
	Id int `gorm:"primaryKey;column:Id"`
	DepartureTime time.Time `gorm:"column:DepartureTime"`
	GeoJson string `gorm:"column:GeoJson"`
	RouteName string `gorm:"column:RouteName"`
}

func (Route) TableName() string { return "Routes" }

func (r *Route) GetId() int {return r.Id}
