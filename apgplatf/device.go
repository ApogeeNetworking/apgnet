package apgplatf

import (
	"encoding/json"
	"fmt"
)

// GetDevices ...
func (s *Service) GetDevices() ([]Device, error) {
	var devices []Device
	uri := "/netplatform/device"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return devices, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return devices, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&devices)
	if err != nil {
		return devices, err
	}
	return devices, nil
}

// GetDevicesBySite ...
func (s *Service) GetDevicesBySite(id string) ([]Device, error) {
	var devices []Device
	uri := fmt.Sprintf("/netplatform/device/%s/site", id)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return devices, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return devices, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&devices)
	if err != nil {
		return devices, err
	}
	return devices, nil
}
