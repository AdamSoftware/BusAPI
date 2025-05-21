package models

type Route struct {
	Id int `gorm:"primaryKey;column:Id"`
	BusId int `gorm:"column:BusId"`
	GeoJson string `gorm:"column:GeoJson"`
	RouteName string `gorm:"column:RouteName"`
}

func (Route) TableName() string { return "Routes" }

func (r *Route) GetId() int {return r.Id}
