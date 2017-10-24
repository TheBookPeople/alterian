package listmanager

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

type DMListType string

const (
	DMListTypeDMLTRECIPIENT DMListType = "DMLTRECIPIENT"

	DMListTypeDMLTSUPPRESS DMListType = "DMLTSUPPRESS"

	DMListTypeDMLTDOMAINSUPPRESSION DMListType = "DMLTDOMAINSUPPRESSION"

	DMListTypeDMLTMD5SUPPRESSION DMListType = "DMLTMD5SUPPRESSION"
)

type DMFieldType string

const (
	DMFieldTypeDMFTEMAIL DMFieldType = "DMFTEMAIL"

	DMFieldTypeDMFTSTRING DMFieldType = "DMFTSTRING"

	DMFieldTypeDMFTTEXT DMFieldType = "DMFTTEXT"

	DMFieldTypeDMFTNUMERIC DMFieldType = "DMFTNUMERIC"

	DMFieldTypeDMFTDATE DMFieldType = "DMFTDATE"

	DMFieldTypeDMFTBOOLEAN DMFieldType = "DMFTBOOLEAN"

	DMFieldTypeDMFTDECIMAL DMFieldType = "DMFTDECIMAL"

	DMFieldTypeDMFTUNICODE DMFieldType = "DMFTUNICODE"
)

type DMListPermission string

const (
	DMListPermissionDMLPEDITPROPERTIES DMListPermission = "DMLPEDITPROPERTIES"

	DMListPermissionDMLPEDITRECORDS DMListPermission = "DMLPEDITRECORDS"

	DMListPermissionDMLPDEPLOY DMListPermission = "DMLPDEPLOY"
)

type DMFieldStorageType string

const (
	DMFieldStorageTypeDMFSNA DMFieldStorageType = "DMFSNA"

	DMFieldStorageTypeDMFSLIST DMFieldStorageType = "DMFSLIST"

	DMFieldStorageTypeDMFSRECIPIENT DMFieldStorageType = "DMFSRECIPIENT"

	DMFieldStorageTypeDMFSRECIPIENTLIST DMFieldStorageType = "DMFSRECIPIENTLIST"
)

type DMPreviewDirection string

const (
	DMPreviewDirectionDMPDFIRST DMPreviewDirection = "DMPDFIRST"

	DMPreviewDirectionDMPDNEXT DMPreviewDirection = "DMPDNEXT"

	DMPreviewDirectionDMPDPREVIOUS DMPreviewDirection = "DMPDPREVIOUS"

	DMPreviewDirectionDMPDLAST DMPreviewDirection = "DMPDLAST"

	DMPreviewDirectionDMPDREFRESH DMPreviewDirection = "DMPDREFRESH"
)

type RecipientSubscribingStatus string

const (
	RecipientSubscribingStatusRecipientSubscribingError RecipientSubscribingStatus = "RecipientSubscribingError"

	RecipientSubscribingStatusRecipientCreated RecipientSubscribingStatus = "RecipientCreated"

	RecipientSubscribingStatusRecipientSkipCreation RecipientSubscribingStatus = "RecipientSkipCreation"

	RecipientSubscribingStatusRecipientResubscribed RecipientSubscribingStatus = "RecipientResubscribed"

	RecipientSubscribingStatusRecipientNotReSubscribed RecipientSubscribingStatus = "RecipientNotReSubscribed"

	RecipientSubscribingStatusRecipientAlreadySubscribed RecipientSubscribingStatus = "RecipientAlreadySubscribed"

	RecipientSubscribingStatusRecipientInvalidPrimaryKey RecipientSubscribingStatus = "RecipientInvalidPrimaryKey"
)

type ExecuteListQuery struct {
	XMLName xml.Name `xml:"DMWebServices ExecuteListQuery"`

	Token string `xml:"Token,omitempty"`

	Criteria *DMListCriteria `xml:"Criteria,omitempty"`
}

type ExecuteListQueryResponse struct {
	XMLName xml.Name `xml:"DMWebServices ExecuteListQueryResponse"`

	ExecuteListQueryResult int32 `xml:"ExecuteListQueryResult,omitempty"`
}

type ExecuteSuppressionListQuery struct {
	XMLName xml.Name `xml:"DMWebServices ExecuteSuppressionListQuery"`

	Token string `xml:"Token,omitempty"`

	Criteria *DMListCriteria `xml:"Criteria,omitempty"`
}

type ExecuteSuppressionListQueryResponse struct {
	XMLName xml.Name `xml:"DMWebServices ExecuteSuppressionListQueryResponse"`

	ExecuteSuppressionListQueryResult int32 `xml:"ExecuteSuppressionListQueryResult,omitempty"`
}

type SaveQueryResults struct {
	XMLName xml.Name `xml:"DMWebServices SaveQueryResults"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	Dividend int32 `xml:"Dividend,omitempty"`

	Divisor int32 `xml:"Divisor,omitempty"`

	Top int32 `xml:"Top,omitempty"`

	Bottom bool `xml:"Bottom,omitempty"`
}

type SaveQueryResultsResponse struct {
	XMLName xml.Name `xml:"DMWebServices SaveQueryResultsResponse"`

	SaveQueryResultsResult int32 `xml:"SaveQueryResultsResult,omitempty"`
}

type SplitQueryResults struct {
	XMLName xml.Name `xml:"DMWebServices SplitQueryResults"`

	Token string `xml:"Token,omitempty"`

	Segments *ArrayOfDMListSegment `xml:"Segments,omitempty"`

	Randomize bool `xml:"Randomize,omitempty"`
}

type SplitQueryResultsResponse struct {
	XMLName xml.Name `xml:"DMWebServices SplitQueryResultsResponse"`
}

type ClearList struct {
	XMLName xml.Name `xml:"DMWebServices ClearList"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type ClearListResponse struct {
	XMLName xml.Name `xml:"DMWebServices ClearListResponse"`
}

type UnsubscribeList struct {
	XMLName xml.Name `xml:"DMWebServices UnsubscribeList"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type UnsubscribeListResponse struct {
	XMLName xml.Name `xml:"DMWebServices UnsubscribeListResponse"`
}

type GetUnsubscribes struct {
	XMLName xml.Name `xml:"DMWebServices GetUnsubscribes"`

	Token string `xml:"Token,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`
}

type GetUnsubscribesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetUnsubscribesResponse"`

	GetUnsubscribesResult *ArrayOfString `xml:"GetUnsubscribesResult,omitempty"`
}

type ResetHealth struct {
	XMLName xml.Name `xml:"DMWebServices ResetHealth"`

	Token string `xml:"Token,omitempty"`

	Domains *ArrayOfString `xml:"Domains,omitempty"`
}

type ResetHealthResponse struct {
	XMLName xml.Name `xml:"DMWebServices ResetHealthResponse"`
}

type GetEmailAddressesByHealth struct {
	XMLName xml.Name `xml:"DMWebServices GetEmailAddressesByHealth"`

	Token string `xml:"Token,omitempty"`

	HealthLow int32 `xml:"HealthLow,omitempty"`

	HealthHigh int32 `xml:"HealthHigh,omitempty"`

	SuppressUnsubscribed bool `xml:"SuppressUnsubscribed,omitempty"`
}

type GetEmailAddressesByHealthResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetEmailAddressesByHealthResponse"`

	GetEmailAddressesByHealthResult *ArrayOfString `xml:"GetEmailAddressesByHealthResult,omitempty"`
}

type GetQueryResults struct {
	XMLName xml.Name `xml:"DMWebServices GetQueryResults"`

	Token string `xml:"Token,omitempty"`

	Fields *ArrayOfInt `xml:"Fields,omitempty"`

	StartIndex int32 `xml:"StartIndex,omitempty"`

	EndIndex int32 `xml:"EndIndex,omitempty"`
}

type GetQueryResultsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetQueryResultsResponse"`

	GetQueryResultsResult *ArrayOfDMRecipientRecord `xml:"GetQueryResultsResult,omitempty"`
}

type GetSuppressionQueryResults struct {
	XMLName xml.Name `xml:"DMWebServices GetSuppressionQueryResults"`

	Token string `xml:"Token,omitempty"`

	StartIndex int32 `xml:"StartIndex,omitempty"`

	EndIndex int32 `xml:"EndIndex,omitempty"`
}

type GetSuppressionQueryResultsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetSuppressionQueryResultsResponse"`

	GetSuppressionQueryResultsResult *ArrayOfDMRecipientRecord `xml:"GetSuppressionQueryResultsResult,omitempty"`
}

