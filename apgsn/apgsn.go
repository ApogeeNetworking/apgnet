package apgsn

import (
	"bytes"
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
func NewService(baseURL string, req *apgreq.Service) *Service {
	return &Service{
		baseURL: baseURL,
		http:    req,
	}
}

// GetIncidentByID ...
func (s *Service) GetIncidentByID(sysID string) (Incident, error) {
	var incident Incident
	uri := fmt.Sprintf("%s/incidents/%s", s.baseURL, sysID)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return incident, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return incident, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&incident); err != nil {
		return incident, err
	}
	return incident, nil
}

// GetIncidentsByCustomerID ...
func (s *Service) GetIncidentsByCustomerID(customerID string) ([]Incident, error) {
	var incidents []Incident
	uri := fmt.Sprintf("%s/incidents?customerId=%s", s.baseURL, customerID)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return incidents, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return incidents, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&incidents); err != nil {
		return incidents, err
	}
	return incidents, nil
}

// CreateIncidentFlow creates Incident and WorkOrder for AP Deployment Issues
func (s *Service) CreateIncidentFlow(incReq *CreateIncidentReq) (Incident, error) {
	var incident Incident
	data, err := json.Marshal(incReq)
	if err != nil {
		return incident, err
	}
	sr := bytes.NewReader(data)
	uri := fmt.Sprintf("%s/internal/incidents", s.baseURL)
	req, err := s.http.GenerateRequest(uri, "POST", sr)
	if err != nil {
		return incident, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return incident, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&incident); err != nil {
		return incident, err
	}
	return incident, nil
}

// GetWorkOrderByID ...
func (s *Service) GetWorkOrderByID(id string) (WorkOrder, error) {
	var workOrder WorkOrder
	uri := fmt.Sprintf("%s/workorder/%s", s.baseURL, id)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return workOrder, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return workOrder, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&workOrder); err != nil {
		return workOrder, err
	}
	return workOrder, nil
}

// GetCustomerByID ...
func (s *Service) GetCustomerByID(id, domainName string) (CustomerAccount, error) {
	var customer CustomerAccount
	uri := fmt.Sprintf("%s/customer?id=%s&domain=%s", s.baseURL, id, domainName)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return customer, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return customer, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&customer); err != nil {
		return customer, err
	}
	return customer, nil
}

// GetCustomerList ...
func (s *Service) GetCustomerList() (CustomerName, error) {
	var customers CustomerName
	uri := fmt.Sprintf("%s/customer/summary", s.baseURL)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return customers, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return customers, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&customers); err != nil {
		return customers, err
	}
	return customers, nil
}

// GetUser ...
func (s *Service) GetUser(email string) (User, error) {
	var user User
	uri := fmt.Sprintf("%s/user?email=%s", s.baseURL, email)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return user, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return user, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}
