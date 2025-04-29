package models

type Route struct {
  Id int `json:"id"`
  BusId int `json:"bus_id"`
  GeoJson string `json:"geo_json"`
  RouteName string `json:"route_name"`
}
