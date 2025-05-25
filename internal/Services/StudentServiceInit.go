package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"Bus-Backend/internal/Logging"
	"fmt"
	"net/http"
	"encoding/json"
	"net/url"
)


type mapboxGeocodeResponse struct {
    Features []struct {
        Center [2]float64 `json:"center"`
    } `json:"features"`
}

type StudentServiceInit struct {
	genericService GenericService[*models.Student]
	repo           *repository.GenericRepoInit[*models.Student]
	Studentrepo 	 *repository.StudentRepoInit
	mapBoxToken  		string 
}

func NewStudentService(repo *repository.GenericRepoInit[*models.Student], mapBoxToken string, Studentrepo *repository.StudentRepoInit) StudentService {
	return &StudentServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
		mapBoxToken:    mapBoxToken,
		Studentrepo:     Studentrepo,
	}
}

func (s *StudentServiceInit) Get() ([]*models.Student, error) {
	return s.genericService.Get()
}

// Find user by School ID
func (s *StudentServiceInit) FindBySchoolId(schoolId int) ([]*models.Student, error) {
	students, err := s.Studentrepo.FindBySchoolId(schoolId) 
	if err != nil {
		return nil, err
	}
	return students, nil
}

// Find user by ID
func (s *StudentServiceInit) FindById(id int) (*models.Student, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *StudentServiceInit) Insert(user *models.Student) (*models.Student, error) {
	lat, lon, err := s.geocodeAddress(user.Address)
	if err != nil {
		Logging.Logs.Warnf("Error geocoding address: %v", err) 
		return nil, fmt.Errorf("geocoding error: %v", err) 
	}

	// Set the latitude and longitude in the user object
	user.Latitude = lat
	user.Longitude = lon

	return s.genericService.Insert(user)
}

// Update an existing user
func (s *StudentServiceInit) Update(user *models.Student) (*models.Student, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *StudentServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}

// createing the geocodeaddress
func (s *StudentServiceInit) geocodeAddress(address string) (float64, float64, error) {
    url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s", url.QueryEscape(address), s.mapBoxToken)
    
    resp, err := http.Get(url)
    if err != nil {
        return 0, 0, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return 0, 0, fmt.Errorf("mapbox API returned status: %s", resp.Status)
    }

    var result mapboxGeocodeResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return 0, 0, err
    }

    if len(result.Features) == 0 {
        return 0, 0, fmt.Errorf("no geocode result found")
    }

    lon := result.Features[0].Center[0]
    lat := result.Features[0].Center[1]

    return lat, lon, nil
}

