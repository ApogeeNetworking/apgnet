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
	WatchList string `json:"watch_list,omitempty"`
	Number    string `json:"number"`
	// WorkOrder State => Indicate whether we can close the Ticket or Not
	SubState string `json:"u_inc_substate,omitempty"`
	// 2 = In Progress
	// 3 = On Hold
	// 6 = Resolved
	// 7 = Closed
	State string `json:"state,omitempty"`
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
	WorkNotesList    string    `json:"work_notes_list,omitempty"`
	ShortDescription string    `json:"short_description"`
	ClosedBy         RefFields `json:"closed_by,omitempty"`
	AssignedTo       RefFields `json:"assigned_to,omitempty"`
	// WorkOrder Info
	UWork           RefFields `json:"u_work"`
	ResolvedBy      RefFields `json:"resolved_by,omitempty"`
	OpenedBy        RefFields `json:"opened_by,omitempty"`
	CallerID        RefFields `json:"caller_id"`
	Subcategory     string    `json:"subcategory,omitempty"`
	WorkNotes       string    `json:"work_notes,omitempty"`
	AssignmentGroup RefFields `json:"assignment_group"`
	Description     string    `json:"description"`
	CloseNotes      string    `json:"close_notes,omitempty"`
	IncidentState   string    `json:"incident_state,omitempty"`
	Company         RefFields `json:"company"`
	// ResNet
	// Managed Campus
	BusinessService RefFields `json:"business_service"`
	Severity        string    `json:"severity"`
	Location        string    `json:"location,omitempty"`
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
	AssignedTo        string `json:"assigned_to,omitempty"`
	Caller            string `json:"caller_id"`
	// Default: 6
	Priority string `json:"priority"`
	// Default: 5
	Impact string `json:"impact"`
	// Default: 3
	Severity  string `json:"severity"`
	WatchList string `json:"watch_list"`
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

