package services


import (
  "errors"
  "Bus-Backend/internal/models"
  "gorm.io/gorm"
)


type GenericRepoInit[T models.GenericModel] struct {
  db *gorm.DB 
}


// createing the generic repo for all the different ones
