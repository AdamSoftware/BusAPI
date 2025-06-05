package dto 

type OrderStopsInfo struct {
	Id int `json:"id"`
	StopId int `json:"stopId"`
	StopName string `json:"stopName"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	StopTypeId int `json:"stopTypeId"`
	StopOrder int `json:"stopOrder"`
}