type GetQueryResultsCSV struct {
	XMLName xml.Name `xml:"DMWebServices GetQueryResultsCSV"`

	Token string `xml:"Token,omitempty"`

	Fields *ArrayOfInt `xml:"Fields,omitempty"`

	StartIndex int32 `xml:"StartIndex,omitempty"`

	EndIndex int32 `xml:"EndIndex,omitempty"`

	IncludeRecipientID bool `xml:"IncludeRecipientID,omitempty"`
}

type GetQueryResultsCSVResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetQueryResultsCSVResponse"`

	GetQueryResultsCSVResult string `xml:"GetQueryResultsCSVResult,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`
}

type GetSuppressionQueryResultsCSV struct {
	XMLName xml.Name `xml:"DMWebServices GetSuppressionQueryResultsCSV"`

	Token string `xml:"Token,omitempty"`

	StartIndex int32 `xml:"StartIndex,omitempty"`

	EndIndex int32 `xml:"EndIndex,omitempty"`
}

type GetSuppressionQueryResultsCSVResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetSuppressionQueryResultsCSVResponse"`

	GetSuppressionQueryResultsCSVResult string `xml:"GetSuppressionQueryResultsCSVResult,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`
}

type GetListsAndFields struct {
	XMLName xml.Name `xml:"DMWebServices GetListsAndFields"`

	Token string `xml:"Token,omitempty"`

	ListCategoryID int32 `xml:"ListCategoryID,omitempty"`

	ListTypes *ArrayOfDMListType `xml:"ListTypes,omitempty"`
}

type GetListsAndFieldsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListsAndFieldsResponse"`

	Lists *ArrayOfDMQList `xml:"Lists,omitempty"`

	Fields *ArrayOfDMQField `xml:"Fields,omitempty"`

	Categories *ArrayOfDMQListCategory `xml:"Categories,omitempty"`
}

type GetCreativesAndEvents struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativesAndEvents"`

	Token string `xml:"Token,omitempty"`
}

type GetCreativesAndEventsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetCreativesAndEventsResponse"`

	Creatives *ArrayOfDMQCreative `xml:"Creatives,omitempty"`

	Events *ArrayOfDMQEvent `xml:"Events,omitempty"`
}

type GetRecipientByID struct {
	XMLName xml.Name `xml:"DMWebServices GetRecipientByID"`

	Token string `xml:"Token,omitempty"`

	RecipientID int32 `xml:"RecipientID,omitempty"`

	FieldIDs *ArrayOfInt `xml:"FieldIDs,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type GetRecipientByIDResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetRecipientByIDResponse"`

	GetRecipientByIDResult *DMRecipientRecord `xml:"GetRecipientByIDResult,omitempty"`
}

type GetRecipientByPK struct {
	XMLName xml.Name `xml:"DMWebServices GetRecipientByPK"`

	Token string `xml:"Token,omitempty"`

	PrimaryKey int32 `xml:"PrimaryKey,omitempty"`

	Value string `xml:"Value,omitempty"`

	FieldIDs *ArrayOfInt `xml:"FieldIDs,omitempty"`

	CreateNew bool `xml:"CreateNew,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type GetRecipientByPKResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetRecipientByPKResponse"`

	GetRecipientByPKResult *DMRecipientRecord `xml:"GetRecipientByPKResult,omitempty"`
}

type GetRecipientIdByPK struct {
	XMLName xml.Name `xml:"DMWebServices GetRecipientIdByPK"`

	Token string `xml:"token,omitempty"`

	PrimaryKeyFieldId int32 `xml:"primaryKeyFieldId,omitempty"`

	Value string `xml:"value,omitempty"`

	ListID int32 `xml:"listID,omitempty"`
}

type GetRecipientIdByPKResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetRecipientIdByPKResponse"`

	GetRecipientIdByPKResult int32 `xml:"GetRecipientIdByPKResult,omitempty"`
}

type GetVisibleUsers struct {
	XMLName xml.Name `xml:"DMWebServices GetVisibleUsers"`

	Token string `xml:"Token,omitempty"`
}

type GetVisibleUsersResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetVisibleUsersResponse"`

	GetVisibleUsersResult *ArrayOfDMQUser `xml:"GetVisibleUsersResult,omitempty"`
}

type CreateListCategory struct {
	XMLName xml.Name `xml:"DMWebServices CreateListCategory"`

	Token string `xml:"Token,omitempty"`

	CategoryName string `xml:"CategoryName,omitempty"`

	ParentID int32 `xml:"parentID,omitempty"`
}

type CreateListCategoryResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateListCategoryResponse"`

	CreateListCategoryResult int32 `xml:"CreateListCategoryResult,omitempty"`
}

type RenameListCategory struct {
	XMLName xml.Name `xml:"DMWebServices RenameListCategory"`

	Token string `xml:"Token,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	CategoryName string `xml:"CategoryName,omitempty"`
}

type RenameListCategoryResponse struct {
	XMLName xml.Name `xml:"DMWebServices RenameListCategoryResponse"`
}

type DeleteListCategory struct {
	XMLName xml.Name `xml:"DMWebServices DeleteListCategory"`

	Token string `xml:"Token,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`
}

type DeleteListCategoryResponse struct {
	XMLName xml.Name `xml:"DMWebServices DeleteListCategoryResponse"`
}

type GetListCategories struct {
	XMLName xml.Name `xml:"DMWebServices GetListCategories"`

	Token string `xml:"Token,omitempty"`
}

type GetListCategoriesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListCategoriesResponse"`

	GetListCategoriesResult *ArrayOfDMCategory `xml:"GetListCategoriesResult,omitempty"`
}

type GetListCategoriesEx struct {
	XMLName xml.Name `xml:"DMWebServices GetListCategoriesEx"`

	Token string `xml:"Token,omitempty"`

	CategoryIDs *ArrayOfInt `xml:"CategoryIDs,omitempty"`

	PermissionIDs *ArrayOfDMListPermission `xml:"PermissionIDs,omitempty"`

	GetLists bool `xml:"GetLists,omitempty"`
}

type GetListCategoriesExResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListCategoriesExResponse"`

	GetListCategoriesExResult *ArrayOfDMListCategory `xml:"GetListCategoriesExResult,omitempty"`
}

type GetLists struct {
	XMLName xml.Name `xml:"DMWebServices GetLists"`

	Token string `xml:"Token,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	ListTypes *ArrayOfDMListType `xml:"ListTypes,omitempty"`

	DefaultListOnly bool `xml:"DefaultListOnly,omitempty"`
}

type GetListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListsResponse"`

	GetListsResult *ArrayOfDMList `xml:"GetListsResult,omitempty"`
}

type GetListFields struct {
	XMLName xml.Name `xml:"DMWebServices GetListFields"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type GetListFieldsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListFieldsResponse"`

	GetListFieldsResult *ArrayOfDMListField `xml:"GetListFieldsResult,omitempty"`

	PrimaryKey int32 `xml:"PrimaryKey,omitempty"`
}

type GetListRecords struct {
	XMLName xml.Name `xml:"DMWebServices GetListRecords"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`

	Cursor *DMCursor `xml:"Cursor,omitempty"`

	Direction *DMPreviewDirection `xml:"Direction,omitempty"`
}

type GetListRecordsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListRecordsResponse"`

	GetListRecordsResult *ArrayOfDMRecipientRecord `xml:"GetListRecordsResult,omitempty"`

	Cursor *DMCursor `xml:"Cursor,omitempty"`
}

type AddRecipientRecord struct {
	XMLName xml.Name `xml:"DMWebServices AddRecipientRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	PrefersText bool `xml:"PrefersText,omitempty"`

	Unsubscribed bool `xml:"Unsubscribed,omitempty"`

	Values *ArrayOfDMFieldValue `xml:"Values,omitempty"`
}

type AddRecipientRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices AddRecipientRecordResponse"`

	AddRecipientRecordResult int32 `xml:"AddRecipientRecordResult,omitempty"`
}

type AddRecipientRecords struct {
	XMLName xml.Name `xml:"DMWebServices AddRecipientRecords"`

	Token string `xml:"Token,omitempty"`

	Records *ArrayOfDMAddRecipientRecord `xml:"Records,omitempty"`
}

type AddRecipientRecordsResponse struct {
	XMLName xml.Name `xml:"DMWebServices AddRecipientRecordsResponse"`
}

type AddSuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices AddSuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	EmailAddress string `xml:"EmailAddress,omitempty"`
}

type AddSuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices AddSuppressionRecordResponse"`
}

type AddDomainSuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices AddDomainSuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	Domain string `xml:"Domain,omitempty"`
}

type AddDomainSuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices AddDomainSuppressionRecordResponse"`
}

type AddMD5SuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices AddMD5SuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	MD5Hash string `xml:"MD5Hash,omitempty"`
}

type AddMD5SuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices AddMD5SuppressionRecordResponse"`
}

type RemoveRecipient struct {
	XMLName xml.Name `xml:"DMWebServices RemoveRecipient"`

	Token string `xml:"Token,omitempty"`

	ListIDs *ArrayOfInt `xml:"ListIDs,omitempty"`

	RecipientID int32 `xml:"RecipientID,omitempty"`
}

type RemoveRecipientResponse struct {
	XMLName xml.Name `xml:"DMWebServices RemoveRecipientResponse"`
}

type RemoveQueryResultsFromLists struct {
	XMLName xml.Name `xml:"DMWebServices RemoveQueryResultsFromLists"`

	Token string `xml:"Token,omitempty"`

	ListIDs *ArrayOfInt `xml:"ListIDs,omitempty"`
}

type RemoveQueryResultsFromListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices RemoveQueryResultsFromListsResponse"`
}

type RemoveSuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices RemoveSuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	EmailAddress string `xml:"EmailAddress,omitempty"`
}

type RemoveSuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices RemoveSuppressionRecordResponse"`
}

type RemoveDomainSuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices RemoveDomainSuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	Domain string `xml:"Domain,omitempty"`
}

type RemoveDomainSuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices RemoveDomainSuppressionRecordResponse"`
}

type RemoveMD5SuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices RemoveMD5SuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	MD5Hash string `xml:"MD5Hash,omitempty"`
}

type RemoveMD5SuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices RemoveMD5SuppressionRecordResponse"`
}

type UpdateSuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices UpdateSuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	OldEmailAddress string `xml:"OldEmailAddress,omitempty"`

	NewEmailAddress string `xml:"NewEmailAddress,omitempty"`
}

type UpdateSuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateSuppressionRecordResponse"`
}

type UpdateDomainSuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices UpdateDomainSuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	OldDomain string `xml:"OldDomain,omitempty"`

	NewDomain string `xml:"NewDomain,omitempty"`
}

type UpdateDomainSuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateDomainSuppressionRecordResponse"`
}

type UpdateMD5SuppressionRecord struct {
	XMLName xml.Name `xml:"DMWebServices UpdateMD5SuppressionRecord"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	OldMD5Hash string `xml:"OldMD5Hash,omitempty"`

	NewMD5Hash string `xml:"NewMD5Hash,omitempty"`
}

type UpdateMD5SuppressionRecordResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateMD5SuppressionRecordResponse"`
}

type UpdateRecipient struct {
	XMLName xml.Name `xml:"DMWebServices UpdateRecipient"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	RecipientID int32 `xml:"RecipientID,omitempty"`

	PrefersText bool `xml:"PrefersText,omitempty"`

	RSSOnly bool `xml:"RSSOnly,omitempty"`

	Unsubscribed bool `xml:"Unsubscribed,omitempty"`

	FieldValues *ArrayOfDMFieldValue `xml:"FieldValues,omitempty"`
}

type UpdateRecipientResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateRecipientResponse"`

	UpdateRecipientResult int32 `xml:"UpdateRecipientResult,omitempty"`
}

type AddRecipientToList struct {
	XMLName xml.Name `xml:"DMWebServices AddRecipientToList"`

	Token string `xml:"Token,omitempty"`

	RecipientID int32 `xml:"RecipientID,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type AddRecipientToListResponse struct {
	XMLName xml.Name `xml:"DMWebServices AddRecipientToListResponse"`
}

type UpdateField struct {
	XMLName xml.Name `xml:"DMWebServices UpdateField"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	FieldID int32 `xml:"FieldID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type UpdateFieldResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateFieldResponse"`
}

type CreateList struct {
	XMLName xml.Name `xml:"DMWebServices CreateList"`

	Token string `xml:"Token,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	AllowOptout bool `xml:"AllowOptout,omitempty"`

	AllowOptin bool `xml:"AllowOptin,omitempty"`

	ListFields *ArrayOfDMListField `xml:"ListFields,omitempty"`

	ListCriteria *DMListCriteria `xml:"ListCriteria,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type CreateListResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateListResponse"`

	CreateListResult int32 `xml:"CreateListResult,omitempty"`
}

type CreateRecipientList struct {
	XMLName xml.Name `xml:"DMWebServices CreateRecipientList"`

	Token string `xml:"Token,omitempty"`

	ListName string `xml:"ListName,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	AllowOptout bool `xml:"AllowOptout,omitempty"`

	AllowOptin bool `xml:"AllowOptin,omitempty"`

	ListFields *ArrayOfDMListField `xml:"ListFields,omitempty"`

	ListCriteria *DMListCriteria `xml:"ListCriteria,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type CreateRecipientListResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateRecipientListResponse"`

	CreateRecipientListResult int32 `xml:"CreateRecipientListResult,omitempty"`

	ListFields *ArrayOfDMListField `xml:"ListFields,omitempty"`
}

type CreateSuppressionList struct {
	XMLName xml.Name `xml:"DMWebServices CreateSuppressionList"`

	Token string `xml:"Token,omitempty"`

	ListName string `xml:"ListName,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type CreateSuppressionListResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateSuppressionListResponse"`

	CreateSuppressionListResult int32 `xml:"CreateSuppressionListResult,omitempty"`
}

type CreateDomainSuppressionList struct {
	XMLName xml.Name `xml:"DMWebServices CreateDomainSuppressionList"`

	Token string `xml:"Token,omitempty"`

	ListName string `xml:"ListName,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type CreateDomainSuppressionListResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateDomainSuppressionListResponse"`

	CreateDomainSuppressionListResult int32 `xml:"CreateDomainSuppressionListResult,omitempty"`
}

type CreateMD5SuppressionList struct {
	XMLName xml.Name `xml:"DMWebServices CreateMD5SuppressionList"`

	Token string `xml:"Token,omitempty"`

	ListName string `xml:"ListName,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type CreateMD5SuppressionListResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateMD5SuppressionListResponse"`

	CreateMD5SuppressionListResult int32 `xml:"CreateMD5SuppressionListResult,omitempty"`
}

type UpdateList struct {
	XMLName xml.Name `xml:"DMWebServices UpdateList"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	ListName string `xml:"ListName,omitempty"`

	Description string `xml:"Description,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	AllowOptout bool `xml:"AllowOptout,omitempty"`

	AllowOptin bool `xml:"AllowOptin,omitempty"`

	ListCriteria *DMListCriteria `xml:"ListCriteria,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`
}

type UpdateListResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateListResponse"`
}

type DeleteList struct {
	XMLName xml.Name `xml:"DMWebServices DeleteList"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type DeleteListResponse struct {
	XMLName xml.Name `xml:"DMWebServices DeleteListResponse"`
}

type RefreshSmartLists struct {
	XMLName xml.Name `xml:"DMWebServices RefreshSmartLists"`

	Token string `xml:"Token,omitempty"`

	ListIDs *ArrayOfInt `xml:"ListIDs,omitempty"`
}

type RefreshSmartListsResponse struct {
	XMLName xml.Name `xml:"DMWebServices RefreshSmartListsResponse"`
}

type UploadData struct {
	XMLName xml.Name `xml:"DMWebServices UploadData"`

	Token string `xml:"Token,omitempty"`

	Data []byte `xml:"Data,omitempty"`
}

type UploadDataResponse struct {
	XMLName xml.Name `xml:"DMWebServices UploadDataResponse"`
}

type UploadDataA struct {
	XMLName xml.Name `xml:"DMWebServices UploadDataA"`

	Token string `xml:"Token,omitempty"`

	Data string `xml:"Data,omitempty"`
}

type UploadDataAResponse struct {
	XMLName xml.Name `xml:"DMWebServices UploadDataAResponse"`
}

type UploadDataCSV struct {
	XMLName xml.Name `xml:"DMWebServices UploadDataCSV"`

	Token string `xml:"Token,omitempty"`

	Delimiter string `xml:"Delimiter,omitempty"`

	Qualifier string `xml:"Qualifier,omitempty"`

	Data string `xml:"Data,omitempty"`
}

type UploadDataCSVResponse struct {
	XMLName xml.Name `xml:"DMWebServices UploadDataCSVResponse"`
}

type InitializeListImport struct {
	XMLName xml.Name `xml:"DMWebServices InitializeListImport"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	FieldOrder *ArrayOfInt `xml:"FieldOrder,omitempty"`

	Unicode bool `xml:"Unicode,omitempty"`
}

