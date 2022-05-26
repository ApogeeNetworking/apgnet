package apgrxg

import (
	"encoding/json"
	"fmt"

	"github.com/ApogeeNetworking/apgnet/apgreq"
)

// Service ...
type Service struct {
	baseURL string
	http    *apgreq.Service
}

// NewService ...
func NewService(uri string, req *apgreq.Service) *Service {
	return &Service{
		baseURL: uri,
		http:    req,
	}
}

// GetRxgs ...
func (s *Service) GetRxgs() ([]RxG, error) {
	var rxgs []RxG
	uri := fmt.Sprintf("%s/site", s.baseURL)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return rxgs, nil
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return rxgs, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&rxgs); err != nil {
		return rxgs, err
	}
	return rxgs, nil
}
