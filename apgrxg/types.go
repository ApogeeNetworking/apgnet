package apgrxg

import (
	"log"
	"time"

	"github.com/ApogeeNetworking/rgnets"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OssSite ...
type RxG struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SiteID    primitive.ObjectID `bson:"siteId" json:"siteId"`
	DbName    string             `json:"dbName" bson:"dbName"`
	Host      string             `json:"dbHost" bson:"dbHost"`
	User      string             `json:"dbUser" bson:"dbUser"`
	Password  string             `json:"dbPassword" bson:"dbPassword"`
	APIKey    string             `json:"apiKey" bson:"apiKey"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

// ExporterSite ...
type ExporterSite struct {
	ID        int       `json:"id" db:"id"`
	DbName    string    `json:"dbName" db:"db_name"`
	Host      string    `json:"dbHost" db:"db_host"`
	User      string    `json:"dbUser" db:"db_user"`
	Password  string    `json:"dbPassword" db:"db_password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// AddSiteErrorResp ...
type AddSiteErrorResp struct {
	Message string `json:"message"`
}

// ExportReqParams ...
type ExportReqParams struct {
	// Location CSV file is Saved Locally when exported from DB
	// We Save to File here as to not consume NVRAM
	SavePath string `json:"savePath"`
	SFTPHost string `json:"sftpHost"`
	SFTPUser string `json:"sftpUser"`
	SFTPPass string `json:"sftpPass"`
	// Location DIR File will be stored on the SFTP Server
	SFTPPath string `json:"sftpPath"`
}

// ExporterServiceParams ...
type ExporterServiceParams struct {
	Sites    []RxG
	Logger   *log.Logger
	SavePath string
	SFTPHost string
	SFTPUser string
	SFTPPass string
	SFTPPath string
}

// ExporterParams ...
type ExporterParams struct {
	DbInfo
	SavePath string
	Logger   *log.Logger
}

// DbInfo ...
type DbInfo struct {
	Name string
	Host string
	User string
	Pass string
}

// AccountsDevices ...
type AccountsDevices struct {
	Accounts []rgnets.Account `json:"accounts"`
	Devices  []rgnets.Device  `json:"devices"`
}
