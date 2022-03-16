package apgplatf

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Site ...
type Site struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name          string             `bson:"name" json:"name"`
	ProjectNumber string             `bson:"projectNumber" json:"projectNumber"`
	Domain        string             `bson:"domain" json:"domain"`
	State         string             `bson:"state" json:"state"`
	BedCount      int64              `bson:"bedCount" json:"bedCount"`
	Website       string             `bson:"website" json:"website"`
	Service       struct {
		Name string `bson:"name" json:"name"`
		SnID string `bson:"snId" json:"snId"`
	} `bson:"service" json:"service"`
	ServiceNowInfo struct {
		CustomerID    string `bson:"customerId" json:"customerId"`
		AccountNumber string `bson:"accountNumber" json:"accountNumber"`
	} `bson:"serviceNowInfo" json:"serviceNowInfo"`
	Location       Location `bson:"location" json:"location"`
	ProjectActive  bool     `bson:"projectActive" json:"projectActive"`
	ProjectManager struct {
		Email string `bson:"email" json:"email"`
		SnID  string `bson:"snId" json:"snId"`
	} `bson:"projectManager" json:"projectManager"`
	InstallLead struct {
		Email string `bson:"email" json:"email"`
		SnID  string `bson:"snId" json:"snId"`
	} `bson:"installLead" json:"installLead"`
	Buildings     []Building `bson:"buildings" json:"buildings"`
	CreatedAt     time.Time  `bson:"createdAt" json:"createdAt"`
	NetProperties struct {
		Credentials   Credentials    `bson:"credentials,omitempty" json:"credentials,omitempty"`
		WlanVendor    string         `bson:"wlanVendor" json:"wlanVendor"`
		NacType       string         `bson:"nacType" json:"nacType"`
		RadiusServers []RadiusServer `bson:"radiusServers,omitempty" json:"radiusServers"`
		Networks      []Network      `bson:"networks" json:"networks"`
		Vlans         []SiteVlan     `bson:"vlans" json:"vlans"`
		SSIDs         []struct {
			Name string `bson:"name" json:"name"`
			Freq string `bson:"freq" json:"freq"`
		} `bson:"ssids" json:"ssids"`
	} `bson:"netProperties" json:"netProperties"`
	InfraDetails InfraStructure `bson:"infraDetails" json:"infraDetails"`
}

// InfraStructure
type InfraStructure struct {
	// CradlePoint Information
	CPInfo      CradlePointInfo `bson:"cpInfo" json:"cpInfo"`
	VpnRtrIP    string          `bson:"vpnRtrIp" json:"vpnRtrIp"`
	CoreSwIP    string          `bson:"coreSwIp" json:"coreSwIp"`
	CoreSwModel string          `bson:"coreSwModel" json:"coreSwModel"`
	PduInfo     PDUDetails      `bson:"pduInfo" json:"pduInfo"`
	TsConsole   []InfraPort     `bson:"tsConsole" json:"tsConsole"`
}

// CradlePointInfo
type CradlePointInfo struct {
	ID     string `bson:"id" json:"id"`
	Name   string `bson:"name" json:"name"`
	IPAddr string `bson:"ipAddr" json:"ipAddr"`
	State  string `bson:"state" json:"state"`
}

// PDUDetails
type PDUDetails struct {
	IPAddr  string      `bson:"ipAddr" json:"ipAddr"`
	Outlets []InfraPort `bson:"outlets" json:"outlets"`
}

// InfraPort
type InfraPort struct {
	PortNum        string `bson:"portNum" json:"portNum"`
	AttachedDevice string `bson:"attachedDevice" json:"attachedDevice"`
}

// Network ...
type Network struct {
	// Network Name: SwitchLanMgmt
	Name   string `bson:"name" json:"name"`
	Subnet string `bson:"subnet" json:"subnet"`
	Mask   string `bson:"mask" json:"mask"`
	Cidr   uint8  `bson:"cidr" json:"cidr"`
}

