package thinobot

import (
	"log"
	"time"

	"github.com/ApogeeNetworking/apgnet/apgplatf"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bot ...
type Bot struct {
	ID              string             `bson:"_id,omitempty" json:"id,omitempty"`
	SiteID          primitive.ObjectID `bson:"siteId" json:"siteId"`
	Name            string             `bson:"name" json:"name"`
	State           string             `bson:"state" json:"state"`
	Campus          string             `bson:"campus" json:"campus"`
	WLCIP           string             `bson:"controllerIp" json:"controllerIp"`
	SpreadsheetID   string             `bson:"spreadSheetId" json:"spreadSheetId"`
	StagingSheetID  int64              `bson:"stagingSheetId" json:"stagingSheetId"`
	DeploymtSheetID int64              `bson:"deploymtSheetId" json:"deploymtSheetId"`
	ChannelID       string             `bson:"channelId" json:"channelId"`
	Expiration      time.Time          `bson:"exp" json:"expiration"`
	DeploymentType  string             `bson:"deploymentType" json:"deploymentType"`
	ProjectActive   bool               `bson:"projectActive" json:"projectActive"`
	Watched         bool               `bson:"watched" json:"watched"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
	SpreadsheetLink string             `bson:"spreadSheetLink" json:"spreadSheetLink"`
	ProjectTeam     []string           `bson:"projectTeam" json:"projectTeam"`
	// Used for Ruckus SmartZones currently only 8_1 OR 10_0
	ApiVersion string `bson:"apiVersion" json:"apiVersion"`
}

// AccessPoint ...
type AccessPoint struct {
	Name         string            `json:"name"`
	MacAddr      string            `json:"macAddr"`
	Serial       string            `json:"serialNumber"`
	Model        string            `json:"model"`
	Status       string            `json:"status"`
	Group        string            `json:"group"`
	WlcIP        string            `json:"controllerIp"`
	ServiceTag   string            `json:"serviceTag"`
	Staged       bool              `json:"staged"`
	Row          int               `json:"row"`
	Installed    bool              `json:"installed"`
	GoodPoll     int64             `json:"goodPoll"`
	Vendor       string            `json:"vendor"`
	RfProfile    string            `json:"rfTag"`
	SiteProfile  string            `json:"siteTag"`
	RemoteSwitch string            `json:"remoteSwitch"`
	RemoteIntf   string            `json:"remoteIntf"`
	FactoryReset bool              `json:"factoryReset"`
	VlanID       int               `json:"vlanId"`
	Notes        string            `json:"notes"`
	Building     string            `json:"bldg"`
	Room         string            `json:"room"`
	Location     apgplatf.Location `json:"location"`
}

// RouterParams ...
type RouterParams struct {
	GoogService *GDriveService
	Logger      *log.Logger
	SSHUser     string
	SSHPassword string
	SSHEnable   string
	RksUser     string
	RksPassword string
	NetPlatURI  string
	SnServURI   string
}

// ApServiceParams ...
type ApServiceParams struct {
	Bot            *Bot
	Logger         *log.Logger
	GoogService    *GDriveService
	User           string
	Password       string
	EnablePW       string
	NormalizeMac   func(string) string
	DeploymentType string
	AsyncCount     int
}

// ApStatus ...
type ApStatus struct {
	ApName     string `json:"apName"`
	ApMacAddr  string `json:"apMac"`
	Speed      string `json:"speed"`
	Duplex     string `json:"duplex,omitempty"`
	RemoteSw   string `json:"remoteSwitch"`
	RemoteIntf string `json:"remoteSwitchport"`
	Polls      int64  `json:"polls,omitempty"`
	Status     string `json:"status,omitempty"`
	Row        int    `json:"row,omitempty"`
}

// ApSummary ...
type ApSummary struct {
	Planned   int    `json:"plannedAps"`
	Installed string `json:"installedAps"`
	FullOp    string `json:"fullOpAps"`
	Degraded  string `json:"degradedAps"`
	Down      string `json:"downAps"`
}

type GDriveService struct {
	SheetSrv   *sheets.Service
	DriveSrv   *drive.Service
	WebHookURI string
	JwtPath    string
	Campus     string
	State      string
	SheetGrid  struct {
		headerEndCol int64
		boolStartCol int64
	}
}
