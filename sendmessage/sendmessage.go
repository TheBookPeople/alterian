package sendmessage

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type DMOperator string

const (
	DMOperatorDMOPEQUAL DMOperator = "DMOPEQUAL"

	DMOperatorDMOPNOTEQUAL DMOperator = "DMOPNOTEQUAL"

	DMOperatorDMOPLESSTHAN DMOperator = "DMOPLESSTHAN"

	DMOperatorDMOPGREATERTHAN DMOperator = "DMOPGREATERTHAN"

	DMOperatorDMOPBEGINSWITH DMOperator = "DMOPBEGINSWITH"

	DMOperatorDMOPNOTBEGINSWITH DMOperator = "DMOPNOTBEGINSWITH"

	DMOperatorDMOPENDSWITH DMOperator = "DMOPENDSWITH"

	DMOperatorDMOPNOTENDSWITH DMOperator = "DMOPNOTENDSWITH"

	DMOperatorDMOPCONTAINS DMOperator = "DMOPCONTAINS"

	DMOperatorDMOPNOTCONTAINS DMOperator = "DMOPNOTCONTAINS"
)

type DMCombine string

const (
	DMCombineDMCMNONE DMCombine = "DMCMNONE"

	DMCombineDMCMAND DMCombine = "DMCMAND"

	DMCombineDMCMOR DMCombine = "DMCMOR"
)

type DMRecurringScheduleType string

const (
	DMRecurringScheduleTypeDMRSTDAILY DMRecurringScheduleType = "DMRSTDAILY"

	DMRecurringScheduleTypeDMRSTWEEKLY DMRecurringScheduleType = "DMRSTWEEKLY"

	DMRecurringScheduleTypeDMRSTMONTHLY DMRecurringScheduleType = "DMRSTMONTHLY"
)

type DMDeploymentStatus string

const (
	DMDeploymentStatusDMDSSETUP DMDeploymentStatus = "DMDSSETUP"

	DMDeploymentStatusDMDSQUEUED DMDeploymentStatus = "DMDSQUEUED"

	DMDeploymentStatusDMDSSCHEDULED DMDeploymentStatus = "DMDSSCHEDULED"

	DMDeploymentStatusDMDSPREFETCHING DMDeploymentStatus = "DMDSPREFETCHING"

	DMDeploymentStatusDMDSPREFETCHED DMDeploymentStatus = "DMDSPREFETCHED"

	DMDeploymentStatusDMDSSENDING DMDeploymentStatus = "DMDSSENDING"

	DMDeploymentStatusDMDSPAUSED DMDeploymentStatus = "DMDSPAUSED"

	DMDeploymentStatusDMDSABORTED DMDeploymentStatus = "DMDSABORTED"

	DMDeploymentStatusDMDSCOMPLETED DMDeploymentStatus = "DMDSCOMPLETED"

	DMDeploymentStatusDMDSPREFECTHINGINPROGRESS DMDeploymentStatus = "DMDSPREFECTHINGINPROGRESS"

	DMDeploymentStatusDMDSFAILED DMDeploymentStatus = "DMDSFAILED"
)

type DMListType string

const (
	DMListTypeDMLTRECIPIENT DMListType = "DMLTRECIPIENT"

	DMListTypeDMLTSUPPRESS DMListType = "DMLTSUPPRESS"

	DMListTypeDMLTDOMAINSUPPRESSION DMListType = "DMLTDOMAINSUPPRESSION"

	DMListTypeDMLTMD5SUPPRESSION DMListType = "DMLTMD5SUPPRESSION"
)

type DMRDateRangeType string

const (
	DMRDateRangeTypeDMDTBEFORE DMRDateRangeType = "DMDTBEFORE"

	DMRDateRangeTypeDMDTAFTER DMRDateRangeType = "DMDTAFTER"

	DMRDateRangeTypeDMDTBETWEEN DMRDateRangeType = "DMDTBETWEEN"

	DMRDateRangeTypeDMDTYEARTODATE DMRDateRangeType = "DMDTYEARTODATE"

	DMRDateRangeTypeDMDTMONTHTODATE DMRDateRangeType = "DMDTMONTHTODATE"

	DMRDateRangeTypeDMDTWEEKTODATE DMRDateRangeType = "DMDTWEEKTODATE"

	DMRDateRangeTypeDMDTTODAY DMRDateRangeType = "DMDTTODAY"
)

type DMFieldQueryType string

const (
	DMFieldQueryTypeDMQTPREFERTEXT DMFieldQueryType = "DMQTPREFERTEXT"

	DMFieldQueryTypeDMQTVERIFIEDFLASH DMFieldQueryType = "DMQTVERIFIEDFLASH"

	DMFieldQueryTypeDMQTVERIFIEDHTML DMFieldQueryType = "DMQTVERIFIEDHTML"

	DMFieldQueryTypeDMQTENABLED DMFieldQueryType = "DMQTENABLED"

	DMFieldQueryTypeDMQTLASTEVENT DMFieldQueryType = "DMQTLASTEVENT"

	DMFieldQueryTypeDMQTCREATED DMFieldQueryType = "DMQTCREATED"

	DMFieldQueryTypeDMQTMODIFIED DMFieldQueryType = "DMQTMODIFIED"

	DMFieldQueryTypeDMQTDATEFIELD DMFieldQueryType = "DMQTDATEFIELD"

	DMFieldQueryTypeDMQTNUMERICFIELD DMFieldQueryType = "DMQTNUMERICFIELD"

	DMFieldQueryTypeDMQTDECIMALFIELD DMFieldQueryType = "DMQTDECIMALFIELD"

	DMFieldQueryTypeDMQTBOOLEANFIELD DMFieldQueryType = "DMQTBOOLEANFIELD"

	DMFieldQueryTypeDMQTSTRINGFIELD DMFieldQueryType = "DMQTSTRINGFIELD"

	DMFieldQueryTypeDMQTRSSONLY DMFieldQueryType = "DMQTRSSONLY"

	DMFieldQueryTypeDMQTLASTENGAGED DMFieldQueryType = "DMQTLASTENGAGED"
)

type DMSQLOperator string

const (
	DMSQLOperatorDMSQEQUAL DMSQLOperator = "DMSQEQUAL"

	DMSQLOperatorDMSQNOTEQUAL DMSQLOperator = "DMSQNOTEQUAL"

	DMSQLOperatorDMSQLESSTHAN DMSQLOperator = "DMSQLESSTHAN"

	DMSQLOperatorDMSQGREATERTHAN DMSQLOperator = "DMSQGREATERTHAN"

	DMSQLOperatorDMSQBEGINSWITH DMSQLOperator = "DMSQBEGINSWITH"

	DMSQLOperatorDMSQNOTBEGINSWITH DMSQLOperator = "DMSQNOTBEGINSWITH"

	DMSQLOperatorDMSQENDSWITH DMSQLOperator = "DMSQENDSWITH"

	DMSQLOperatorDMSQNOTENDSWITH DMSQLOperator = "DMSQNOTENDSWITH"

	DMSQLOperatorDMSQCONTAINS DMSQLOperator = "DMSQCONTAINS"

	DMSQLOperatorDMSQNOTCONTAINS DMSQLOperator = "DMSQNOTCONTAINS"

	DMSQLOperatorDMSQIN DMSQLOperator = "DMSQIN"

	DMSQLOperatorDMSQNOTIN DMSQLOperator = "DMSQNOTIN"

	DMSQLOperatorDMSQBETWEEN DMSQLOperator = "DMSQBETWEEN"

	DMSQLOperatorDMSQNOTBETWEEN DMSQLOperator = "DMSQNOTBETWEEN"
)

type DMFieldSubType string

const (
	DMFieldSubTypeDMFSTNA DMFieldSubType = "DMFSTNA"

	DMFieldSubTypeDMFSTEMAILUSER DMFieldSubType = "DMFSTEMAILUSER"

	DMFieldSubTypeDMFSTEMAILDOMAIN DMFieldSubType = "DMFSTEMAILDOMAIN"

	DMFieldSubTypeDMFSTEMAILHEALTH DMFieldSubType = "DMFSTEMAILHEALTH"

	DMFieldSubTypeDMFSTDATEYEAR DMFieldSubType = "DMFSTDATEYEAR"

	DMFieldSubTypeDMFSTDATEMONTH DMFieldSubType = "DMFSTDATEMONTH"

	DMFieldSubTypeDMFSTDATEDAY DMFieldSubType = "DMFSTDATEDAY"

	DMFieldSubTypeDMFSTDATETIME DMFieldSubType = "DMFSTDATETIME"

	DMFieldSubTypeDMFSTDATEDAYSFROMNOW DMFieldSubType = "DMFSTDATEDAYSFROMNOW"
)

type DMDeploymentField string

