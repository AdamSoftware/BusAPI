package models

type RouteGroup struct {
	Id int `gorm:"primaryKey;column:Id"`
	Name string `gorm:"column:Name"`
}

func (RouteGroup) TableName() string { return "RouteGroups" }

func (r *RouteGroup) GetId() int { return r.Id }