type InitializeListImportResponse struct {
	XMLName xml.Name `xml:"DMWebServices InitializeListImportResponse"`
}

type InsertListData struct {
	XMLName xml.Name `xml:"DMWebServices InsertListData"`

	Token string `xml:"Token,omitempty"`
}

type InsertListDataResponse struct {
	XMLName xml.Name `xml:"DMWebServices InsertListDataResponse"`

	InsertListDataResult *DMImportStatus `xml:"InsertListDataResult,omitempty"`
}

type FinalizeListImport struct {
	XMLName xml.Name `xml:"DMWebServices FinalizeListImport"`

	Token string `xml:"Token,omitempty"`
}

type FinalizeListImportResponse struct {
	XMLName xml.Name `xml:"DMWebServices FinalizeListImportResponse"`
}

type InsertTableData struct {
	XMLName xml.Name `xml:"DMWebServices InsertTableData"`

	Token string `xml:"Token,omitempty"`

	LookupID int32 `xml:"LookupID,omitempty"`

	FieldOrder *ArrayOfInt `xml:"FieldOrder,omitempty"`

	Unicode bool `xml:"Unicode,omitempty"`
}

type InsertTableDataResponse struct {
	XMLName xml.Name `xml:"DMWebServices InsertTableDataResponse"`

	InsertTableDataResult *DMImportStatus `xml:"InsertTableDataResult,omitempty"`
}

type GetLookupTables struct {
	XMLName xml.Name `xml:"DMWebServices GetLookupTables"`

	Token string `xml:"Token,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`
}

type GetLookupTablesResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetLookupTablesResponse"`

	GetLookupTablesResult *ArrayOfDMLookupTable `xml:"GetLookupTablesResult,omitempty"`
}

type CreateLookupTable struct {
	XMLName xml.Name `xml:"DMWebServices CreateLookupTable"`

	Token string `xml:"Token,omitempty"`

	CreativeID int32 `xml:"CreativeID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Shared bool `xml:"Shared,omitempty"`

	ColumnNames *ArrayOfString `xml:"ColumnNames,omitempty"`
}

type CreateLookupTableResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateLookupTableResponse"`

	CreateLookupTableResult int32 `xml:"CreateLookupTableResult,omitempty"`
}

type GetColumnMapping struct {
	XMLName xml.Name `xml:"DMWebServices GetColumnMapping"`

	Token string `xml:"Token,omitempty"`

	SourceColumns *ArrayOfDMSourceColumn `xml:"SourceColumns,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type GetColumnMappingResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetColumnMappingResponse"`

	GetColumnMappingResult *ArrayOfDMListField `xml:"GetColumnMappingResult,omitempty"`

	SourceColumns *ArrayOfDMSourceColumn `xml:"SourceColumns,omitempty"`

	DefaultKey int32 `xml:"DefaultKey,omitempty"`
}

type GetColumnMappingAll struct {
	XMLName xml.Name `xml:"DMWebServices GetColumnMappingAll"`

	Token string `xml:"Token,omitempty"`

	SourceColumns *ArrayOfDMSourceColumn `xml:"SourceColumns,omitempty"`
}

type GetColumnMappingAllResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetColumnMappingAllResponse"`

	GetColumnMappingAllResult *ArrayOfDMListField `xml:"GetColumnMappingAllResult,omitempty"`

	SourceColumns *ArrayOfDMSourceColumn `xml:"SourceColumns,omitempty"`

	DefaultKey int32 `xml:"DefaultKey,omitempty"`
}

type AddFieldsToList struct {
	XMLName xml.Name `xml:"DMWebServices AddFieldsToList"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	Fields *ArrayOfInt `xml:"Fields,omitempty"`
}

type AddFieldsToListResponse struct {
	XMLName xml.Name `xml:"DMWebServices AddFieldsToListResponse"`
}

type RemoveFieldsFromList struct {
	XMLName xml.Name `xml:"DMWebServices RemoveFieldsFromList"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	Fields *ArrayOfInt `xml:"Fields,omitempty"`
}

type RemoveFieldsFromListResponse struct {
	XMLName xml.Name `xml:"DMWebServices RemoveFieldsFromListResponse"`
}

type GetListInfo struct {
	XMLName xml.Name `xml:"DMWebServices GetListInfo"`

	Token string `xml:"Token,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`
}

type GetListInfoResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListInfoResponse"`

	GetListInfoResult *DMList `xml:"GetListInfoResult,omitempty"`
}

type GetPrimaryKeyCollection struct {
	XMLName xml.Name `xml:"DMWebServices GetPrimaryKeyCollection"`

	Token string `xml:"Token,omitempty"`
}

type GetPrimaryKeyCollectionResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetPrimaryKeyCollectionResponse"`

	GetPrimaryKeyCollectionResult *ArrayOfInt `xml:"GetPrimaryKeyCollectionResult,omitempty"`
}

type GetEventsList struct {
	XMLName xml.Name `xml:"DMWebServices GetEventsList"`

	Token string `xml:"Token,omitempty"`

	Creatives *ArrayOfInt `xml:"Creatives,omitempty"`
}

type GetEventsListResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetEventsListResponse"`

	GetEventsListResult *ArrayOfDMQEvent `xml:"GetEventsListResult,omitempty"`
}

type GetFieldSuggestions struct {
	XMLName xml.Name `xml:"DMWebServices GetFieldSuggestions"`

	Token string `xml:"Token,omitempty"`

	SourceColumn string `xml:"SourceColumn,omitempty"`

	SampleValues *ArrayOfString `xml:"SampleValues,omitempty"`
}

type GetFieldSuggestionsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetFieldSuggestionsResponse"`

	GetFieldSuggestionsResult *ArrayOfDMSuggestedFieldScore `xml:"GetFieldSuggestionsResult,omitempty"`

	SuggestedFieldType *DMFieldType `xml:"SuggestedFieldType,omitempty"`
}

type GetFields struct {
	XMLName xml.Name `xml:"DMWebServices GetFields"`

	Token string `xml:"Token,omitempty"`
}

type GetFieldsResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetFieldsResponse"`

	GetFieldsResult *ArrayOfDMField `xml:"GetFieldsResult,omitempty"`
}

type CreateField struct {
	XMLName xml.Name `xml:"DMWebServices CreateField"`

	Token string `xml:"Token,omitempty"`

	FieldType *DMFieldType `xml:"FieldType,omitempty"`

	ListField bool `xml:"ListField,omitempty"`

	FieldName string `xml:"FieldName,omitempty"`

	SourceColumn string `xml:"SourceColumn,omitempty"`

	UserAccess bool `xml:"UserAccess,omitempty"`

	StorageType *DMFieldStorageType `xml:"StorageType,omitempty"`

	UserOptOut bool `xml:"UserOptOut,omitempty"`
}

type CreateFieldResponse struct {
	XMLName xml.Name `xml:"DMWebServices CreateFieldResponse"`

	CreateFieldResult int32 `xml:"CreateFieldResult,omitempty"`
}

type RemoveField struct {
	XMLName xml.Name `xml:"DMWebServices RemoveField"`

	Token string `xml:"Token,omitempty"`

	FieldID int32 `xml:"FieldID,omitempty"`
}

type RemoveFieldResponse struct {
	XMLName xml.Name `xml:"DMWebServices RemoveFieldResponse"`
}

type GetListsWithSamePrimaryKeys struct {
	XMLName xml.Name `xml:"DMWebServices GetListsWithSamePrimaryKeys"`

	Token string `xml:"Token,omitempty"`

	ListIDs *ArrayOfInt `xml:"ListIDs,omitempty"`
}

type GetListsWithSamePrimaryKeysResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListsWithSamePrimaryKeysResponse"`

	GetListsWithSamePrimaryKeysResult *ArrayOfDMList `xml:"GetListsWithSamePrimaryKeysResult,omitempty"`

	Categories *ArrayOfDMCategory `xml:"Categories,omitempty"`
}