const (
	DMDeploymentFieldDMDFCREATIVECATEGORYNAME DMDeploymentField = "DMDFCREATIVECATEGORYNAME"

	DMDeploymentFieldDMDFCREATIVENAME DMDeploymentField = "DMDFCREATIVENAME"

	DMDeploymentFieldDMDFNAME DMDeploymentField = "DMDFNAME"

	DMDeploymentFieldDMDFLISTNAMES DMDeploymentField = "DMDFLISTNAMES"

	DMDeploymentFieldDMDFUSERNAME DMDeploymentField = "DMDFUSERNAME"

	DMDeploymentFieldDMDFFROMEMAIL DMDeploymentField = "DMDFFROMEMAIL"

	DMDeploymentFieldDMDFFROMALIAS DMDeploymentField = "DMDFFROMALIAS"

	DMDeploymentFieldDMDFTOEMAIL DMDeploymentField = "DMDFTOEMAIL"

	DMDeploymentFieldDMDFTOALIAS DMDeploymentField = "DMDFTOALIAS"

	DMDeploymentFieldDMDFREPLYTO DMDeploymentField = "DMDFREPLYTO"

	DMDeploymentFieldDMDFSUBJECT DMDeploymentField = "DMDFSUBJECT"

	DMDeploymentFieldDMDFCREATED DMDeploymentField = "DMDFCREATED"

	DMDeploymentFieldDMDFMODIFIED DMDeploymentField = "DMDFMODIFIED"

	DMDeploymentFieldDMDFSTARTED DMDeploymentField = "DMDFSTARTED"

	DMDeploymentFieldDMDFFINISHED DMDeploymentField = "DMDFFINISHED"

	DMDeploymentFieldDMDFSCHEDULE DMDeploymentField = "DMDFSCHEDULE"

	DMDeploymentFieldDMDFTOTAL DMDeploymentField = "DMDFTOTAL"

	DMDeploymentFieldDMDFSENT DMDeploymentField = "DMDFSENT"

	DMDeploymentFieldDMDFQUEUED DMDeploymentField = "DMDFQUEUED"

	DMDeploymentFieldDMDFPERCENTDONE DMDeploymentField = "DMDFPERCENTDONE"

	DMDeploymentFieldDMDFTHROTTLE DMDeploymentField = "DMDFTHROTTLE"

	DMDeploymentFieldDMDFSTATUS DMDeploymentField = "DMDFSTATUS"

	DMDeploymentFieldDMDFID DMDeploymentField = "DMDFID"

	DMDeploymentFieldDMDFLISTTOTAL DMDeploymentField = "DMDFLISTTOTAL"

	DMDeploymentFieldDMDFRECIPIENTSUPPRESSED DMDeploymentField = "DMDFRECIPIENTSUPPRESSED"

	DMDeploymentFieldDMDFLISTSUPPRESSED DMDeploymentField = "DMDFLISTSUPPRESSED"

	DMDeploymentFieldDMDFHEALTHSUPPRESSED DMDeploymentField = "DMDFHEALTHSUPPRESSED"

	DMDeploymentFieldDMDFEVENTSUPPRESSED DMDeploymentField = "DMDFEVENTSUPPRESSED"

	DMDeploymentFieldDMDFFIELDSUPPRESSED DMDeploymentField = "DMDFFIELDSUPPRESSED"

	DMDeploymentFieldDMDFMAILERSUBMISSIONERROR DMDeploymentField = "DMDFMAILERSUBMISSIONERROR"

	DMDeploymentFieldDMDFDCINFO DMDeploymentField = "DMDFDCINFO"

	DMDeploymentFieldDMDFRECURRENCESCHEDULE DMDeploymentField = "DMDFRECURRENCESCHEDULE"

	DMDeploymentFieldDMDFDUPESUPPRESSED DMDeploymentField = "DMDFDUPESUPPRESSED"
)

type DMCreativePermission string

const (
	DMCreativePermissionDMCPMODIFY DMCreativePermission = "DMCPMODIFY"

	DMCreativePermissionDMCPDEPLOY DMCreativePermission = "DMCPDEPLOY"

	DMCreativePermissionDMCPREPORT DMCreativePermission = "DMCPREPORT"
)

type DMVariableType string

const (
	DMVariableTypeDMVTFROMEMAIL DMVariableType = "DMVTFROMEMAIL"

	DMVariableTypeDMVTFROMALIAS DMVariableType = "DMVTFROMALIAS"

	DMVariableTypeDMVTTOEMAIL DMVariableType = "DMVTTOEMAIL"

	DMVariableTypeDMVTTOALIAS DMVariableType = "DMVTTOALIAS"

	DMVariableTypeDMVTREPLYTO DMVariableType = "DMVTREPLYTO"

	DMVariableTypeDMVTSUBJECT DMVariableType = "DMVTSUBJECT"

	DMVariableTypeDMVTSTRING DMVariableType = "DMVTSTRING"

	DMVariableTypeDMVTEMAIL DMVariableType = "DMVTEMAIL"

	DMVariableTypeDMVTTEXT DMVariableType = "DMVTTEXT"

	DMVariableTypeDMVTHTML DMVariableType = "DMVTHTML"

	DMVariableTypeDMVTDATE DMVariableType = "DMVTDATE"

	DMVariableTypeDMVTNUMERIC DMVariableType = "DMVTNUMERIC"
)

type DMEditorType string

const (
	DMEditorTypeDMETSTRING DMEditorType = "DMETSTRING"

	DMEditorTypeDMETTEXT DMEditorType = "DMETTEXT"

	DMEditorTypeDMETHTML DMEditorType = "DMETHTML"

	DMEditorTypeDMETDATE DMEditorType = "DMETDATE"

	DMEditorTypeDMETCOMBO DMEditorType = "DMETCOMBO"
)

type VolumeWarning struct {
	XMLName xml.Name `xml:"DMWebServices VolumeWarning"`
}

type VolumeWarningResponse struct {
	XMLName xml.Name `xml:"DMWebServices VolumeWarningResponse"`

	VolumeWarningResult bool `xml:"VolumeWarningResult,omitempty"`
}

type CreateDeployment struct {
	XMLName xml.Name `xml:"DMWebServices CreateDeployment"`

	Token string `xml:"Token,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type CreateDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateDeploymentResponse"`

	CreateDeploymentResult int32 `xml:"CreateDeploymentResult,omitempty"`
}

type UpdateDeployment struct {
	XMLName xml.Name `xml:"DMWebServices UpdateDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type UpdateDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateDeploymentResponse"`
}

type SetDeploymentVariables struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentVariables"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	TemplateValue *DMTemplateValue `xml:"TemplateValue,omitempty"`

	RecipientLists *ArrayOfInt `xml:"RecipientLists,omitempty"`

	Attachments *ArrayOfInt `xml:"Attachments,omitempty"`

	VariableMapping *ArrayOfDMVariableMap `xml:"VariableMapping,omitempty"`

	ProofDeployment bool `xml:"ProofDeployment,omitempty"`
}

type SetDeploymentVariablesResponse struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentVariablesResponse"`
}

type SetDeploymentDeliveryContext struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentDeliveryContext"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	DeliveryContextID int32 `xml:"DeliveryContextID,omitempty"`
}

type SetDeploymentDeliveryContextResponse struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentDeliveryContextResponse"`
}

type SetDeploymentSmartlistRefresh struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentSmartlistRefresh"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	RefreshSmartlists bool `xml:"RefreshSmartlists,omitempty"`
}

type SetDeploymentSmartlistRefreshResponse struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentSmartlistRefreshResponse"`
}

type SetDeploymentSuppression struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentSuppression"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	SuppressLists *ArrayOfInt `xml:"SuppressLists,omitempty"`

	SuppressEvents *ArrayOfInt `xml:"SuppressEvents,omitempty"`

	HealthThreshold int32 `xml:"HealthThreshold,omitempty"`
}

type SetDeploymentSuppressionResponse struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentSuppressionResponse"`
}

type GetDeploymentNotify struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentNotify"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type GetDeploymentNotifyResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentNotifyResponse"`

	GetDeploymentNotifyResult *ArrayOfDMEventNotify `xml:"GetDeploymentNotifyResult,omitempty"`

	NotifyComplete bool `xml:"NotifyComplete,omitempty"`

	NotifyEmail string `xml:"NotifyEmail,omitempty"`
}

type SetDeploymentNotify struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentNotify"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	NotifyComplete bool `xml:"NotifyComplete,omitempty"`

	NotifyEvents *ArrayOfInt `xml:"NotifyEvents,omitempty"`

	NotifyEmail string `xml:"NotifyEmail,omitempty"`
}

type SetDeploymentNotifyResponse struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentNotifyResponse"`
}

type SetDeploymentThrottle struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentThrottle"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	Choke int32 `xml:"Choke,omitempty"`

	ChokeInterval int32 `xml:"ChokeInterval,omitempty"`
}

type SetDeploymentThrottleResponse struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentThrottleResponse"`
}

type SetDeploymentSchedule struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentSchedule"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	Schedule time.Time `xml:"Schedule,omitempty"`
}

type SetDeploymentScheduleResponse struct {
	XMLName xml.Name `xml:"DMWebServices SetDeploymentScheduleResponse"`
}

type FinalizeDeployment struct {
	XMLName xml.Name `xml:"DMWebServices FinalizeDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	SendNow bool `xml:"SendNow,omitempty"`
}

type FinalizeDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices FinalizeDeploymentResponse"`
}

type FinalizeRecurringDeployment struct {
	XMLName xml.Name `xml:"DMWebServices FinalizeRecurringDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	Enddate time.Time `xml:"Enddate,omitempty"`

	Frequency *DMRecurringScheduleType `xml:"Frequency,omitempty"`
}

type FinalizeRecurringDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices FinalizeRecurringDeploymentResponse"`
}

type UpdateRecurringDeployment struct {
	XMLName xml.Name `xml:"DMWebServices UpdateRecurringDeployment"`

	Token string `xml:"Token,omitempty"`

	Deployment *DMDeployment `xml:"Deployment,omitempty"`
}

type UpdateRecurringDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateRecurringDeploymentResponse"`
}

