package handler

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/services"
	"encoding/json"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)


type GenericHandler[T any] struct {
	genericService services.GenericService[T]
}