type GetListCategoriesWithSamePrimaryKeys struct {
	XMLName xml.Name `xml:"DMWebServices GetListCategoriesWithSamePrimaryKeys"`

	Token string `xml:"Token,omitempty"`

	ListIDs *ArrayOfInt `xml:"ListIDs,omitempty"`
}

type GetListCategoriesWithSamePrimaryKeysResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListCategoriesWithSamePrimaryKeysResponse"`

	GetListCategoriesWithSamePrimaryKeysResult *ArrayOfDMCategory `xml:"GetListCategoriesWithSamePrimaryKeysResult,omitempty"`
}

type GetListsFromCategoryWithSamePrimaryKeys struct {
	XMLName xml.Name `xml:"DMWebServices GetListsFromCategoryWithSamePrimaryKeys"`

	Token string `xml:"Token,omitempty"`

	ListIDs *ArrayOfInt `xml:"ListIDs,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`
}

type GetListsFromCategoryWithSamePrimaryKeysResponse struct {
	XMLName xml.Name `xml:"DMWebServices GetListsFromCategoryWithSamePrimaryKeysResponse"`

	GetListsFromCategoryWithSamePrimaryKeysResult *ArrayOfDMListDTO `xml:"GetListsFromCategoryWithSamePrimaryKeysResult,omitempty"`
}

type CancelListQuery struct {
	XMLName xml.Name `xml:"DMWebServices CancelListQuery"`

	Token string `xml:"Token,omitempty"`
}

type CancelListQueryResponse struct {
	XMLName xml.Name `xml:"DMWebServices CancelListQueryResponse"`
}

type UpdateFieldSettings struct {
	XMLName xml.Name `xml:"DMWebServices UpdateFieldSettings"`

	Token string `xml:"Token,omitempty"`

	Field *DMListField `xml:"Field,omitempty"`
}

type UpdateFieldSettingsResponse struct {
	XMLName xml.Name `xml:"DMWebServices UpdateFieldSettingsResponse"`
}

type SubscribeRecipientsByPK struct {
	XMLName xml.Name `xml:"DMWebServices SubscribeRecipientsByPK"`

	Token string `xml:"token,omitempty"`

	PrimaryKeyFieldId int32 `xml:"primaryKeyFieldId,omitempty"`

	PrimaryKeysValues *ArrayOfAnyType `xml:"primaryKeysValues,omitempty"`

	Resubscribe bool `xml:"resubscribe,omitempty"`

	CreateNewRecipient bool `xml:"createNewRecipient,omitempty"`

	ListId int32 `xml:"listId,omitempty"`

	DeploymentId int32 `xml:"deploymentId,omitempty"`

	AddExistingSubscribedRecipientsToDeployment bool `xml:"addExistingSubscribedRecipientsToDeployment,omitempty"`
}

type SubscribeRecipientsByPKResponse struct {
	XMLName xml.Name `xml:"DMWebServices SubscribeRecipientsByPKResponse"`

	SubscribeRecipientsByPKResult *RecipientSubscibingListDTO `xml:"SubscribeRecipientsByPKResult,omitempty"`
}

type DMListCriteria struct {
	// 	XMLName xml.Name `xml:"DMWebServices DMListCriteria"`

	IncludeLists *ArrayOfInt `xml:"IncludeLists,omitempty"`

	ExcludeLists *ArrayOfInt `xml:"ExcludeLists,omitempty"`

	EventCriteria *DMListEventCriteria `xml:"EventCriteria,omitempty"`

	FieldCriteria *ArrayOfDMFieldCriteria `xml:"FieldCriteria,omitempty"`
}

type ArrayOfInt struct {
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfInt"`

	Int []int32 `xml:"int,omitempty"`
}

type DMListEventCriteria struct {
	//	XMLName xml.Name `xml:"DMWebServices DMListEventCriteria"`

	Creatives *ArrayOfInt `xml:"Creatives,omitempty"`

	Users *ArrayOfInt `xml:"Users,omitempty"`

	DateRanges *ArrayOfDMRDateRange `xml:"DateRanges,omitempty"`

	Events *ArrayOfDMEventCriteria `xml:"Events,omitempty"`
}

type ArrayOfDMRDateRange struct {
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfDMRDateRange"`

	DMRDateRange []*DMRDateRange `xml:"DMRDateRange,omitempty"`
}

type DMRDateRange struct {
	XMLName xml.Name `xml:"DMWebServices DMRDateRange"`

	RangeType *DMRDateRangeType `xml:"RangeType,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`
}

type ArrayOfDMEventCriteria struct {
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfDMEventCriteria"`

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
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfDMFieldCriteria"`

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

type ArrayOfString struct {
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfString"`

	String []string `xml:"string,omitempty"`
}

type ArrayOfDouble struct {
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfDouble"`

	Double []float64 `xml:"double,omitempty"`
}

type ArrayOfBoolean struct {
	// XMLName xml.Name `xml:"DMWebServices ArrayOfBoolean"`

	Boolean []bool `xml:"boolean,omitempty"`
}

type ArrayOfDMListSegment struct {
	// XMLName xml.Name `xml:"DMWebServices ArrayOfDMListSegment"`

	DMListSegment []*DMListSegment `xml:"DMListSegment,omitempty"`
}

type DMListSegment struct {
	//	XMLName xml.Name `xml:"DMWebServices DMListSegment"`

	ListID int32 `xml:"ListID,omitempty"`

	Offset int32 `xml:"Offset,omitempty"`

	Count int32 `xml:"Count,omitempty"`
}

type ArrayOfDMRecipientRecord struct {
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfDMRecipientRecord"`

	DMRecipientRecord []*DMRecipientRecord `xml:"DMRecipientRecord,omitempty"`
}

type DMRecipientRecord struct {
	XMLName xml.Name `xml:"DMWebServices DMRecipientRecord"`

	RecipientID int32 `xml:"RecipientID,omitempty"`

	PrefersText bool `xml:"PrefersText,omitempty"`

	Unsubscribed bool `xml:"Unsubscribed,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	Fields *ArrayOfString `xml:"Fields,omitempty"`

	RSSOnly bool `xml:"RSSOnly,omitempty"`

	VerifiedFlash bool `xml:"VerifiedFlash,omitempty"`

	VerifiedHTML bool `xml:"VerifiedHTML,omitempty"`

	IP int32 `xml:"IP,omitempty"`

	OS string `xml:"OS,omitempty"`

	WB string `xml:"WB,omitempty"`

	LastEventID int32 `xml:"LastEventID,omitempty"`

	Health int32 `xml:"Health,omitempty"`
}

type ArrayOfDMListType struct {
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfDMListType"`
	// xml: name "ListTypes" in tag of listmanager.GetLists.ListTypes conflicts with name "ArrayOfDMListType" in *listmanager.ArrayOfDMListType.XMLName

	DMListType []*DMListType `xml:"DMListType,omitempty"`
}

type ArrayOfDMQList struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMQList"`

	DMQList []*DMQList `xml:"DMQList,omitempty"`
}

type DMQList struct {
	XMLName xml.Name `xml:"DMWebServices DMQList"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	PrimaryKey int32 `xml:"PrimaryKey,omitempty"`

	Fields *ArrayOfInt `xml:"Fields,omitempty"`

	ListType *DMListType `xml:"ListType,omitempty"`
}

type ArrayOfDMQField struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMQField"`

	DMQField []*DMQField `xml:"DMQField,omitempty"`
}

type DMQField struct {
	XMLName xml.Name `xml:"DMWebServices DMQField"`

	ID int32 `xml:"ID,omitempty"`

	FieldType *DMFieldType `xml:"FieldType,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type ArrayOfDMQListCategory struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMQListCategory"`

	DMQListCategory []*DMQListCategory `xml:"DMQListCategory,omitempty"`
}

type DMQListCategory struct {
	XMLName xml.Name `xml:"DMWebServices DMQListCategory"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Lists *ArrayOfInt `xml:"Lists,omitempty"`
}

type ArrayOfDMQCreative struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMQCreative"`

	DMQCreative []*DMQCreative `xml:"DMQCreative,omitempty"`
}

type DMQCreative struct {
	XMLName xml.Name `xml:"DMWebServices DMQCreative"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Events *ArrayOfInt `xml:"Events,omitempty"`
}

type ArrayOfDMQEvent struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMQEvent"`

	DMQEvent []*DMQEvent `xml:"DMQEvent,omitempty"`
}