type GetDeploymentLists struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentLists"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type GetDeploymentListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentListsResponse"`

	GetDeploymentListsResult *ArrayOfDMList `xml:"GetDeploymentListsResult,omitempty"`
}

type GetDeploymentsLists struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentsLists"`

	Token string `xml:"Token,omitempty"`

	DeploymentIDs *ArrayOfInt `xml:"DeploymentIDs,omitempty"`
}

type GetDeploymentsListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentsListsResponse"`

	GetDeploymentsListsResult *ArrayOfDMDeploymentLists `xml:"GetDeploymentsListsResult,omitempty"`
}

type GetDeploymentsEvents struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentsEvents"`

	Token string `xml:"Token,omitempty"`

	DeploymentIDs *ArrayOfInt `xml:"DeploymentIDs,omitempty"`
}

type GetDeploymentsEventsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentsEventsResponse"`

	GetDeploymentsEventsResult *ArrayOfDMEventInfo `xml:"GetDeploymentsEventsResult,omitempty"`

	DeploymentEvents *ArrayOfDMDeploymentEvents `xml:"DeploymentEvents,omitempty"`
}

type GetDeployment struct {
	XMLName xml.Name `xml:"DMWebServices GetDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type GetDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentResponse"`

	GetDeploymentResult *DMDeployment `xml:"GetDeploymentResult,omitempty"`
}

type AddRecipientToDeployment struct {
	XMLName xml.Name `xml:"DMWebServices AddRecipientToDeployment"`

	Token string `xml:"Token,omitempty"`

	RecipientID int32 `xml:"RecipientID,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type AddRecipientToDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices AddRecipientToDeploymentResponse"`
}

type GetDeployments struct {
	XMLName xml.Name `xml:"DMWebServices GetDeployments"`

	Token string `xml:"Token,omitempty"`

	DateRange *DMRDateRange `xml:"DateRange,omitempty"`

	Scheduled bool `xml:"Scheduled,omitempty"`

	Active bool `xml:"Active,omitempty"`

	Completed bool `xml:"Completed,omitempty"`

	ExcludeOneOffs bool `xml:"ExcludeOneOffs,omitempty"`

	Fields *ArrayOfDMDeploymentField `xml:"Fields,omitempty"`

	SearchByScheduledDate bool `xml:"SearchByScheduledDate,omitempty"`
}

type GetDeploymentsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetDeploymentsResponse"`

	GetDeploymentsResult *ArrayOfDMDeployment `xml:"GetDeploymentsResult,omitempty"`
}

type PauseDeployment struct {
	XMLName xml.Name `xml:"DMWebServices PauseDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type PauseDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices PauseDeploymentResponse"`

	PauseDeploymentResult *DMDeploymentStatus `xml:"PauseDeploymentResult,omitempty"`
}

type PauseRecurringDeployment struct {
	XMLName xml.Name `xml:"DMWebServices PauseRecurringDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type PauseRecurringDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices PauseRecurringDeploymentResponse"`
}

type ResumeDeployment struct {
	XMLName xml.Name `xml:"DMWebServices ResumeDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type ResumeDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices ResumeDeploymentResponse"`

	ResumeDeploymentResult *DMDeploymentStatus `xml:"ResumeDeploymentResult,omitempty"`
}

type ResumeRecurringDeployment struct {
	XMLName xml.Name `xml:"DMWebServices ResumeRecurringDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type ResumeRecurringDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices ResumeRecurringDeploymentResponse"`
}

type AbortDeployment struct {
	XMLName xml.Name `xml:"DMWebServices AbortDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type AbortDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices AbortDeploymentResponse"`

	AbortDeploymentResult *DMDeploymentStatus `xml:"AbortDeploymentResult,omitempty"`
}

type AbortRecurringDeployment struct {
	XMLName xml.Name `xml:"DMWebServices AbortRecurringDeployment"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type AbortRecurringDeploymentResponse struct {
	XMLName xml.Name `xml:"DMWebServices AbortRecurringDeploymentResponse"`
}

type SAMGetTotalRecipients struct {
	XMLName xml.Name `xml:"DMWebServices SAMGetTotalRecipients"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	RecipientLists *ArrayOfInt `xml:"RecipientLists,omitempty"`

	SuppressionLists *ArrayOfInt `xml:"SuppressionLists,omitempty"`

	SuppressEvents *ArrayOfInt `xml:"SuppressEvents,omitempty"`
}

type SAMGetTotalRecipientsResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMGetTotalRecipientsResponse"`

	SAMGetTotalRecipientsResult int32 `xml:"SAMGetTotalRecipientsResult,omitempty"`
}

type PreviewDeploymentTotal struct {
	XMLName xml.Name `xml:"DMWebServices PreviewDeploymentTotal"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	RecipientLists *ArrayOfInt `xml:"RecipientLists,omitempty"`

	SuppressionLists *ArrayOfInt `xml:"SuppressionLists,omitempty"`

	SuppressEvents *ArrayOfInt `xml:"SuppressEvents,omitempty"`
}

type PreviewDeploymentTotalResponse struct {
	XMLName xml.Name `xml:"DMWebServices PreviewDeploymentTotalResponse"`

	PreviewDeploymentTotalResult *ArrayOfHealthCount `xml:"PreviewDeploymentTotalResult,omitempty"`
}

type SAMPreviewFirst struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewFirst"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	RecipientLists *ArrayOfInt `xml:"RecipientLists,omitempty"`

	SuppressionLists *ArrayOfInt `xml:"SuppressionLists,omitempty"`

	SuppressEvents *ArrayOfInt `xml:"SuppressEvents,omitempty"`
}

type SAMPreviewFirstResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewFirstResponse"`

	SAMPreviewFirstResult string `xml:"SAMPreviewFirstResult,omitempty"`
}

type SAMPreviewLast struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewLast"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	RecipientLists *ArrayOfInt `xml:"RecipientLists,omitempty"`

	SuppressionLists *ArrayOfInt `xml:"SuppressionLists,omitempty"`

	SuppressEvents *ArrayOfInt `xml:"SuppressEvents,omitempty"`
}

type SAMPreviewLastResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewLastResponse"`

	SAMPreviewLastResult string `xml:"SAMPreviewLastResult,omitempty"`
}

type SAMPreviewNext struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewNext"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	MessageID string `xml:"MessageID,omitempty"`

	RecipientLists *ArrayOfInt `xml:"RecipientLists,omitempty"`

	SuppressionLists *ArrayOfInt `xml:"SuppressionLists,omitempty"`

	SuppressEvents *ArrayOfInt `xml:"SuppressEvents,omitempty"`
}

type SAMPreviewNextResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewNextResponse"`

	SAMPreviewNextResult string `xml:"SAMPreviewNextResult,omitempty"`
}

type SAMPreviewPrevious struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewPrevious"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	MessageID string `xml:"MessageID,omitempty"`

	RecipientLists *ArrayOfInt `xml:"RecipientLists,omitempty"`

	SuppressionLists *ArrayOfInt `xml:"SuppressionLists,omitempty"`

	SuppressEvents *ArrayOfInt `xml:"SuppressEvents,omitempty"`
}

type SAMPreviewPreviousResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewPreviousResponse"`

	SAMPreviewPreviousResult string `xml:"SAMPreviewPreviousResult,omitempty"`
}

type SAMPreviewEx struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewEx"`

	Token string `xml:"Token,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	RecordNumber int32 `xml:"RecordNumber,omitempty"`
}

type SAMPreviewExResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMPreviewExResponse"`

	SAMPreviewExResult string `xml:"SAMPreviewExResult,omitempty"`
}

type GetSAReport struct {
	XMLName xml.Name `xml:"DMWebServices GetSAReport"`

	Token string `xml:"Token,omitempty"`

	MessageID string `xml:"MessageID,omitempty"`
}

type GetSAReportResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetSAReportResponse"`

	GetSAReportResult *SAReport `xml:"GetSAReportResult,omitempty"`
}

type GetCreativeCategories struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativeCategories"`

	Token string `xml:"Token,omitempty"`
}

type GetCreativeCategoriesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativeCategoriesResponse"`

	GetCreativeCategoriesResult *ArrayOfDMCategory `xml:"GetCreativeCategoriesResult,omitempty"`
}

type GetCreatives struct {
	XMLName xml.Name `xml:"DMWebServices GetCreatives"`

	Token string `xml:"Token,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	CreativeIDList *ArrayOfInt `xml:"CreativeIDList,omitempty"`
}

type GetCreativesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativesResponse"`

	GetCreativesResult *ArrayOfDMCreative `xml:"GetCreativesResult,omitempty"`
}

type GetCreativeCategoriesEx struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativeCategoriesEx"`

	Token string `xml:"Token,omitempty"`

	CategoryIDs *ArrayOfInt `xml:"CategoryIDs,omitempty"`

	PermissionIDs *ArrayOfDMCreativePermission `xml:"PermissionIDs,omitempty"`

	GetCreatives bool `xml:"GetCreatives,omitempty"`
}

type GetCreativeCategoriesExResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativeCategoriesExResponse"`

	GetCreativeCategoriesExResult *ArrayOfDMCreativeCategory `xml:"GetCreativeCategoriesExResult,omitempty"`
}

type GetAllCreatives struct {
	XMLName xml.Name `xml:"DMWebServices GetAllCreatives"`

	Token string `xml:"Token,omitempty"`
}

type GetAllCreativesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetAllCreativesResponse"`

	GetAllCreativesResult *ArrayOfDMCreative `xml:"GetAllCreativesResult,omitempty"`
}

type GetCreativeVariables struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativeVariables"`

	Token string `xml:"Token,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`
}

type GetCreativeVariablesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativeVariablesResponse"`

	GetCreativeVariablesResult *ArrayOfDMSAMVariable `xml:"GetCreativeVariablesResult,omitempty"`
}

type GetVariableComboValues struct {
	XMLName xml.Name `xml:"DMWebServices GetVariableComboValues"`

	Token string `xml:"Token,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`

	VariableID int32 `xml:"VariableID,omitempty"`
}

type GetVariableComboValuesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetVariableComboValuesResponse"`

	GetVariableComboValuesResult *ArrayOfDMComboValue `xml:"GetVariableComboValuesResult,omitempty"`
}

type GetListCategories struct {
	XMLName xml.Name `xml:"DMWebServices GetListCategories"`

	Token string `xml:"Token,omitempty"`
}

type GetListCategoriesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListCategoriesResponse"`

	GetListCategoriesResult *ArrayOfDMCategory `xml:"GetListCategoriesResult,omitempty"`
}

type SAMGetLists struct {
	XMLName xml.Name `xml:"DMWebServices SAMGetLists"`

	Token string `xml:"Token,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	DefaultListsOnly bool `xml:"DefaultListsOnly,omitempty"`
}

type SAMGetListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMGetListsResponse"`

	SAMGetListsResult *ArrayOfDMSAMList `xml:"SAMGetListsResult,omitempty"`
}

type SAMGetAllLists struct {
	XMLName xml.Name `xml:"DMWebServices SAMGetAllLists"`

	Token string `xml:"Token,omitempty"`

	DefaultListsOnly bool `xml:"DefaultListsOnly,omitempty"`
}

type SAMGetAllListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMGetAllListsResponse"`

	SAMGetAllListsResult *ArrayOfDMSAMList `xml:"SAMGetAllListsResult,omitempty"`
}

type GetSelectedLists struct {
	XMLName xml.Name `xml:"DMWebServices GetSelectedLists"`

	Token string `xml:"Token,omitempty"`

	ListIDCollection *ArrayOfInt `xml:"ListIDCollection,omitempty"`
}

type GetSelectedListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetSelectedListsResponse"`

	GetSelectedListsResult *ArrayOfDMSAMList `xml:"GetSelectedListsResult,omitempty"`
}

type SAMGetSuppressionLists struct {
	XMLName xml.Name `xml:"DMWebServices SAMGetSuppressionLists"`

	Token string `xml:"Token,omitempty"`

	DefaultListsOnly bool `xml:"DefaultListsOnly,omitempty"`
}

type SAMGetSuppressionListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices SAMGetSuppressionListsResponse"`

	SAMGetSuppressionListsResult *ArrayOfDMSAMList `xml:"SAMGetSuppressionListsResult,omitempty"`
}

type GetCompatibleFields struct {
	XMLName xml.Name `xml:"DMWebServices GetCompatibleFields"`

	Token string `xml:"Token,omitempty"`

	VariableType *DMVariableType `xml:"VariableType,omitempty"`

	Lists *ArrayOfInt `xml:"Lists,omitempty"`
}

type GetCompatibleFieldsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetCompatibleFieldsResponse"`

	GetCompatibleFieldsResult *ArrayOfDMCompatibleField `xml:"GetCompatibleFieldsResult,omitempty"`
}

type GetAllCompatibleFields struct {
	XMLName xml.Name `xml:"DMWebServices GetAllCompatibleFields"`

	Token string `xml:"Token,omitempty"`

	VariableType *DMVariableType `xml:"VariableType,omitempty"`
}

type GetAllCompatibleFieldsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetAllCompatibleFieldsResponse"`

	GetAllCompatibleFieldsResult *ArrayOfDMCompatibleField `xml:"GetAllCompatibleFieldsResult,omitempty"`
}

type GetAttachments struct {
	XMLName xml.Name `xml:"DMWebServices GetAttachments"`

	Token string `xml:"Token,omitempty"`
}

type GetAttachmentsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetAttachmentsResponse"`

	GetAttachmentsResult *ArrayOfDMSAMAttachment `xml:"GetAttachmentsResult,omitempty"`
}

type CreateAttachment struct {
	XMLName xml.Name `xml:"DMWebServices CreateAttachment"`

	Token string `xml:"Token,omitempty"`

	FileName string `xml:"FileName,omitempty"`

	Description string `xml:"Description,omitempty"`

	Data []byte `xml:"Data,omitempty"`
}

type CreateAttachmentResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateAttachmentResponse"`

	CreateAttachmentResult int32 `xml:"CreateAttachmentResult,omitempty"`
}

type UpdateAttachment struct {
	XMLName xml.Name `xml:"DMWebServices UpdateAttachment"`

	Token string `xml:"Token,omitempty"`

	AttachmentID int32 `xml:"AttachmentID,omitempty"`

	FileName string `xml:"FileName,omitempty"`

	Description string `xml:"Description,omitempty"`

	Data []byte `xml:"Data,omitempty"`
}

type UpdateAttachmentResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateAttachmentResponse"`
}

type DeleteAttachment struct {
	XMLName xml.Name `xml:"DMWebServices DeleteAttachment"`

	Token string `xml:"Token,omitempty"`

	AttachmentID int32 `xml:"AttachmentID,omitempty"`
}

type DeleteAttachmentResponse struct {
	XMLName xml.Name `xml:"DMWebServices DeleteAttachmentResponse"`
}

type GetTemplates struct {
	XMLName xml.Name `xml:"DMWebServices GetTemplates"`

	Token string `xml:"Token,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`
}

type GetTemplatesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetTemplatesResponse"`

	GetTemplatesResult *ArrayOfDMSAMTemplate `xml:"GetTemplatesResult,omitempty"`
}

type GetHTMLTemplates struct {
	XMLName xml.Name `xml:"DMWebServices GetHTMLTemplates"`

	Token string `xml:"Token,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`
}

type GetHTMLTemplatesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetHTMLTemplatesResponse"`

	GetHTMLTemplatesResult *ArrayOfDMSAMTemplate `xml:"GetHTMLTemplatesResult,omitempty"`
}

type GetTextTemplates struct {
	XMLName xml.Name `xml:"DMWebServices GetTextTemplates"`

	Token string `xml:"Token,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`
}

type GetTextTemplatesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetTextTemplatesResponse"`

	GetTextTemplatesResult *ArrayOfDMSAMTemplate `xml:"GetTextTemplatesResult,omitempty"`
}

type GetInboundAddresses struct {
	XMLName xml.Name `xml:"DMWebServices GetInboundAddresses"`

	Token string `xml:"Token,omitempty"`
}

type GetInboundAddressesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetInboundAddressesResponse"`

	GetInboundAddressesResult *ArrayOfDMInboundAddress `xml:"GetInboundAddressesResult,omitempty"`

	Default int32 `xml:"Default,omitempty"`
}

type CreateInboundAddress struct {
	XMLName xml.Name `xml:"DMWebServices CreateInboundAddress"`

	Token string `xml:"Token,omitempty"`

	Email string `xml:"Email,omitempty"`

	Forward string `xml:"Forward,omitempty"`

	Default bool `xml:"Default,omitempty"`

	ForwardEvents *ArrayOfDMInboundEventForward `xml:"ForwardEvents,omitempty"`
}

type CreateInboundAddressResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateInboundAddressResponse"`

	CreateInboundAddressResult int32 `xml:"CreateInboundAddressResult,omitempty"`
}

type UpdateInboundAddress struct {
	XMLName xml.Name `xml:"DMWebServices UpdateInboundAddress"`

	Token string `xml:"Token,omitempty"`

	ID int32 `xml:"ID,omitempty"`

	Email string `xml:"Email,omitempty"`

	Forward string `xml:"Forward,omitempty"`

	Default bool `xml:"Default,omitempty"`

	ForwardEvents *ArrayOfDMInboundEventForward `xml:"ForwardEvents,omitempty"`
}

type UpdateInboundAddressResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateInboundAddressResponse"`
}

type DeleteInboundAddress struct {
	XMLName xml.Name `xml:"DMWebServices DeleteInboundAddress"`

	Token string `xml:"Token,omitempty"`

	ID int32 `xml:"ID,omitempty"`
}

type DeleteInboundAddressResponse struct {
	XMLName xml.Name `xml:"DMWebServices DeleteInboundAddressResponse"`
}

type DMTemplateValue struct {
	XMLName xml.Name `xml:"DMWebServices DMTemplateValue"`

	HtmlTemplate int32 `xml:"HtmlTemplate,omitempty"`

	TextTemplate int32 `xml:"TextTemplate,omitempty"`

	ConditionalValues *ArrayOfDMConditionalTemplate `xml:"ConditionalValues,omitempty"`
}

type ArrayOfDMConditionalTemplate struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMConditionalTemplate"`

	DMConditionalTemplate []*DMConditionalTemplate `xml:"DMConditionalTemplate,omitempty"`
}

type DMConditionalTemplate struct {
	XMLName xml.Name `xml:"DMWebServices DMConditionalTemplate"`

	HtmlTemplate int32 `xml:"HtmlTemplate,omitempty"`

	TextTemplate int32 `xml:"TextTemplate,omitempty"`

	Conditions *ArrayOfDMCondition `xml:"Conditions,omitempty"`
}

