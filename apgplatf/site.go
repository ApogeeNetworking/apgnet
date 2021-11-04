package apgplatf

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ApogeeNetworking/apgnet/apgreq"
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
	uri := fmt.Sprintf("/netplatform/site/%s", id)
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

// CreateSite ...
func (s *Service) CreateSite(site Site) (Site, error) {
	var newSite Site
	data, err := json.Marshal(&site)
	if err != nil {
		return newSite, err
	}
	sr := bytes.NewReader(data)
	uri := "/netplatform/site"
	req, err := s.http.GenerateRequest(uri, "POST", sr)
	if err != nil {
		return newSite, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return newSite, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&newSite); err != nil {
		return newSite, err
	}
	return newSite, nil
}

// UpdateSite ...
func (s *Service) UpdateSite(site Site) (Site, error) {
	var updatedSite Site
	data, err := json.Marshal(&site)
	if err != nil {
		return updatedSite, err
	}
	sr := bytes.NewReader(data)
	uri := fmt.Sprintf("/netplatform/site/%s", site.ID.Hex())
	req, err := s.http.GenerateRequest(uri, "PUT", sr)
	if err != nil {
		return updatedSite, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return updatedSite, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&updatedSite); err != nil {
		return updatedSite, err
	}
	return updatedSite, nil
}

// DeleteSite ...
func (s *Service) DeleteSite(id string) (int, error) {
	uri := fmt.Sprintf("/netplatform/site/%s", id)
	req, err := s.http.GenerateRequest(uri, "DELETE", nil)
	if err != nil {
		return 0, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	return res.StatusCode, nil
}