// Location ...
type Location struct {
	Lat float64 `bson:"latitude" json:"latitude"`
	Lng float64 `bson:"longitude" json:"longitude"`
}

// Building ...
type Building struct {
	Name      string   `bson:"name" json:"name"`
	ShortName string   `bson:"shortName" json:"shortName"`
	Location  Location `bson:"location" json:"location"`
}

// SiteVlan ...
type SiteVlan struct {
	ID   uint16 `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

// RadiusServer ...
type RadiusServer struct {
	Hostname string `bson:"hostname" json:"hostname"`
	IPAddr   string `bson:"ipAddress" json:"ipAddress"`
	Key      string `bson:"key" json:"key"`
}

// Credentials ...
type Credentials struct {
	User     string `bson:"user,omitempty" json:"user,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

// Device ...
type Device struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SiteID        primitive.ObjectID `bson:"siteId" json:"siteId"`
	SiteName      string             `bson:"siteName" json:"siteName"`
	Info          DeviceInfo         `bson:"deviceInfo" json:"deviceInfo"`
	Configuration Configuration      `bson:"configuration" json:"configuration"`
}

// DeviceInfo ...
type DeviceInfo struct {
	Hostname            string             `bson:"hostname" json:"hostname"`
	Serial              string             `bson:"serialNumber" json:"serialNumber"`
	ProductID           string             `bson:"pid" json:"pid"`
	Vendor              string             `bson:"vendor" json:"vendor"`
	Category            string             `bson:"category" json:"category"`
	State               string             `bson:"state" json:"state"`
	OnboardingState     string             `bson:"onbState" json:"onbState"`
	ImgVersion          string             `bson:"imageVersion" json:"imageVersion"`
	ImageFile           string             `bson:"imageFile" json:"imageFile"`
	FileSysList         []FileSystem       `bson:"fileSystemList" json:"fileSystemList"`
	DeviceSerialNumbers []string           `bson:"deviceSerialNumbers" json:"deviceSerialNumbers"`
	IsStack             bool               `bson:"stack" json:"stack"`
	ChassisMembers      []Member           `bson:"chassisMembers,omitempty" json:"chassisMembers,omitempty"`
	Created             time.Time          `bson:"createdAt" json:"createdAt"`
	LastUpdated         time.Time          `bson:"lastUpdated" json:"lastUpdated"`
	Topology            []TopologyNeighbor `bson:"topology,omitempty" json:"topology,omitempty"`
}

// Configuration represent the logical configuration of a Network Device
type Configuration struct {
	Mgmt  MgmtInfo `bson:"mgmt" json:"mgmt"`
	Image struct {
		Version  string `bson:"version" json:"version"`
		Filename string `bson:"filename" json:"filename"`
	} `bson:"image" json:"image"`
	// The Logical Number defined for a Cisco 9800
	// Values can only be 1 or 2
	ChassisNumber uint8  `bson:"chassisNumber,omitemtpy" json:"chassisNumber,omitempty"`
	RadiusServer  string `bson:"radiusServer,omitempty" json:"radiusServer,omitempty"`
	// Used for Cisco Catalyst 9800 HA Pairing
	HaInitCmds  []string `bson:"haInitCmds,omitempty" json:"haInitCmds,omitempty"`
	CliCfg      string   `bson:"cli,omitempty" json:"cli,omitempty"`
	CfgFilename string   `bson:"cfgFilename" json:"cfgFilename"`
}

// MgmtInfo ...
type MgmtInfo struct {
	IPAddr         string `bson:"ipAddress" json:"ipAddress"`
	NetworkMask    string `bson:"networkMask" json:"networkMask"`
	Interface      string `bson:"interface" json:"interface"`
	VlanID         uint8  `bson:"vlanId,omitempty" json:"vlanId"`
	VlanName       string `bson:"vlanName,omitempty" json:"vlanName"`
	DefaultGateway string `bson:"defaultGateway" json:"defaultGateway"`
	// Redundancy Management Interface IP for HA Cisco Catalyst 9800s
	Chassis1RMIAddr string `bson:"chassis1RmiAddress,omitemtpy" json:"chassis1RmiAddress,omitempty"`
	Chassis2RMIAddr string `bson:"chassis2RmiAddress,omitemtpy" json:"chassis2RmiAddress,omitempty"`
}

