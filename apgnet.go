package apgnet

import (
	"github.com/ApogeeNetworking/apgnet/apgplatf"
	"github.com/ApogeeNetworking/apgnet/apgreq"
	"github.com/ApogeeNetworking/apgnet/apgsn"
	"github.com/ApogeeNetworking/apgnet/thinobot"
)

type Service struct {
	Platform *apgplatf.Service
	Snow     *apgsn.Service
	Thinobot *thinobot.Service
}

func NewService(host string, insecureSSL bool) *Service {
	req := apgreq.NewService(host, insecureSSL)
	return &Service{
		Platform: apgplatf.NewService(req),
		Thinobot: thinobot.NewService("/thino", req),
		Snow:     apgsn.NewService("/apgsn", req),
	}
}