type ArrayOfDMCondition struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMCondition"`

	DMCondition []*DMCondition `xml:"DMCondition,omitempty"`
}

type DMCondition struct {
	XMLName xml.Name `xml:"DMWebServices DMCondition"`

	Operand1 string `xml:"Operand1,omitempty"`

	Operator *DMOperator `xml:"Operator,omitempty"`

	Operand2 string `xml:"Operand2,omitempty"`

	Combine *DMCombine `xml:"Combine,omitempty"`
}

type ArrayOfInt struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfInt"`

	Int []int32 `xml:"int,omitempty"`
}

type ArrayOfDMVariableMap struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMVariableMap"`

	DMVariableMap []*DMVariableMap `xml:"DMVariableMap,omitempty"`
}

type DMVariableMap struct {
	XMLName xml.Name `xml:"DMWebServices DMVariableMap"`

	VariableID int32 `xml:"VariableID,omitempty"`

	Value *DMVariableValue `xml:"Value,omitempty"`

	FieldID int32 `xml:"FieldID,omitempty"`
}

type DMVariableValue struct {
	XMLName xml.Name `xml:"DMWebServices DMVariableValue"`

	Value string `xml:"Value,omitempty"`

	ConditionalValues *ArrayOfDMConditionalValue `xml:"ConditionalValues,omitempty"`
}

type ArrayOfDMConditionalValue struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMConditionalValue"`

	DMConditionalValue []*DMConditionalValue `xml:"DMConditionalValue,omitempty"`
}

type DMConditionalValue struct {
	XMLName xml.Name `xml:"DMWebServices DMConditionalValue"`

	Value string `xml:"Value,omitempty"`

	Conditions *ArrayOfDMCondition `xml:"Conditions,omitempty"`
}

type ArrayOfDMEventNotify struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMEventNotify"`

	DMEventNotify []*DMEventNotify `xml:"DMEventNotify,omitempty"`
}

type DMEventNotify struct {
	XMLName xml.Name `xml:"DMWebServices DMEventNotify"`

	EventID int32 `xml:"EventID,omitempty"`

	EventName string `xml:"EventName,omitempty"`

	Notify bool `xml:"Notify,omitempty"`

	Default bool `xml:"Default,omitempty"`
}

type DMDeployment struct {
	XMLName xml.Name `xml:"DMWebServices DMDeployment"`

	Client *DMClient `xml:"Client,omitempty"`

	ID int32 `xml:"ID,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`

	CreativeName string `xml:"CreativeName,omitempty"`

	SenderID int32 `xml:"SenderID,omitempty"`

	Status *DMDeploymentStatus `xml:"Status,omitempty"`

	Total int32 `xml:"Total,omitempty"`

	Sent int32 `xml:"Sent,omitempty"`

	Choke int32 `xml:"Choke,omitempty"`

	ChokeInterval int32 `xml:"ChokeInterval,omitempty"`

	OneOff bool `xml:"OneOff,omitempty"`

	ProofMode bool `xml:"ProofMode,omitempty"`

	DeliveryContextID int32 `xml:"DeliveryContextID,omitempty"`

	DeliveryContextVirtualMTA string `xml:"DeliveryContextVirtualMTA,omitempty"`

	Notify bool `xml:"Notify,omitempty"`

	Started time.Time `xml:"Started,omitempty"`

	PagesServed int32 `xml:"PagesServed,omitempty"`

	SMSShortCode string `xml:"SMSShortCode,omitempty"`

	SMSUsername string `xml:"SMSUsername,omitempty"`

	SMSPassword string `xml:"SMSPassword,omitempty"`

	SMSTariffClass string `xml:"SMSTariffClass,omitempty"`

	CreativeCategoryName string `xml:"CreativeCategoryName,omitempty"`

	Name string `xml:"Name,omitempty"`

	ListNames string `xml:"ListNames,omitempty"`

	UserName string `xml:"UserName,omitempty"`

	FromEmail string `xml:"FromEmail,omitempty"`

	FromAlias string `xml:"FromAlias,omitempty"`

	ToEmail string `xml:"ToEmail,omitempty"`

	ToAlias string `xml:"ToAlias,omitempty"`

	ReplyTo string `xml:"ReplyTo,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	Finished time.Time `xml:"Finished,omitempty"`

	Schedule time.Time `xml:"Schedule,omitempty"`

	Queued int32 `xml:"Queued,omitempty"`

	PercentDone int32 `xml:"PercentDone,omitempty"`

	HealthThreshold int32 `xml:"HealthThreshold,omitempty"`

	Throttle string `xml:"Throttle,omitempty"`

	OldStatus *DMDeploymentStatus `xml:"OldStatus,omitempty"`

	ListTotal int32 `xml:"ListTotal,omitempty"`

	RecipientSuppressed int32 `xml:"RecipientSuppressed,omitempty"`

	ListSuppressed int32 `xml:"ListSuppressed,omitempty"`

	HealthSuppressed int32 `xml:"HealthSuppressed,omitempty"`

	EventSuppressed int32 `xml:"EventSuppressed,omitempty"`

	FieldSuppressed int32 `xml:"FieldSuppressed,omitempty"`

	MailerSubmissionError int32 `xml:"MailerSubmissionError,omitempty"`

	DCInfo string `xml:"DCInfo,omitempty"`

	SuppressionLists *ArrayOfInt `xml:"SuppressionLists,omitempty"`

	SuppressEvents *ArrayOfInt `xml:"SuppressEvents,omitempty"`

	NotifyEvents *ArrayOfInt `xml:"NotifyEvents,omitempty"`

	RefreshSmartLists bool `xml:"RefreshSmartLists,omitempty"`

	RecurrenceSchedule *DMDeploymentSchedule `xml:"RecurrenceSchedule,omitempty"`

	NotifyEmail string `xml:"NotifyEmail,omitempty"`

	DuplicatesSuppressed int32 `xml:"DuplicatesSuppressed,omitempty"`

	LastChecked time.Time `xml:"LastChecked,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`

	TotalMessagesLeft int32 `xml:"TotalMessagesLeft,omitempty"`
}

type DMClient struct {
	XMLName xml.Name `xml:"DMWebServices DMClient"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	VirtualMTA string `xml:"VirtualMTA,omitempty"`

	ConnectionString string `xml:"ConnectionString,omitempty"`

	HashTable string `xml:"HashTable,omitempty"`

	ContentPath *ArrayOfString `xml:"ContentPath,omitempty"`

	ContentURL string `xml:"ContentURL,omitempty"`

	Seats int32 `xml:"Seats,omitempty"`

	SetID int32 `xml:"SetID,omitempty"`

	Icon string `xml:"Icon,omitempty"`

	ContentServerID int32 `xml:"ContentServerID,omitempty"`

	Directory string `xml:"Directory,omitempty"`

	DBServerID int32 `xml:"DBServerID,omitempty"`

	DBName string `xml:"DBName,omitempty"`

	MTAServerID int32 `xml:"MTAServerID,omitempty"`

	DefaultDC int32 `xml:"DefaultDC,omitempty"`

	DeliveryContexts *ArrayOfInt `xml:"DeliveryContexts,omitempty"`

	Settings *DMClientSettings `xml:"Settings,omitempty"`

	ImageBasePath string `xml:"ImageBasePath,omitempty"`

	PasswordPolicy *DMPasswordPolicy `xml:"PasswordPolicy,omitempty"`

	StaticTokenEnabled bool `xml:"StaticTokenEnabled,omitempty"`

	NotifyEmail string `xml:"NotifyEmail,omitempty"`

	ShareLookUpOnSaveAs string `xml:"ShareLookUpOnSaveAs,omitempty"`

	SEVExportEnabled bool `xml:"SEVExportEnabled,omitempty"`
}

type ArrayOfString struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfString"`

	String []string `xml:"string,omitempty"`
}

type DMClientSettings struct {
	XMLName xml.Name `xml:"DMWebServices DMClientSettings"`

	InboundDomain string `xml:"InboundDomain,omitempty"`

	ContentDomain string `xml:"ContentDomain,omitempty"`

	UploadImageDefault bool `xml:"UploadImageDefault,omitempty"`

	DMWCMList *ArrayOfDMWCMEntry `xml:"DMWCMList,omitempty"`

	PVID int32 `xml:"PVID,omitempty"`

	UserCulture string `xml:"UserCulture,omitempty"`

	OmnitureSettings *DMOmnitureSettings `xml:"OmnitureSettings,omitempty"`

	DefaultTemplateCodePage int32 `xml:"DefaultTemplateCodePage,omitempty"`

	PVUrl string `xml:"PVUrl,omitempty"`

	SM2Url string `xml:"SM2Url,omitempty"`

	MMUrl string `xml:"MMUrl,omitempty"`

	MMCertThumbprint string `xml:"MMCertThumbprint,omitempty"`

	GalleryPreviewEnabled bool `xml:"GalleryPreviewEnabled,omitempty"`

	StaticTokenEnabled bool `xml:"StaticTokenEnabled,omitempty"`

	NotifyEmail string `xml:"NotifyEmail,omitempty"`

	ShareLookUpOnSaveAs bool `xml:"ShareLookUpOnSaveAs,omitempty"`

	RecipientsSubscribeCommandTimout int32 `xml:"RecipientsSubscribeCommandTimout,omitempty"`

	SEVExportEnabled bool `xml:"SEVExportEnabled,omitempty"`

	MaxUpdatedDeploymentsToProcessByPrefetch int32 `xml:"MaxUpdatedDeploymentsToProcessByPrefetch,omitempty"`

	DomainFilteringEnabled bool `xml:"DomainFilteringEnabled,omitempty"`

	TrackLastEngagementDate bool `xml:"TrackLastEngagementDate,omitempty"`

	ProcessSuppressionChunksMaxDegreeOfParallelism byte `xml:"ProcessSuppressionChunksMaxDegreeOfParallelism,omitempty"`
}

