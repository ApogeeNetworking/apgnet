package apgnet

import (
	"github.com/ApogeeNetworking/apgnet/apgplatf"
	"github.com/ApogeeNetworking/apgnet/apgreq"
	"github.com/ApogeeNetworking/apgnet/apgsn"
)

type Service struct {
	Platform *apgplatf.Service
	Snow     *apgsn.Service
}

func NewService(host string, insecureSSL bool) *Service {
	req := apgreq.NewService(host, insecureSSL)
	return &Service{
		Platform: apgplatf.NewService(req),
		Snow:     apgsn.NewService("/apgsn", req),
	}
}