type DMQEvent struct {
	XMLName xml.Name `xml:"DMWebServices DMQEvent"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type ArrayOfDMQUser struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMQUser"`

	DMQUser []*DMQUser `xml:"DMQUser,omitempty"`
}

type DMQUser struct {
	XMLName xml.Name `xml:"DMWebServices DMQUser"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`
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

type ArrayOfDMListPermission struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMListPermission"`

	DMListPermission []*DMListPermission `xml:"DMListPermission,omitempty"`
}

type ArrayOfDMListCategory struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMListCategory"`

	DMListCategory []*DMListCategory `xml:"DMListCategory,omitempty"`
}

type DMListCategory struct {
	XMLName xml.Name `xml:"DMWebServices DMListCategory"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Lists *ArrayOfDMList `xml:"Lists,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty"`
}

type ArrayOfDMList struct {
	//	XMLName xml.Name `xml:"DMWebServices ArrayOfDMList"`

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

type ArrayOfDMListField struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMListField"`

	DMListField []*DMListField `xml:"DMListField,omitempty"`
}

type DMListField struct {
	XMLName xml.Name `xml:"DMWebServices DMListField"`

	ID int32 `xml:"ID,omitempty"`

	FieldType *DMFieldType `xml:"FieldType,omitempty"`

	Name string `xml:"Name,omitempty"`

	SourceColumn string `xml:"SourceColumn,omitempty"`

	PrimaryKey bool `xml:"PrimaryKey,omitempty"`

	UserAccess bool `xml:"UserAccess,omitempty"`

	AllowDupes bool `xml:"AllowDupes,omitempty"`

	ListField bool `xml:"ListField,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	StorageType *DMFieldStorageType `xml:"StorageType,omitempty"`

	UserOptOut bool `xml:"UserOptOut,omitempty"`
}

type DMCursor struct {
	//	XMLName xml.Name `xml:"DMWebServices DMCursor"`

	High int32 `xml:"High,omitempty"`

	Low int32 `xml:"Low,omitempty"`

	High2 int32 `xml:"High2,omitempty"`

	Low2 int32 `xml:"Low2,omitempty"`
}

type ArrayOfDMFieldValue struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMFieldValue"`

	DMFieldValue []*DMFieldValue `xml:"DMFieldValue,omitempty"`
}

type DMFieldValue struct {
	XMLName xml.Name `xml:"DMWebServices DMFieldValue"`

	FieldID int32 `xml:"FieldID,omitempty"`

	Value string `xml:"Value,omitempty"`

	BValue []byte `xml:"BValue,omitempty"`
}

type ArrayOfDMAddRecipientRecord struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMAddRecipientRecord"`

	DMAddRecipientRecord []*DMAddRecipientRecord `xml:"DMAddRecipientRecord,omitempty"`
}

type DMAddRecipientRecord struct {
	XMLName xml.Name `xml:"DMWebServices DMAddRecipientRecord"`

	PrefersText bool `xml:"PrefersText,omitempty"`

	Unsubscribed bool `xml:"Unsubscribed,omitempty"`

	Values *ArrayOfDMFieldValue `xml:"Values,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	DeploymentID int32 `xml:"DeploymentID,omitempty"`
}

type DMImportStatus struct {
	XMLName xml.Name `xml:"DMWebServices DMImportStatus"`

	Inserted int32 `xml:"Inserted,omitempty"`

	Ignored int32 `xml:"Ignored,omitempty"`

	Updated int32 `xml:"Updated,omitempty"`

	Created int32 `xml:"Created,omitempty"`

	FieldsChanged int32 `xml:"FieldsChanged,omitempty"`

	FieldsAdded int32 `xml:"FieldsAdded,omitempty"`

	Added int32 `xml:"Added,omitempty"`

	TickCount int32 `xml:"TickCount,omitempty"`

	Counters *ArrayOfInt `xml:"Counters,omitempty"`

	LookupId int32 `xml:"LookupId,omitempty"`
}

type ArrayOfDMLookupTable struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMLookupTable"`

	DMLookupTable []*DMLookupTable `xml:"DMLookupTable,omitempty"`
}

type DMLookupTable struct {
	XMLName xml.Name `xml:"DMWebServices DMLookupTable"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Shared bool `xml:"Shared,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	Columns *ArrayOfDMLookupColumn `xml:"Columns,omitempty"`
}

type ArrayOfDMLookupColumn struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMLookupColumn"`

	DMLookupColumn []*DMLookupColumn `xml:"DMLookupColumn,omitempty"`
}

type DMLookupColumn struct {
	XMLName xml.Name `xml:"DMWebServices DMLookupColumn"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type ArrayOfDMSourceColumn struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMSourceColumn"`

	DMSourceColumn []*DMSourceColumn `xml:"DMSourceColumn,omitempty"`
}

type DMSourceColumn struct {
	XMLName xml.Name `xml:"DMWebServices DMSourceColumn"`

	Name string `xml:"Name,omitempty"`

	LooksLike *DMFieldType `xml:"LooksLike,omitempty"`

	LooksUnique bool `xml:"LooksUnique,omitempty"`

	LooksStatic bool `xml:"LooksStatic,omitempty"`

	LooksEmpty bool `xml:"LooksEmpty,omitempty"`

	FieldIndex int32 `xml:"FieldIndex,omitempty"`

	NewColumn string `xml:"NewColumn,omitempty"`
}

type ArrayOfDMSuggestedFieldScore struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMSuggestedFieldScore"`

	DMSuggestedFieldScore []*DMSuggestedFieldScore `xml:"DMSuggestedFieldScore,omitempty"`
}

type DMSuggestedFieldScore struct {
	XMLName xml.Name `xml:"DMWebServices DMSuggestedFieldScore"`

	FieldID int32 `xml:"FieldID,omitempty"`

	Score float64 `xml:"Score,omitempty"`
}

type ArrayOfDMField struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMField"`

	DMField []*DMField `xml:"DMField,omitempty"`
}

type DMField struct {
	XMLName xml.Name `xml:"DMWebServices DMField"`

	ID int32 `xml:"ID,omitempty"`

	FieldType *DMFieldType `xml:"FieldType,omitempty"`

	Name string `xml:"Name,omitempty"`

	SourceColumn string `xml:"SourceColumn,omitempty"`

	UserAccess bool `xml:"UserAccess,omitempty"`

	ListField bool `xml:"ListField,omitempty"`

	PrimaryKey bool `xml:"PrimaryKey,omitempty"`

	Created time.Time `xml:"Created,omitempty"`

	Modified time.Time `xml:"Modified,omitempty"`

	StorageType *DMFieldStorageType `xml:"StorageType,omitempty"`

	UserOptOut bool `xml:"UserOptOut,omitempty"`

	IsSeed bool `xml:"IsSeed,omitempty"`
}

type ArrayOfDMListDTO struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfDMListDTO"`

	DMListDTO []*DMListDTO `xml:"DMListDTO,omitempty"`
}

type DMListDTO struct {
	XMLName xml.Name `xml:"DMWebServices DMListDTO"`

	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type ArrayOfAnyType struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfAnyType"`

	AnyType struct {
	} `xml:"anyType,omitempty"`
}

type RecipientSubscibingListDTO struct {
	XMLName xml.Name `xml:"DMWebServices RecipientSubscibingListDTO"`

	Recipients *ArrayOfRecipientSubscibingDTO `xml:"Recipients,omitempty"`

	ErrorMessage string `xml:"ErrorMessage,omitempty"`
}

type ArrayOfRecipientSubscibingDTO struct {
	XMLName xml.Name `xml:"DMWebServices ArrayOfRecipientSubscibingDTO"`

	RecipientSubscibingDTO []*RecipientSubscibingDTO `xml:"RecipientSubscibingDTO,omitempty"`
}

type RecipientSubscibingDTO struct {
	XMLName xml.Name `xml:"DMWebServices RecipientSubscibingDTO"`

	PKValue struct {
	} `xml:"PKValue,omitempty"`

	RecipientId int32 `xml:"RecipientId,omitempty"`

	Status *RecipientSubscribingStatus `xml:"Status,omitempty"`
}

type DMListManagerSoap struct {
	client *SOAPClient
}

func NewDMListManagerSoap(url string, tls bool, auth *BasicAuth) *DMListManagerSoap {
	if url == "" {
		url = "https://uk56.em.sdlproducts.com/listmanager.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &DMListManagerSoap{
		client: client,
	}
}