type ArrayOfDMWCMEntry struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMWCMEntry"`

	DMWCMEntry []*DMWCMEntry `xml:"DMWCMEntry,omitempty"`
}

type DMWCMEntry struct {
	XMLName xml.Name `xml:"DMWebServices DMWCMEntry"`

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type DMOmnitureSettings struct {
	XMLName xml.Name `xml:"DMWebServices DMOmnitureSettings"`

	Enabled bool `xml:"Enabled,omitempty"`

	IntegrationNum string `xml:"IntegrationNum,omitempty"`

	CompanyName string `xml:"CompanyName,omitempty"`

	ReportSuite string `xml:"ReportSuite,omitempty"`

	RemarketingSegPath string `xml:"RemarketingSegPath,omitempty"`

	RemarketingSeg string `xml:"RemarketingSeg,omitempty"`

	EmailAddr string `xml:"EmailAddr,omitempty"`
}

type DMPasswordPolicy struct {
	XMLName xml.Name `xml:"DMWebServices DMPasswordPolicy"`

	ClientID int32 `xml:"ClientID,omitempty"`

	IsUpperCaseRequired bool `xml:"IsUpperCaseRequired,omitempty"`

	IsLowerCaseRequired bool `xml:"IsLowerCaseRequired,omitempty"`

	IsDigitRequired bool `xml:"IsDigitRequired,omitempty"`

	IsSpecialCharRequired bool `xml:"IsSpecialCharRequired,omitempty"`

	PasswordAge int32 `xml:"PasswordAge,omitempty"`

	MinimumLength int32 `xml:"MinimumLength,omitempty"`

	MaximumLength int32 `xml:"MaximumLength,omitempty"`

	PasswordHistoryMax int32 `xml:"PasswordHistoryMax,omitempty"`

	IsPreviousPasswordsAllowed bool `xml:"IsPreviousPasswordsAllowed,omitempty"`
}

type DMDeploymentSchedule struct {
	XMLName xml.Name `xml:"DMWebServices DMDeploymentSchedule"`

	Frequency *DMRecurringScheduleType `xml:"Frequency,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty"`
}

type ArrayOfDMList struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMList"`

	DMList []*DMList `xml:"DMList,omitempty"`
}

type DMList struct {
	XMLName xml.Name `xml:"DMWebServices DMList"`

	ID int32 `xml:"ID,omitempty"`

	OptinEventID int32 `xml:"OptinEventID,omitempty"`

	OptoutEventID int32 `xml:"OptoutEventID,omitempty"`

	ListType *DMListType `xml:"ListType,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	PrimaryKey int32 `xml:"PrimaryKey,omitempty"`

	AllowOptin bool `xml:"AllowOptin,omitempty"`

	AllowOptout bool `xml:"AllowOptout,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	EditPropertiesAccess bool `xml:"EditPropertiesAccess,omitempty"`

	EditRecordsAccess bool `xml:"EditRecordsAccess,omitempty"`

	DeployAccess bool `xml:"DeployAccess,omitempty"`

	ListCriteria *DMListCriteria `xml:"ListCriteria,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type DMListCriteria struct {
	XMLName xml.Name `xml:"DMWebServices DMListCriteria"`

	IncludeLists *ArrayOfInt `xml:"IncludeLists,omitempty"`

	ExcludeLists *ArrayOfInt `xml:"ExcludeLists,omitempty"`

	EventCriteria *DMListEventCriteria `xml:"EventCriteria,omitempty"`

	FieldCriteria *ArrayOfDMFieldCriteria `xml:"FieldCriteria,omitempty"`
}

type DMListEventCriteria struct {
	XMLName xml.Name `xml:"DMWebServices DMListEventCriteria"`

	Creatives *ArrayOfInt `xml:"Creatives,omitempty"`

	Users *ArrayOfInt `xml:"Users,omitempty"`

	DateRanges *ArrayOfDMRDateRange `xml:"DateRanges,omitempty"`

	Events *ArrayOfDMEventCriteria `xml:"Events,omitempty"`
}

type ArrayOfDMRDateRange struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMRDateRange"`

	DMRDateRange []*DMRDateRange `xml:"DMRDateRange,omitempty"`
}

type DMRDateRange struct {
	XMLName xml.Name `xml:"DMWebServices DMRDateRange"`

	RangeType *DMRDateRangeType `xml:"RangeType,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`
}

type ArrayOfDMEventCriteria struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMEventCriteria"`

	DMEventCriteria []*DMEventCriteria `xml:"DMEventCriteria,omitempty"`
}

type DMEventCriteria struct {
	XMLName xml.Name `xml:"DMWebServices DMEventCriteria"`

	OpenParenthesis int32 `xml:"OpenParenthesis,omitempty"`

	CloseParenthesis int32 `xml:"CloseParenthesis,omitempty"`

	Operator *DMOperator `xml:"Operator,omitempty"`

	Combine *DMCombine `xml:"Combine,omitempty"`

	EventID int32 `xml:"EventID,omitempty"`

	Operand int32 `xml:"Operand,omitempty"`
}

type ArrayOfDMFieldCriteria struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMFieldCriteria"`

	DMFieldCriteria []*DMFieldCriteria `xml:"DMFieldCriteria,omitempty"`
}

type DMFieldCriteria struct {
	XMLName xml.Name `xml:"DMWebServices DMFieldCriteria"`

	OpenParenthesis int32 `xml:"OpenParenthesis,omitempty"`

	CloseParenthesis int32 `xml:"CloseParenthesis,omitempty"`

	QueryType *DMFieldQueryType `xml:"QueryType,omitempty"`

	Parameter int32 `xml:"Parameter,omitempty"`

	Operator *DMSQLOperator `xml:"Operator,omitempty"`

	Combine *DMCombine `xml:"Combine,omitempty"`

	SValue *ArrayOfString `xml:"SValue,omitempty"`

	IValue *ArrayOfInt `xml:"IValue,omitempty"`

	DValue *ArrayOfDouble `xml:"DValue,omitempty"`

	BValue *ArrayOfBoolean `xml:"BValue,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`

	FieldSubType *DMFieldSubType `xml:"FieldSubType,omitempty"`
}

type ArrayOfDouble struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDouble"`

	Double []float64 `xml:"double,omitempty"`
}

type ArrayOfBoolean struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfBoolean"`

	Boolean []bool `xml:"boolean,omitempty"`
}

type ArrayOfDMDeploymentLists struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMDeploymentLists"`

	DMDeploymentLists []*DMDeploymentLists `xml:"DMDeploymentLists,omitempty"`
}

type DMDeploymentLists struct {
	XMLName xml.Name `xml:"DMWebServices DMDeploymentLists"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	Include *ArrayOfInt `xml:"Include,omitempty"`

	Exclude *ArrayOfInt `xml:"Exclude,omitempty"`

	Suppress *ArrayOfInt `xml:"Suppress,omitempty"`
}

type ArrayOfDMEventInfo struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMEventInfo"`

	DMEventInfo []*DMEventInfo `xml:"DMEventInfo,omitempty"`
}

type DMEventInfo struct {
	XMLName xml.Name `xml:"DMWebServices DMEventInfo"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Health int32 `xml:"Health,omitempty"`

	System bool `xml:"System,omitempty"`

	UserDefined bool `xml:"UserDefined,omitempty"`

	AutoNamed bool `xml:"AutoNamed,omitempty"`

	Shared bool `xml:"Shared,omitempty"`

	HasValue bool `xml:"HasValue,omitempty"`

	Disable bool `xml:"Disable,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	URL string `xml:"URL,omitempty"`
}

type ArrayOfDMDeploymentEvents struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMDeploymentEvents"`

	DMDeploymentEvents []*DMDeploymentEvents `xml:"DMDeploymentEvents,omitempty"`
}

type DMDeploymentEvents struct {
	XMLName xml.Name `xml:"DMWebServices DMDeploymentEvents"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`

	EventIDs *ArrayOfInt `xml:"EventIDs,omitempty"`
}

type ArrayOfDMDeploymentField struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMDeploymentField"`

	DMDeploymentField []*DMDeploymentField `xml:"DMDeploymentField,omitempty"`
}

type ArrayOfDMDeployment struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMDeployment"`

	DMDeployment []*DMDeployment `xml:"DMDeployment,omitempty"`
}

type ArrayOfHealthCount struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfHealthCount"`

	HealthCount []*HealthCount `xml:"HealthCount,omitempty"`
}

type HealthCount struct {
	XMLName xml.Name `xml:"DMWebServices HealthCount"`

	Health int32 `xml:"Health,omitempty"`

	Count int32 `xml:"Count,omitempty"`
}

type SAReport struct {
	XMLName xml.Name `xml:"DMWebServices SAReport"`

	Score string `xml:"score,omitempty"`

	Rows *ArrayOfSAReportRow `xml:"Rows,omitempty"`
}

type ArrayOfSAReportRow struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfSAReportRow"`

	SAReportRow []*SAReportRow `xml:"SAReportRow,omitempty"`
}

