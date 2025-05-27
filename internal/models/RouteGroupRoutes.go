// junc table for RouteGroup and Route 

package models

type RouteGroupRoutes struct {
	Id int `gorm:"primaryKey;column:Id"`
	RouteGroupId int `gorm:"column:RouteGroupId"`
	RouteId int `gorm:"column:RouteId"`
}

func (RouteGroupRoutes) TableName() string { return "RouteGroupRoutes" }

func (r *RouteGroupRoutes) GetId() int { return r.Id }
