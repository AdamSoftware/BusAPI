package models

type BusRouteGroup struct {
	Id int 	`gorm:"primaryKey;column:Id"`
	BusId int `gorm:"column:BusId"`
	RouteGroupId int `gorm:"column:RouteGroupId"`
}

func (BusRouteGroup) TableName() string { return "BusRouteGroups" }

func (b *BusRouteGroup) GetId() int { return b.Id }
