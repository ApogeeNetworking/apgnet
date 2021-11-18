package apgsn

// CustomerAccount ...
type CustomerAccount struct {
	ID         string  `json:"id"`
	Number     string  `json:"number"`
	BedCount   int     `json:"bedCount"`
	ClientName string  `json:"clientName"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Website    string  `json:"website"`
}

// CustomerAccountResp ...
type CustomerAccountResp struct {
	ID         string `json:"sys_id"`
	Number     string `json:"number"`
	BedCount   string `json:"u_bed_count"`
	ClientName string `json:"name"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Website    string `json:"website"`
}

// CustomerName ...
type CustomerName struct {
	Name string `json:"name"`
	ID   string `json:"sys_id"`
}

// Incident ...
type Incident struct {
	ID        string `json:"sys_id"`
	WatchList string `json:"watch_list"`
	Number    string `json:"number"`
	// WorkOrder State => Indicate whether we can close the Ticket or Not
	SubState string `json:"u_inc_substate"`
	// 2 = In Progress
	// 3 = On Hold
	// 6 = Resolved
	// 7 = Closed
	State string `json:"state"`
	// 6 = Project (default)
	// 1 = COLO
	// 2 = P0
	// 3 = P1
	// 4 = P2
	// 5 = P3
	Priority string `json:"priority"`
	// 1 = Vital
	// 2 = High
	// 3 = Medium
	// 4 = Low (default)
	Urgency          string    `json:"urgency"`
	SysCreatedBy     string    `json:"sys_created_by"`
	WorkNotesList    string    `json:"work_notes_list"`
	ShortDescription string    `json:"short_description"`
	ClosedBy         RefFields `json:"closed_by"`
	AssignedTo       RefFields `json:"assigned_to"`
	// WorkOrder Info
	UWork           RefFields `json:"u_work"`
	ResolvedBy      RefFields `json:"resolved_by"`
	OpenedBy        RefFields `json:"opened_by"`
	CallerID        RefFields `json:"caller_id"`
	Subcategory     string    `json:"subcategory"`
	WorkNotes       string    `json:"work_notes"`
	AssignmentGroup RefFields `json:"assignment_group"`
	Description     string    `json:"description"`
	CloseNotes      string    `json:"close_notes"`
	IncidentState   string    `json:"incident_state"`
	Company         RefFields `json:"company"`
	// ResNet
	// Managed Campus
	BusinessService RefFields `json:"business_service"`
	Severity        string    `json:"severity"`
	Location        string    `json:"location"`
	Category        string    `json:"category"`
}

// RefFields ...
type RefFields struct {
	Link        string `json:"link"`
	RefID       string `json:"value,omitempty"`
	DisplayName string `json:"display_value,omitempty"`
}

// CreateIncidentReq ...
type CreateIncidentReq struct {
	Description       string `json:"description"`
	ShortDescription  string `json:"short_description"`
	Category          string `json:"category"`
	SubCategory       string `json:"subcategory"`
	CompanyID         string `json:"company"`
	BusinessServiceID string `json:"business_service"`
	// Email Address of Creator
	CreatedBy         string `json:"sys_created_by"`
	AssignmentGroupID string `json:"assignment_group"`
	AssignedTo        string `json:"assigned_to"`
	Caller            string `json:"caller_id,omitempty"`
	// Default: 6
	Priority string `json:"priority"`
	// Default: 5
	Impact string `json:"impact"`
	// Default: 3
	Severity  string `json:"severity"`
	WatchList string `json:"watch_list,omitempty"`
}

// WorkOrder ...
type WorkOrder struct {
	ID            string `json:"sys_id"`
	Number        string `json:"number"`
	State         string `json:"state"`
	WorkNotesList string `json:"work_notes_list"`
}

// CreateWorkOrderReq ...
type CreateWorkOrderReq struct {
	Description       string `json:"description"`
	ShortDescription  string `json:"short_description"`
	Category          string `json:"category"`
	SubCategory       string `json:"subcategory"`
	AssignmentGroupID string `json:"assignment_group"`
	IncidentID        string `json:"initiated_from"`
	CompanyID         string `json:"account"`
	BusinessService   string `json:"business_service"`
	CallBackNumber    string `json:"u_best_call_back_number"`
	ContactType       string `json:"u_contact_type"`
	// Default: 6
	Priority string `json:"priority"`
	// Default: 5
	Impact string `json:"impact"`
	// Default: 3
	Severity string `json:"severity"`
}

// User ...
type User struct {
	ID           string    `json:"sys_id"`
	State        string    `json:"state"`
	ZipCode      string    `json:"zip"`
	Phone        string    `json:"phone"`
	Name         string    `json:"name"`
	City         string    `json:"city"`
	Username     string    `json:"user_name"`
	Title        string    `json:"title"`
	DefaultGroup RefFields `json:"u_default_group"`
	Department   RefFields `json:"department"`
	FirstName    string    `json:"first_name"`
	Email        string    `json:"email"`
	Manager      RefFields `json:"manager"`
	LastName     string    `json:"last_name"`
}

// https://apogeedev.service-now.com/customer_account_list.do?sysparm_query=customer%3Dtrue%5Eaccount_parentISEMPTY&sysparm_view=case
