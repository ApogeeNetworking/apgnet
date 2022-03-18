package apgers

import (
	"time"

	"github.com/ApogeeNetworking/apgnet/apgreq"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type ErsAlert struct {
	Alert                string    `json:"alert"`
	NodeID               int       `json:"nodeId"`
	NodeName             string    `json:"nodeName"`
	Intf                 string    `json:"interface"`
	Organization         string    `json:"organization"`
	CarrierName          string    `json:"carrierName"`
	CircuitID            string    `json:"circuitId"`
	CarrierSupportNumber string    `json:"carrierSupportNumber"`
	WanIpVrf             string    `json:"wanIpVrf"`
	GlueIp               string    `json:"glueIp"`
	RemoteNodeName       string    `json:"remoteNodeName"`
	RemoteNodeInterface  string    `json:"remoteNodeIntf"`
	TriggeredAt          time.Time `json:"triggeredAt"`
}

// Job ...
type Job struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SiteID          primitive.ObjectID `bson:"siteId" json:"siteId"`
	IncID           string             `bson:"incId" json:"incId"`
	IncURL          string             `bson:"incUrl" json:"incUrl"`
	Inc             string             `bson:"inc" json:"inc"`
	CarrierIncID    string             `bson:"carrierIncId" json:"carrierIncId"`
	CarrierIncURL   string             `bson:"carrierIncUrl" json:"carrierIncUrl"`
	AffectedCircuit string             `bson:"affectedCircuit" json:"affectedCircuit"`
	ReasonForOutage string             `bson:"rfo" json:"rfo"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
	LastUpdated     time.Time          `bson:"lastUpdated" json:"lastUpdated"`
	Active          bool               `bson:"active" json:"active"`
	Actions         []HistoryItem      `bson:"actions" json:"actions"`
}

type HistoryItem struct {
	Action    string `bson:"action" json:"action"`
	Timestamp int64  `bson:"timestamp" json:"timestamp"`
	Details   string `bson:"details" json:"details"`
	PassFail  string `bson:"passFail" json:"passFail"`
}
