package models

type Student struct {
  Id int `json:"id"`
  SchoolId int `json:"school_id"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Photo string `json:"photo"` 
}