func (service *DMListManagerSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

func (service *DMListManagerSoap) ExecuteListQuery(request *ExecuteListQuery) (*ExecuteListQueryResponse, error) {
	response := new(ExecuteListQueryResponse)
	err := service.client.Call("DMWebServices/ExecuteListQuery", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) ExecuteSuppressionListQuery(request *ExecuteSuppressionListQuery) (*ExecuteSuppressionListQueryResponse, error) {
	response := new(ExecuteSuppressionListQueryResponse)
	err := service.client.Call("DMWebServices/ExecuteSuppressionListQuery", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) SaveQueryResults(request *SaveQueryResults) (*SaveQueryResultsResponse, error) {
	response := new(SaveQueryResultsResponse)
	err := service.client.Call("DMWebServices/SaveQueryResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) SplitQueryResults(request *SplitQueryResults) (*SplitQueryResultsResponse, error) {
	response := new(SplitQueryResultsResponse)
	err := service.client.Call("DMWebServices/SplitQueryResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) ClearList(request *ClearList) (*ClearListResponse, error) {
	response := new(ClearListResponse)
	err := service.client.Call("DMWebServices/ClearList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UnsubscribeList(request *UnsubscribeList) (*UnsubscribeListResponse, error) {
	response := new(UnsubscribeListResponse)
	err := service.client.Call("DMWebServices/UnsubscribeList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetUnsubscribes(request *GetUnsubscribes) (*GetUnsubscribesResponse, error) {
	response := new(GetUnsubscribesResponse)
	err := service.client.Call("DMWebServices/GetUnsubscribes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) ResetHealth(request *ResetHealth) (*ResetHealthResponse, error) {
	response := new(ResetHealthResponse)
	err := service.client.Call("DMWebServices/ResetHealth", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetEmailAddressesByHealth(request *GetEmailAddressesByHealth) (*GetEmailAddressesByHealthResponse, error) {
	response := new(GetEmailAddressesByHealthResponse)
	err := service.client.Call("DMWebServices/GetEmailAddressesByHealth", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetQueryResults(request *GetQueryResults) (*GetQueryResultsResponse, error) {
	response := new(GetQueryResultsResponse)
	err := service.client.Call("DMWebServices/GetQueryResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetSuppressionQueryResults(request *GetSuppressionQueryResults) (*GetSuppressionQueryResultsResponse, error) {
	response := new(GetSuppressionQueryResultsResponse)
	err := service.client.Call("DMWebServices/GetSuppressionQueryResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetQueryResultsCSV(request *GetQueryResultsCSV) (*GetQueryResultsCSVResponse, error) {
	response := new(GetQueryResultsCSVResponse)
	err := service.client.Call("DMWebServices/GetQueryResultsCSV", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetSuppressionQueryResultsCSV(request *GetSuppressionQueryResultsCSV) (*GetSuppressionQueryResultsCSVResponse, error) {
	response := new(GetSuppressionQueryResultsCSVResponse)
	err := service.client.Call("DMWebServices/GetSuppressionQueryResultsCSV", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListsAndFields(request *GetListsAndFields) (*GetListsAndFieldsResponse, error) {
	response := new(GetListsAndFieldsResponse)
	err := service.client.Call("DMWebServices/GetListsAndFields", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetCreativesAndEvents(request *GetCreativesAndEvents) (*GetCreativesAndEventsResponse, error) {
	response := new(GetCreativesAndEventsResponse)
	err := service.client.Call("DMWebServices/GetCreativesAndEvents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetRecipientByID(request *GetRecipientByID) (*GetRecipientByIDResponse, error) {
	response := new(GetRecipientByIDResponse)
	err := service.client.Call("DMWebServices/GetRecipientByID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetRecipientByPK(request *GetRecipientByPK) (*GetRecipientByPKResponse, error) {
	response := new(GetRecipientByPKResponse)
	err := service.client.Call("DMWebServices/GetRecipientByPK", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetRecipientIdByPK(request *GetRecipientIdByPK) (*GetRecipientIdByPKResponse, error) {
	response := new(GetRecipientIdByPKResponse)
	err := service.client.Call("DMWebServices/GetRecipientIdByPK", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetVisibleUsers(request *GetVisibleUsers) (*GetVisibleUsersResponse, error) {
	response := new(GetVisibleUsersResponse)
	err := service.client.Call("DMWebServices/GetVisibleUsers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CreateListCategory(request *CreateListCategory) (*CreateListCategoryResponse, error) {
	response := new(CreateListCategoryResponse)
	err := service.client.Call("DMWebServices/CreateListCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RenameListCategory(request *RenameListCategory) (*RenameListCategoryResponse, error) {
	response := new(RenameListCategoryResponse)
	err := service.client.Call("DMWebServices/RenameListCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) DeleteListCategory(request *DeleteListCategory) (*DeleteListCategoryResponse, error) {
	response := new(DeleteListCategoryResponse)
	err := service.client.Call("DMWebServices/DeleteListCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListCategories(request *GetListCategories) (*GetListCategoriesResponse, error) {
	response := new(GetListCategoriesResponse)
	err := service.client.Call("DMWebServices/GetListCategories", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListCategoriesEx(request *GetListCategoriesEx) (*GetListCategoriesExResponse, error) {
	response := new(GetListCategoriesExResponse)
	err := service.client.Call("DMWebServices/GetListCategoriesEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetLists(request *GetLists) (*GetListsResponse, error) {
	response := new(GetListsResponse)
	err := service.client.Call("DMWebServices/GetLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListFields(request *GetListFields) (*GetListFieldsResponse, error) {
	response := new(GetListFieldsResponse)
	err := service.client.Call("DMWebServices/GetListFields", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListRecords(request *GetListRecords) (*GetListRecordsResponse, error) {
	response := new(GetListRecordsResponse)
	err := service.client.Call("DMWebServices/GetListRecords", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) AddRecipientRecord(request *AddRecipientRecord) (*AddRecipientRecordResponse, error) {
	response := new(AddRecipientRecordResponse)
	err := service.client.Call("DMWebServices/AddRecipientRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) AddRecipientRecords(request *AddRecipientRecords) (*AddRecipientRecordsResponse, error) {
	response := new(AddRecipientRecordsResponse)
	err := service.client.Call("DMWebServices/AddRecipientRecords", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) AddSuppressionRecord(request *AddSuppressionRecord) (*AddSuppressionRecordResponse, error) {
	response := new(AddSuppressionRecordResponse)
	err := service.client.Call("DMWebServices/AddSuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) AddDomainSuppressionRecord(request *AddDomainSuppressionRecord) (*AddDomainSuppressionRecordResponse, error) {
	response := new(AddDomainSuppressionRecordResponse)
	err := service.client.Call("DMWebServices/AddDomainSuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) AddMD5SuppressionRecord(request *AddMD5SuppressionRecord) (*AddMD5SuppressionRecordResponse, error) {
	response := new(AddMD5SuppressionRecordResponse)
	err := service.client.Call("DMWebServices/AddMD5SuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RemoveRecipient(request *RemoveRecipient) (*RemoveRecipientResponse, error) {
	response := new(RemoveRecipientResponse)
	err := service.client.Call("DMWebServices/RemoveRecipient", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RemoveQueryResultsFromLists(request *RemoveQueryResultsFromLists) (*RemoveQueryResultsFromListsResponse, error) {
	response := new(RemoveQueryResultsFromListsResponse)
	err := service.client.Call("DMWebServices/RemoveQueryResultsFromLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RemoveSuppressionRecord(request *RemoveSuppressionRecord) (*RemoveSuppressionRecordResponse, error) {
	response := new(RemoveSuppressionRecordResponse)
	err := service.client.Call("DMWebServices/RemoveSuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RemoveDomainSuppressionRecord(request *RemoveDomainSuppressionRecord) (*RemoveDomainSuppressionRecordResponse, error) {
	response := new(RemoveDomainSuppressionRecordResponse)
	err := service.client.Call("DMWebServices/RemoveDomainSuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RemoveMD5SuppressionRecord(request *RemoveMD5SuppressionRecord) (*RemoveMD5SuppressionRecordResponse, error) {
	response := new(RemoveMD5SuppressionRecordResponse)
	err := service.client.Call("DMWebServices/RemoveMD5SuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UpdateSuppressionRecord(request *UpdateSuppressionRecord) (*UpdateSuppressionRecordResponse, error) {
	response := new(UpdateSuppressionRecordResponse)
	err := service.client.Call("DMWebServices/UpdateSuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UpdateDomainSuppressionRecord(request *UpdateDomainSuppressionRecord) (*UpdateDomainSuppressionRecordResponse, error) {
	response := new(UpdateDomainSuppressionRecordResponse)
	err := service.client.Call("DMWebServices/UpdateDomainSuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UpdateMD5SuppressionRecord(request *UpdateMD5SuppressionRecord) (*UpdateMD5SuppressionRecordResponse, error) {
	response := new(UpdateMD5SuppressionRecordResponse)
	err := service.client.Call("DMWebServices/UpdateMD5SuppressionRecord", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UpdateRecipient(request *UpdateRecipient) (*UpdateRecipientResponse, error) {
	response := new(UpdateRecipientResponse)
	err := service.client.Call("DMWebServices/UpdateRecipient", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) AddRecipientToList(request *AddRecipientToList) (*AddRecipientToListResponse, error) {
	response := new(AddRecipientToListResponse)
	err := service.client.Call("DMWebServices/AddRecipientToList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UpdateField(request *UpdateField) (*UpdateFieldResponse, error) {
	response := new(UpdateFieldResponse)
	err := service.client.Call("DMWebServices/UpdateField", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CreateList(request *CreateList) (*CreateListResponse, error) {
	response := new(CreateListResponse)
	err := service.client.Call("DMWebServices/CreateList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CreateRecipientList(request *CreateRecipientList) (*CreateRecipientListResponse, error) {
	response := new(CreateRecipientListResponse)
	err := service.client.Call("DMWebServices/CreateRecipientList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CreateSuppressionList(request *CreateSuppressionList) (*CreateSuppressionListResponse, error) {
	response := new(CreateSuppressionListResponse)
	err := service.client.Call("DMWebServices/CreateSuppressionList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CreateDomainSuppressionList(request *CreateDomainSuppressionList) (*CreateDomainSuppressionListResponse, error) {
	response := new(CreateDomainSuppressionListResponse)
	err := service.client.Call("DMWebServices/CreateDomainSuppressionList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CreateMD5SuppressionList(request *CreateMD5SuppressionList) (*CreateMD5SuppressionListResponse, error) {
	response := new(CreateMD5SuppressionListResponse)
	err := service.client.Call("DMWebServices/CreateMD5SuppressionList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UpdateList(request *UpdateList) (*UpdateListResponse, error) {
	response := new(UpdateListResponse)
	err := service.client.Call("DMWebServices/UpdateList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) DeleteList(request *DeleteList) (*DeleteListResponse, error) {
	response := new(DeleteListResponse)
	err := service.client.Call("DMWebServices/DeleteList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RefreshSmartLists(request *RefreshSmartLists) (*RefreshSmartListsResponse, error) {
	response := new(RefreshSmartListsResponse)
	err := service.client.Call("DMWebServices/RefreshSmartLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UploadData(request *UploadData) (*UploadDataResponse, error) {
	response := new(UploadDataResponse)
	err := service.client.Call("DMWebServices/UploadData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UploadDataA(request *UploadDataA) (*UploadDataAResponse, error) {
	response := new(UploadDataAResponse)
	err := service.client.Call("DMWebServices/UploadDataA", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UploadDataCSV(request *UploadDataCSV) (*UploadDataCSVResponse, error) {
	response := new(UploadDataCSVResponse)
	err := service.client.Call("DMWebServices/UploadDataCSV", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) InitializeListImport(request *InitializeListImport) (*InitializeListImportResponse, error) {
	response := new(InitializeListImportResponse)
	err := service.client.Call("DMWebServices/InitializeListImport", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) InsertListData(request *InsertListData) (*InsertListDataResponse, error) {
	response := new(InsertListDataResponse)
	err := service.client.Call("DMWebServices/InsertListData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) FinalizeListImport(request *FinalizeListImport) (*FinalizeListImportResponse, error) {
	response := new(FinalizeListImportResponse)
	err := service.client.Call("DMWebServices/FinalizeListImport", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) InsertTableData(request *InsertTableData) (*InsertTableDataResponse, error) {
	response := new(InsertTableDataResponse)
	err := service.client.Call("DMWebServices/InsertTableData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetLookupTables(request *GetLookupTables) (*GetLookupTablesResponse, error) {
	response := new(GetLookupTablesResponse)
	err := service.client.Call("DMWebServices/GetLookupTables", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CreateLookupTable(request *CreateLookupTable) (*CreateLookupTableResponse, error) {
	response := new(CreateLookupTableResponse)
	err := service.client.Call("DMWebServices/CreateLookupTable", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetColumnMapping(request *GetColumnMapping) (*GetColumnMappingResponse, error) {
	response := new(GetColumnMappingResponse)
	err := service.client.Call("DMWebServices/GetColumnMapping", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetColumnMappingAll(request *GetColumnMappingAll) (*GetColumnMappingAllResponse, error) {
	response := new(GetColumnMappingAllResponse)
	err := service.client.Call("DMWebServices/GetColumnMappingAll", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) AddFieldsToList(request *AddFieldsToList) (*AddFieldsToListResponse, error) {
	response := new(AddFieldsToListResponse)
	err := service.client.Call("DMWebServices/AddFieldsToList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RemoveFieldsFromList(request *RemoveFieldsFromList) (*RemoveFieldsFromListResponse, error) {
	response := new(RemoveFieldsFromListResponse)
	err := service.client.Call("DMWebServices/RemoveFieldsFromList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListInfo(request *GetListInfo) (*GetListInfoResponse, error) {
	response := new(GetListInfoResponse)
	err := service.client.Call("DMWebServices/GetListInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetPrimaryKeyCollection(request *GetPrimaryKeyCollection) (*GetPrimaryKeyCollectionResponse, error) {
	response := new(GetPrimaryKeyCollectionResponse)
	err := service.client.Call("DMWebServices/GetPrimaryKeyCollection", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetEventsList(request *GetEventsList) (*GetEventsListResponse, error) {
	response := new(GetEventsListResponse)
	err := service.client.Call("DMWebServices/GetEventsList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetFieldSuggestions(request *GetFieldSuggestions) (*GetFieldSuggestionsResponse, error) {
	response := new(GetFieldSuggestionsResponse)
	err := service.client.Call("DMWebServices/GetFieldSuggestions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetFields(request *GetFields) (*GetFieldsResponse, error) {
	response := new(GetFieldsResponse)
	err := service.client.Call("DMWebServices/GetFields", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CreateField(request *CreateField) (*CreateFieldResponse, error) {
	response := new(CreateFieldResponse)
	err := service.client.Call("DMWebServices/CreateField", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) RemoveField(request *RemoveField) (*RemoveFieldResponse, error) {
	response := new(RemoveFieldResponse)
	err := service.client.Call("DMWebServices/RemoveField", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListsWithSamePrimaryKeys(request *GetListsWithSamePrimaryKeys) (*GetListsWithSamePrimaryKeysResponse, error) {
	response := new(GetListsWithSamePrimaryKeysResponse)
	err := service.client.Call("DMWebServices/GetListsWithSamePrimaryKeys", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListCategoriesWithSamePrimaryKeys(request *GetListCategoriesWithSamePrimaryKeys) (*GetListCategoriesWithSamePrimaryKeysResponse, error) {
	response := new(GetListCategoriesWithSamePrimaryKeysResponse)
	err := service.client.Call("DMWebServices/GetListCategoriesWithSamePrimaryKeys", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) GetListsFromCategoryWithSamePrimaryKeys(request *GetListsFromCategoryWithSamePrimaryKeys) (*GetListsFromCategoryWithSamePrimaryKeysResponse, error) {
	response := new(GetListsFromCategoryWithSamePrimaryKeysResponse)
	err := service.client.Call("DMWebServices/GetListsFromCategoryWithSamePrimaryKeys", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) CancelListQuery(request *CancelListQuery) (*CancelListQueryResponse, error) {
	response := new(CancelListQueryResponse)
	err := service.client.Call("DMWebServices/CancelListQuery", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) UpdateFieldSettings(request *UpdateFieldSettings) (*UpdateFieldSettingsResponse, error) {
	response := new(UpdateFieldSettingsResponse)
	err := service.client.Call("DMWebServices/UpdateFieldSettings", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *DMListManagerSoap) SubscribeRecipientsByPK(request *SubscribeRecipientsByPK) (*SubscribeRecipientsByPKResponse, error) {
	response := new(SubscribeRecipientsByPKResponse)
	err := service.client.Call("DMWebServices/SubscribeRecipientsByPK", request, response)
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
