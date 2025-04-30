package models

type Route struct {
  RouteId int `json:"id"`
  BusId int `json:"bus_id"`
  GeoJson string `json:"geo_json"`
  RouteName string `json:"route_name"`
}