type SAReportRow struct {
	XMLName xml.Name `xml:"DMWebServices SAReportRow"`

	Points string `xml:"points,omitempty"`

	Rule string `xml:"rule,omitempty"`

	Desc string `xml:"desc,omitempty"`
}

type ArrayOfDMCategory struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMCategory"`

	DMCategory []*DMCategory `xml:"DMCategory,omitempty"`
}

type DMCategory struct {
	XMLName xml.Name `xml:"DMWebServices DMCategory"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type ArrayOfDMCreative struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMCreative"`

	DMCreative []*DMCreative `xml:"DMCreative,omitempty"`
}

type DMCreative struct {
	XMLName xml.Name `xml:"DMWebServices DMCreative"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	Thumbnail string `xml:"Thumbnail,omitempty"`

	TemplateDefaults *DMTemplateValue `xml:"TemplateDefaults,omitempty"`

	EditDefaults bool `xml:"EditDefaults,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	ModifyAccess bool `xml:"ModifyAccess,omitempty"`

	DeployAccess bool `xml:"DeployAccess,omitempty"`

	ReportAccess bool `xml:"ReportAccess,omitempty"`
}

type ArrayOfDMCreativePermission struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMCreativePermission"`

	DMCreativePermission []*DMCreativePermission `xml:"DMCreativePermission,omitempty"`
}

type ArrayOfDMCreativeCategory struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMCreativeCategory"`

	DMCreativeCategory []*DMCreativeCategory `xml:"DMCreativeCategory,omitempty"`
}

type DMCreativeCategory struct {
	XMLName xml.Name `xml:"DMWebServices DMCreativeCategory"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Creatives *ArrayOfDMCreative `xml:"Creatives,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty"`
}

type ArrayOfDMSAMVariable struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMSAMVariable"`

	DMSAMVariable []*DMSAMVariable `xml:"DMSAMVariable,omitempty"`
}

type DMSAMVariable struct {
	XMLName xml.Name `xml:"DMWebServices DMSAMVariable"`

	ID int32 `xml:"ID,omitempty"`

	VariableType *DMVariableType `xml:"VariableType,omitempty"`

	EditorType *DMEditorType `xml:"EditorType,omitempty"`

	Name string `xml:"Name,omitempty"`

	ParseName string `xml:"ParseName,omitempty"`

	DefaultValue *DMVariableValue `xml:"DefaultValue,omitempty"`

	ForceMap bool `xml:"ForceMap,omitempty"`

	DefaultField int32 `xml:"DefaultField,omitempty"`

	Required bool `xml:"Required,omitempty"`

	Readonly bool `xml:"Readonly,omitempty"`
}

type ArrayOfDMComboValue struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMComboValue"`

	DMComboValue []*DMComboValue `xml:"DMComboValue,omitempty"`
}

type DMComboValue struct {
	XMLName xml.Name `xml:"DMWebServices DMComboValue"`

	Value string `xml:"Value,omitempty"`

	Display string `xml:"Display,omitempty"`
}

type ArrayOfDMSAMList struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMSAMList"`

	DMSAMList []*DMSAMList `xml:"DMSAMList,omitempty"`
}

type DMSAMList struct {
	XMLName xml.Name `xml:"DMWebServices DMSAMList"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type ArrayOfDMCompatibleField struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMCompatibleField"`

	DMCompatibleField []*DMCompatibleField `xml:"DMCompatibleField,omitempty"`
}

type DMCompatibleField struct {
	XMLName xml.Name `xml:"DMWebServices DMCompatibleField"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	CommonLists *ArrayOfInt `xml:"CommonLists,omitempty"`
}

type ArrayOfDMSAMAttachment struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMSAMAttachment"`

	DMSAMAttachment []*DMSAMAttachment `xml:"DMSAMAttachment,omitempty"`
}

type DMSAMAttachment struct {
	XMLName xml.Name `xml:"DMWebServices DMSAMAttachment"`

	ID int32 `xml:"ID,omitempty"`

	FileName string `xml:"FileName,omitempty"`

	Description string `xml:"Description,omitempty"`

	MessageID string `xml:"MessageID,omitempty"`
}

type ArrayOfDMSAMTemplate struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMSAMTemplate"`

	DMSAMTemplate []*DMSAMTemplate `xml:"DMSAMTemplate,omitempty"`
}

type DMSAMTemplate struct {
	XMLName xml.Name `xml:"DMWebServices DMSAMTemplate"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Text bool `xml:"Text,omitempty"`
}

type ArrayOfDMInboundAddress struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMInboundAddress"`

	DMInboundAddress []*DMInboundAddress `xml:"DMInboundAddress,omitempty"`
}

type DMInboundAddress struct {
	XMLName xml.Name `xml:"DMWebServices DMInboundAddress"`

	ID int32 `xml:"ID,omitempty"`

	Email string `xml:"Email,omitempty"`

	ModifyAccess bool `xml:"ModifyAccess,omitempty"`

	DeployAccess bool `xml:"DeployAccess,omitempty"`

	ForwardEvents *ArrayOfDMInboundEventForward `xml:"ForwardEvents,omitempty"`
}

type ArrayOfDMInboundEventForward struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMInboundEventForward"`

	DMInboundEventForward []*DMInboundEventForward `xml:"DMInboundEventForward,omitempty"`
}

type DMInboundEventForward struct {
	XMLName xml.Name `xml:"DMWebServices DMInboundEventForward"`

	EventId int32 `xml:"EventId,omitempty"`

	Email string `xml:"Email,omitempty"`
}

type DMSendMessageSoap struct {
	client *SOAPClient
}

func NewDMSendMessageSoap(url string, tls bool, auth *BasicAuth) *DMSendMessageSoap {
	if url == "" {
		url = "https://uk56.em.sdlproducts.com/sendmessage.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &DMSendMessageSoap{
		client: client,
	}
}

