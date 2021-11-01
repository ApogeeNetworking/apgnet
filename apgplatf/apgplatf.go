package apgplatf

import (
	"encoding/json"
	"fmt"

	"github.com/ApogeeNetworking/apgnet/apgreq"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service ...
type Service struct {
	http *apgreq.Service
}

// NewService ...
func NewService(req *apgreq.Service) *Service {
	return &Service{http: req}
}

// GetSites ...
func (s *Service) GetSites() ([]Site, error) {
	var sites []Site
	uri := "/netplatform/site"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return sites, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return sites, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&sites)
	if err != nil {
		return sites, err
	}
	return sites, nil
}

func (s *Service) GetSite(id string) (Site, error) {
	var site Site
	siteID, _ := primitive.ObjectIDFromHex(id)
	uri := fmt.Sprintf("/netplatform/site/%s", siteID)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return site, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return site, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&site)
	if err != nil {
		return site, err
	}
	return site, nil
}