// UserGroup ...
type UserGroup struct {
	ID          string    `json:"sys_id"`
	Name        string    `json:"name"`
	Manager     RefFields `json:"manager"`
	MemberCount string    `json:"sys_mod_count"`
	Active      string    `json:"active"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	HourlyRate  string    `json:"hourly_rate"`
	UpdatedOn   string    `json:"sys_updated_on"`
	CreatedOn   string    `json:"sys_created_on"`
	Email       string    `json:"email"`
}

// GetIncidentsParams ...
type GetIncidentsParams struct {
	CustomerID        string
	Email             string
	AssignmentGroupID string
}

// Circuit ...
type Circuit struct {
	IspEnd                string    `json:"u_isp_end"`
	Type                  string    `json:"u_type_circuit"`
	CirFoc                string    `json:"u_cir_foc"`
	BsrApogeePort         string    `json:"u_bsr_apogee_port"`
	CirSdi                string    `json:"u_cir_sdi"`
	AcctNo                string    `json:"u_acct_no"`
	CirProj               string    `json:"u_cir_proj"`
	SysUpdatedOn          string    `json:"sys_updated_on"`
	AccountClient         RefFields `json:"u_account_client"`
	CirComplete           string    `json:"u_cir_complete"`
	CircuitNotes          string    `json:"u_circuit_notes"`
	SysUpdatedBy          string    `json:"sys_updated_by"`
	CirNni                string    `json:"u_cir_nni"`
	SysCreatedOn          string    `json:"sys_created_on"`
	CirSigdate            string    `json:"u_cir_sigdate"`
	CirCust               string    `json:"u_cir_cust"`
	CirIspdate            string    `json:"u_cir_ispdate"`
	SysCreatedBy          string    `json:"sys_created_by"`
	Vendor                RefFields `json:"u_cir_dc_vendor"`
	CirXconnectPatchpan   string    `json:"u_cir_xconnect_patchpan"`
	MsaRenew              string    `json:"u_msa_renew"`
	Isp                   RefFields `json:"u_isp"`
	CirIspfoc             string    `json:"u_cir_ispfoc"`
	CirProjtype           string    `json:"u_cir_projtype"`
	State                 string    `json:"u_state"`
	CircuitCust           string    `json:"u_circuit_cust"`
	CirXconnectInstall    string    `json:"u_cir_xconnect_install"`
	Demarc1               string    `json:"u_demarc1"`
	BsrProjStat           string    `json:"u_bsr_proj_stat"`
	CirYear               string    `json:"u_cir_year"`
	Demarc2               string    `json:"u_demarc2"`
	Email1                string    `json:"u_email_1"`
	ApgContractSig        string    `json:"u_apg_contract_sig"`
	BillDate              string    `json:"u_bill_date"`
	CirXconnectApogeeport string    `json:"u_cir_xconnect_apogeeport"`
	CirXconnectPort       string    `json:"u_cir_xconnect_port"`
	CircuitType           string    `json:"u_circuit_type"`
	ColoDep               string    `json:"u_colo_dep"`
	Portal                string    `json:"u_portal"`
	CirApogeecab          string    `json:"u_cir_apogeecab"`
	CircuitSetInactive    string    `json:"u_circuit_set_inactive"`
	CirIsppm              string    `json:"u_cir_isppm"`
	CirBackuppe           string    `json:"u_cir_backuppe"`
	CirStatus             string    `json:"u_cir_status"`
	CirXconnectContype    string    `json:"u_cir_xconnect_contype"`
	CirNewcustname        string    `json:"u_cir_newcustname"`
	DirTurnup             string    `json:"u_dir_turnup"`
	Basicspeed            string    `json:"u_basicspeed"`
	CirGreenlight         string    `json:"u_cir_greenlight"`
	BurstAvail            string    `json:"u_burst_avail"`
	IspOrderNum           string    `json:"u_isp_order_num"`
	LecCid                string    `json:"u_lec_cid"`
	ExecDate              string    `json:"u_exec_date"`
	SysID                 string    `json:"sys_id"`
	CirPrimepe            string    `json:"u_cir_primepe"`
	GlideDuration1        string    `json:"u_glide_duration_1"`
	ContractEnd           string    `json:"u_contract_end"`
	IspHandoffA           string    `json:"u_isp_handoff_a"`
	CircuitColo           RefFields `json:"u_circuit_colo"`
	CircuitActive         string    `json:"u_circuit_active"`
	Transport             string    `json:"u_transport"`
	VenContact            string    `json:"u_ven_contact"`
	CirDcTerm             string    `json:"u_cir_dc_term"`
	CirNewcost            string    `json:"u_cir_newcost"`
	CircuitCurrent        string    `json:"u_circuit_current"`
	Lec                   string    `json:"u_lec"`
	MtmLetter             string    `json:"u_mtm_letter"`
	SigDate               string    `json:"u_sig_date"`
	CirXconnectExp        string    `json:"u_cir_xconnect_exp"`
	Number                string    `json:"u_number"`
	Term                  string    `json:"u_term"`
	BsrIspDelDate         string    `json:"u_bsr_isp_del_date"`
	CirProjnotes          string    `json:"u_cir_projnotes"`
	CirXconnectMedia      string    `json:"u_cir_xconnect_media"`
	Journal1              string    `json:"u_journal_1"`
	CirNslink             string    `json:"u_cir_nslink"`
	CircuitSupport        string    `json:"u_circuit_support"`
	CirApgPm              string    `json:"u_cir_apogeepm"`
	SysModCount           string    `json:"sys_mod_count"`
	CirPacksize           string    `json:"u_cir_packsize"`
	Cid                   string    `json:"u_cid"`
	SysTags               string    `json:"sys_tags"`
	ProvicerCid           string    `json:"u_provicercid"`
	Sharepoint            string    `json:"u_sharepoint"`
	CirNsprojnum          string    `json:"u_cir_nsprojnum"`
	CurBedcount           string    `json:"u_cur_bedcount"`
	BsrHandoff            string    `json:"u_bsr_handoff"`
	CirBsr                string    `json:"u_cir_bsr"`
	Bandwidth             string    `json:"u_bandwidth"`
	Contacts              string    `json:"u_contacts"`
	CustomerCircuitID     string    `json:"u_customer_circuit_id"`
}

// https://apogeedev.service-now.com/customer_account_list.do?sysparm_query=customer%3Dtrue%5Eaccount_parentISEMPTY&sysparm_view=case