// DeviceState ...
type DeviceState string

// DeviceStates ...
type DeviceStates struct {
	// "Created"
	Created DeviceState
	// "Initializing"
	Initializing DeviceState
}

// TopologyNeighbor ...
type TopologyNeighbor struct {
	IntfName       string `bson:"intfName" json:"intfName"`
	ShortIntfName  string `bson:"shortIntfName" json:"shortIntfName"`
	MacAddr        string `bson:"macAddr" json:"macAddr"`
	RemoteSwitch   string `bson:"remoteSwitch" json:"remoteSwitch"`
	RemoteIntf     string `bson:"remoteIntf" json:"remoteIntf"`
	RemotePlatform string `bson:"remotePlatform" json:"remotePlatform"`
	RemoteVersion  string `bson:"remoteVersion" json:"remoteVersion"`
}

// FileSystem ...
type FileSystem struct {
	Name      string `bson:"name" json:"name"`
	Type      string `bson:"type" json:"type"`
	Readable  bool   `bson:"readable" json:"readable"`
	Writeable bool   `bson:"writeable" json:"writeable"`
	FreeSpace uint32 `bson:"freespace" json:"freespace"`
	Size      uint32 `bson:"size" json:"size"`
}

// StackInfo ...
type StackInfo struct {
	// If TRUE then all Other Properties of this Struct will be omitted
	// Only IOS XE Versions SupportWorkFlows
	SupportsWorkflows bool     `bson:"supportsStackWorkflows" json:"supportsStackWorkflows"`
	FirmwareMismatch  bool     `bson:"firmwareMismatch" json:"firmwareMismatch"`
	IsFullRing        bool     `bson:"isFullring" json:"isFullRing,omitempty"`
	MemberList        []Member `bson:"stackMemberList" json:"stackMemberList,omitempty"`
}

// Member ...
type Member struct {
	// Switch or Chassis Number
	UnitNumber      int    `bson:"unitNumber" json:"unitNumber"`
	Serial          string `bson:"serialNumber" json:"serialNumber,omitempty"`
	Role            string `bson:"role" json:"role"`
	Priority        uint8  `bson:"priority" json:"priority"`
	MacAddr         string `bson:"macAddress" json:"macAddress"`
	State           string `bson:"state" json:"state"`
	FirmwareVersion string `bson:"fwVersion" json:"fwVersion"`
	// The Fully Qualified Cisco Device Model #
	ProductID string `bson:"pid" json:"pid"`
	// The Redundacy Management Interface IP Address in the Case of Redundant
	// Cisco Catalyst 9800 WLC Chassis'
	RMIAddr string `bson:"rmiAddress,omitempty" json:"rmiAddress,omitempty"`
	// The RP Redundant Port (Self Generated) IP Address on Cisco Catalyst 9800-L|40|80
	RPAddr string `bson:"rpAddress,omitempty" json:"rpAddress,omitempty"`
}

// DeviceHistory ...
type DeviceHistory struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	DeviceID primitive.ObjectID  `bson:"deviceId" json:"deviceId"`
	Tasks    []DeviceHistoryTask `bson:"tasks" json:"tasks"`
}

// DeviceHistoryTask ...
type DeviceHistoryTask struct {
	SessionID string    `bson:"sessionId" json:"sessionId"`
	TimeStamp time.Time `bson:"timestamp" json:"timestamp"`
	Details   string    `bson:"details" json:"details"`
	Completed bool      `bson:"completed" json:"completed"`
}

// RealtimeDevice ...
type RealtimeDevice struct {
	ID         primitive.ObjectID `json:"id"`
	SiteID     primitive.ObjectID `json:"siteId"`
	Hostname   string             `json:"hostname"`
	State      string             `json:"state"`
	ActionName string             `json:"actionName"`
}
