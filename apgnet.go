package apgnet

import (
	"github.com/ApogeeNetworking/apgnet/apgplatf"
	"github.com/ApogeeNetworking/apgnet/apgreq"
)

type Service struct {
	Platform *apgplatf.Service
}

func NewService(host string, insecureSSL bool) *Service {
	req := apgreq.NewService(host, insecureSSL)
	return &Service{
		Platform: apgplatf.NewService(req),
	}
}