func (service *DMSendMessageSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

func (service *DMSendMessageSoap) VolumeWarning(request *VolumeWarning) (*VolumeWarningResponse, error) {
	response := new(VolumeWarningResponse)
	err := service.client.Call("DMWebServices/VolumeWarning", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) CreateDeployment(request *CreateDeployment) (*CreateDeploymentResponse, error) {
	response := new(CreateDeploymentResponse)
	err := service.client.Call("DMWebServices/CreateDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) UpdateDeployment(request *UpdateDeployment) (*UpdateDeploymentResponse, error) {
	response := new(UpdateDeploymentResponse)
	err := service.client.Call("DMWebServices/UpdateDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SetDeploymentVariables(request *SetDeploymentVariables) (*SetDeploymentVariablesResponse, error) {
	response := new(SetDeploymentVariablesResponse)
	err := service.client.Call("DMWebServices/SetDeploymentVariables", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SetDeploymentDeliveryContext(request *SetDeploymentDeliveryContext) (*SetDeploymentDeliveryContextResponse, error) {
	response := new(SetDeploymentDeliveryContextResponse)
	err := service.client.Call("DMWebServices/SetDeploymentDeliveryContext", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SetDeploymentSmartlistRefresh(request *SetDeploymentSmartlistRefresh) (*SetDeploymentSmartlistRefreshResponse, error) {
	response := new(SetDeploymentSmartlistRefreshResponse)
	err := service.client.Call("DMWebServices/SetDeploymentSmartlistRefresh", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SetDeploymentSuppression(request *SetDeploymentSuppression) (*SetDeploymentSuppressionResponse, error) {
	response := new(SetDeploymentSuppressionResponse)
	err := service.client.Call("DMWebServices/SetDeploymentSuppression", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetDeploymentNotify(request *GetDeploymentNotify) (*GetDeploymentNotifyResponse, error) {
	response := new(GetDeploymentNotifyResponse)
	err := service.client.Call("DMWebServices/GetDeploymentNotify", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SetDeploymentNotify(request *SetDeploymentNotify) (*SetDeploymentNotifyResponse, error) {
	response := new(SetDeploymentNotifyResponse)
	err := service.client.Call("DMWebServices/SetDeploymentNotify", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SetDeploymentThrottle(request *SetDeploymentThrottle) (*SetDeploymentThrottleResponse, error) {
	response := new(SetDeploymentThrottleResponse)
	err := service.client.Call("DMWebServices/SetDeploymentThrottle", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SetDeploymentSchedule(request *SetDeploymentSchedule) (*SetDeploymentScheduleResponse, error) {
	response := new(SetDeploymentScheduleResponse)
	err := service.client.Call("DMWebServices/SetDeploymentSchedule", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) FinalizeDeployment(request *FinalizeDeployment) (*FinalizeDeploymentResponse, error) {
	response := new(FinalizeDeploymentResponse)
	err := service.client.Call("DMWebServices/FinalizeDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) FinalizeRecurringDeployment(request *FinalizeRecurringDeployment) (*FinalizeRecurringDeploymentResponse, error) {
	response := new(FinalizeRecurringDeploymentResponse)
	err := service.client.Call("DMWebServices/FinalizeRecurringDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) UpdateRecurringDeployment(request *UpdateRecurringDeployment) (*UpdateRecurringDeploymentResponse, error) {
	response := new(UpdateRecurringDeploymentResponse)
	err := service.client.Call("DMWebServices/UpdateRecurringDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetDeploymentLists(request *GetDeploymentLists) (*GetDeploymentListsResponse, error) {
	response := new(GetDeploymentListsResponse)
	err := service.client.Call("DMWebServices/GetDeploymentLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetDeploymentsLists(request *GetDeploymentsLists) (*GetDeploymentsListsResponse, error) {
	response := new(GetDeploymentsListsResponse)
	err := service.client.Call("DMWebServices/GetDeploymentsLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetDeploymentsEvents(request *GetDeploymentsEvents) (*GetDeploymentsEventsResponse, error) {
	response := new(GetDeploymentsEventsResponse)
	err := service.client.Call("DMWebServices/GetDeploymentsEvents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetDeployment(request *GetDeployment) (*GetDeploymentResponse, error) {
	response := new(GetDeploymentResponse)
	err := service.client.Call("DMWebServices/GetDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) AddRecipientToDeployment(request *AddRecipientToDeployment) (*AddRecipientToDeploymentResponse, error) {
	response := new(AddRecipientToDeploymentResponse)
	err := service.client.Call("DMWebServices/AddRecipientToDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetDeployments(request *GetDeployments) (*GetDeploymentsResponse, error) {
	response := new(GetDeploymentsResponse)
	err := service.client.Call("DMWebServices/GetDeployments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) PauseDeployment(request *PauseDeployment) (*PauseDeploymentResponse, error) {
	response := new(PauseDeploymentResponse)
	err := service.client.Call("DMWebServices/PauseDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) PauseRecurringDeployment(request *PauseRecurringDeployment) (*PauseRecurringDeploymentResponse, error) {
	response := new(PauseRecurringDeploymentResponse)
	err := service.client.Call("DMWebServices/PauseRecurringDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) ResumeDeployment(request *ResumeDeployment) (*ResumeDeploymentResponse, error) {
	response := new(ResumeDeploymentResponse)
	err := service.client.Call("DMWebServices/ResumeDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) ResumeRecurringDeployment(request *ResumeRecurringDeployment) (*ResumeRecurringDeploymentResponse, error) {
	response := new(ResumeRecurringDeploymentResponse)
	err := service.client.Call("DMWebServices/ResumeRecurringDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) AbortDeployment(request *AbortDeployment) (*AbortDeploymentResponse, error) {
	response := new(AbortDeploymentResponse)
	err := service.client.Call("DMWebServices/AbortDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) AbortRecurringDeployment(request *AbortRecurringDeployment) (*AbortRecurringDeploymentResponse, error) {
	response := new(AbortRecurringDeploymentResponse)
	err := service.client.Call("DMWebServices/AbortRecurringDeployment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMGetTotalRecipients(request *SAMGetTotalRecipients) (*SAMGetTotalRecipientsResponse, error) {
	response := new(SAMGetTotalRecipientsResponse)
	err := service.client.Call("DMWebServices/SAMGetTotalRecipients", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) PreviewDeploymentTotal(request *PreviewDeploymentTotal) (*PreviewDeploymentTotalResponse, error) {
	response := new(PreviewDeploymentTotalResponse)
	err := service.client.Call("DMWebServices/PreviewDeploymentTotal", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMPreviewFirst(request *SAMPreviewFirst) (*SAMPreviewFirstResponse, error) {
	response := new(SAMPreviewFirstResponse)
	err := service.client.Call("DMWebServices/SAMPreviewFirst", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMPreviewLast(request *SAMPreviewLast) (*SAMPreviewLastResponse, error) {
	response := new(SAMPreviewLastResponse)
	err := service.client.Call("DMWebServices/SAMPreviewLast", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMPreviewNext(request *SAMPreviewNext) (*SAMPreviewNextResponse, error) {
	response := new(SAMPreviewNextResponse)
	err := service.client.Call("DMWebServices/SAMPreviewNext", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMPreviewPrevious(request *SAMPreviewPrevious) (*SAMPreviewPreviousResponse, error) {
	response := new(SAMPreviewPreviousResponse)
	err := service.client.Call("DMWebServices/SAMPreviewPrevious", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMPreviewEx(request *SAMPreviewEx) (*SAMPreviewExResponse, error) {
	response := new(SAMPreviewExResponse)
	err := service.client.Call("DMWebServices/SAMPreviewEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetSAReport(request *GetSAReport) (*GetSAReportResponse, error) {
	response := new(GetSAReportResponse)
	err := service.client.Call("DMWebServices/GetSAReport", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetCreativeCategories(request *GetCreativeCategories) (*GetCreativeCategoriesResponse, error) {
	response := new(GetCreativeCategoriesResponse)
	err := service.client.Call("DMWebServices/GetCreativeCategories", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetCreatives(request *GetCreatives) (*GetCreativesResponse, error) {
	response := new(GetCreativesResponse)
	err := service.client.Call("DMWebServices/GetCreatives", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetCreativeCategoriesEx(request *GetCreativeCategoriesEx) (*GetCreativeCategoriesExResponse, error) {
	response := new(GetCreativeCategoriesExResponse)
	err := service.client.Call("DMWebServices/GetCreativeCategoriesEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetAllCreatives(request *GetAllCreatives) (*GetAllCreativesResponse, error) {
	response := new(GetAllCreativesResponse)
	err := service.client.Call("DMWebServices/GetAllCreatives", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetCreativeVariables(request *GetCreativeVariables) (*GetCreativeVariablesResponse, error) {
	response := new(GetCreativeVariablesResponse)
	err := service.client.Call("DMWebServices/GetCreativeVariables", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetVariableComboValues(request *GetVariableComboValues) (*GetVariableComboValuesResponse, error) {
	response := new(GetVariableComboValuesResponse)
	err := service.client.Call("DMWebServices/GetVariableComboValues", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetListCategories(request *GetListCategories) (*GetListCategoriesResponse, error) {
	response := new(GetListCategoriesResponse)
	err := service.client.Call("DMWebServices/GetListCategories", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMGetLists(request *SAMGetLists) (*SAMGetListsResponse, error) {
	response := new(SAMGetListsResponse)
	err := service.client.Call("DMWebServices/SAMGetLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMGetAllLists(request *SAMGetAllLists) (*SAMGetAllListsResponse, error) {
	response := new(SAMGetAllListsResponse)
	err := service.client.Call("DMWebServices/SAMGetAllLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetSelectedLists(request *GetSelectedLists) (*GetSelectedListsResponse, error) {
	response := new(GetSelectedListsResponse)
	err := service.client.Call("DMWebServices/GetSelectedLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) SAMGetSuppressionLists(request *SAMGetSuppressionLists) (*SAMGetSuppressionListsResponse, error) {
	response := new(SAMGetSuppressionListsResponse)
	err := service.client.Call("DMWebServices/SAMGetSuppressionLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetCompatibleFields(request *GetCompatibleFields) (*GetCompatibleFieldsResponse, error) {
	response := new(GetCompatibleFieldsResponse)
	err := service.client.Call("DMWebServices/GetCompatibleFields", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetAllCompatibleFields(request *GetAllCompatibleFields) (*GetAllCompatibleFieldsResponse, error) {
	response := new(GetAllCompatibleFieldsResponse)
	err := service.client.Call("DMWebServices/GetAllCompatibleFields", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetAttachments(request *GetAttachments) (*GetAttachmentsResponse, error) {
	response := new(GetAttachmentsResponse)
	err := service.client.Call("DMWebServices/GetAttachments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) CreateAttachment(request *CreateAttachment) (*CreateAttachmentResponse, error) {
	response := new(CreateAttachmentResponse)
	err := service.client.Call("DMWebServices/CreateAttachment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) UpdateAttachment(request *UpdateAttachment) (*UpdateAttachmentResponse, error) {
	response := new(UpdateAttachmentResponse)
	err := service.client.Call("DMWebServices/UpdateAttachment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) DeleteAttachment(request *DeleteAttachment) (*DeleteAttachmentResponse, error) {
	response := new(DeleteAttachmentResponse)
	err := service.client.Call("DMWebServices/DeleteAttachment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetTemplates(request *GetTemplates) (*GetTemplatesResponse, error) {
	response := new(GetTemplatesResponse)
	err := service.client.Call("DMWebServices/GetTemplates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetHTMLTemplates(request *GetHTMLTemplates) (*GetHTMLTemplatesResponse, error) {
	response := new(GetHTMLTemplatesResponse)
	err := service.client.Call("DMWebServices/GetHTMLTemplates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetTextTemplates(request *GetTextTemplates) (*GetTextTemplatesResponse, error) {
	response := new(GetTextTemplatesResponse)
	err := service.client.Call("DMWebServices/GetTextTemplates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) GetInboundAddresses(request *GetInboundAddresses) (*GetInboundAddressesResponse, error) {
	response := new(GetInboundAddressesResponse)
	err := service.client.Call("DMWebServices/GetInboundAddresses", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) CreateInboundAddress(request *CreateInboundAddress) (*CreateInboundAddressResponse, error) {
	response := new(CreateInboundAddressResponse)
	err := service.client.Call("DMWebServices/CreateInboundAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) UpdateInboundAddress(request *UpdateInboundAddress) (*UpdateInboundAddressResponse, error) {
	response := new(UpdateInboundAddressResponse)
	err := service.client.Call("DMWebServices/UpdateInboundAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMSendMessageSoap) DeleteInboundAddress(request *DeleteInboundAddress) (*DeleteInboundAddressResponse, error) {
	response := new(DeleteInboundAddressResponse)
	err := service.client.Call("DMWebServices/DeleteInboundAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url    string
	tls    bool
	auth   *BasicAuth
	header interface{}
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) SetHeader(header interface{}) {
	s.header = header
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.header != nil {
		envelope.Header = &SOAPHeader{Header: s.header}
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
