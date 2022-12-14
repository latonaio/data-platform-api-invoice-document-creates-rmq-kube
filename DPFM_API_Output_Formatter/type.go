package dpfm_api_output_formatter

type InvoiceDocumentCreates struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	InvoiceDocument string `json:"invoice_document_code"`
	Deleted         bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type CreatesMessage struct {
	HeaderCreates []HeaderCreates `json:"Header"`
	HeaderPartner []HeaderPartner `json:"HeaderPartner"`
	Item          []Item          `json:"Item"`
}

type UpdatesMessage struct {
	HeaderUpdates *HeaderUpdates `json:"Header"`
}

type HeaderCreates struct {
	InvoiceDocument                   *int     `json:"InvoiceDocument"`
	CreationDate                      *string  `json:"CreationDate"`
	LastChangeDate                    *string  `json:"LastChangeDate"`
	BillFromParty                     *int     `json:"BillFromParty"`
	BillToParty                       *int     `json:"BillToParty"`
	BillFromCountry                   *string  `json:"BillFromCountry"`
	BillToCountry                     *string  `json:"BillToCountry"`
	Payer                             *int     `json:"Payer"`
	Payee                             *int     `json:"Payee"`
	InvoiceDocumentDate               *string  `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime               *string  `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate            *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              *string  `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string  `json:"AccountingPostingDate"`
	InvoiceDocumentIsCancelled        *bool    `json:"InvoiceDocumentIsCancelled"`
	CancelledInvoiceDocument          *int     `json:"CancelledInvoiceDocument"`
	IsExportImportDelivery            *bool    `json:"IsExportImportDelivery"`
	HeaderBillingIsConfirmed          *bool    `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus           *string  `json:"HeaderBillingConfStatus"`
	TotalNetAmount                    *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                    *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                  *float32 `json:"TotalGrossAmount"`
	TransactionCurrency               *string  `json:"TransactionCurrency"`
	Incoterms                         *string  `json:"Incoterms"`
	PaymentTerms                      *string  `json:"PaymentTerms"`
	DueCalculationBaseDate            *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                    *string  `json:"PaymentDueDate"`
	NetPaymentDays                    *int     `json:"NetPaymentDays"`
	PaymentMethod                     *string  `json:"PaymentMethod"`
	HeaderPaymentBlockStatus          *bool    `json:"HeaderPaymentBlockStatus"`
	ExternalReferenceDocument         *string  `json:"ExternalReferenceDocument"`
	DocumentHeaderText                *string  `json:"DocumentHeaderText"`
	HeaderPaymentRequisitionIsCreated *bool    `json:"HeaderPaymentRequisitionIsCreated"`
}

type HeaderUpdates struct {
}

type HeaderPartner struct {
	InvoiceDocument         *int    `json:"InvoiceDocument"`
	PartnerFunction         *string `json:"PartnerFunction"`
	BusinessPartner         *int    `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type HeaderPartnerContact struct {
	InvoiceDocument   *int    `json:"InvoiceDocument"`
	PartnerFunction   *string `json:"PartnerFunction"`
	BusinessPartner   *int    `json:"BusinessPartner"`
	ContactID         *int    `json:"ContactID"`
	ContactPersonName *string `json:"ContactPersonName"`
	EmailAddress      *string `json:"EmailAddress"`
	PhoneNumber       *string `json:"PhoneNumber"`
	MobilePhoneNumber *string `json:"MobilePhoneNumber"`
	FaxNumber         *string `json:"FaxNumber"`
	ContactTag1       *string `json:"ContactTag1"`
	ContactTag2       *string `json:"ContactTag2"`
	ContactTag3       *string `json:"ContactTag3"`
	ContactTag4       *string `json:"ContactTag4"`
}

type HeaderPDF struct {
	InvoiceDocument          *int    `json:"InvoiceDocument"`
	DocType                  *string `json:"DocType"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FileName                 *string `json:"FileName"`
}

type Address struct {
	InvoiceDocument *int    `json:"InvoiceDocument"`
	AddressID       *int    `json:"AddressID"`
	PostalCode      *string `json:"PostalCode"`
	LocalRegion     *string `json:"LocalRegion"`
	Country         *string `json:"Country"`
	District        *string `json:"District"`
	StreetName      *string `json:"StreetName"`
	CityName        *string `json:"CityName"`
	Building        *string `json:"Building"`
	Floor           *int    `json:"Floor"`
	Room            *int    `json:"Room"`
}

type Item struct {
	InvoiceDocument                 int      `json:"InvoiceDocument"`
	InvoiceDocumentItem             int      `json:"InvoiceDocumentItem"`
	InvoiceDocumentItemCategory     *string  `json:"InvoiceDocumentItemCategory"`
	InvoiceDocumentItemText         *string  `json:"InvoiceDocumentItemText"`
	InvoiceDocumentItemTextByBuyer  string   `json:"InvoiceDocumentItemTextByBuyer"`
	InvoiceDocumentItemTextBySeller string   `json:"InvoiceDocumentItemTextBySeller"`
	CreationDate                    *string  `json:"CreationDate"`
	CreationTime                    *string  `json:"CreationTime"`
	ItemBillingIsConfirmed          *bool    `json:"ItemBillingIsConfirmed"`
	ItemBillingConfStatus           *string  `json:"ItemBillingConfStatus"`
	Buyer                           *int     `json:"Buyer"`
	Seller                          *int     `json:"Seller"`
	DeliverFromParty                *int     `json:"DeliverFromParty"`
	DeliverToParty                  *int     `json:"DeliverToParty"`
	ProductStandardID               *string  `json:"ProductStandardID"`
	Batch                           *string  `json:"Batch"`
	Product                         *string  `json:"Product"`
	ProductGroup                    *string  `json:"ProductGroup"`
	ProductionPlantPartnerFunction  *string  `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner  *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                 *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation  *string  `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction     *string  `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner     *int     `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                    *string  `json:"IssuingPlant"`
	IssuingPlantStorageLocation     *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction   *string  `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner   *int     `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                  *string  `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation   *string  `json:"ReceivingPlantStorageLocation"`
	ServicesRenderedDate            *string  `json:"ServicesRenderedDate"`
	InvoiceQuantity                 *float32 `json:"InvoiceQuantity"`
	InvoiceQuantityUnit             *string  `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit       *float32 `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                        *string  `json:"BaseUnit"`
	ActualGoodsIssueDate            *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime            *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate          *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime          *string  `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight                 *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                   *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                  *string  `json:"ItemWeightUnit"`
	NetAmount                       *float32 `json:"NetAmount"`
	TaxAmount                       *float32 `json:"TaxAmount"`
	GrossAmount                     *float32 `json:"GrossAmount"`
	ItemPaymentBlockStatus          *bool    `json:"ItemPaymentBlockStatus"`
	GoodsIssueOrReceiptSlipNumber   *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency             *string  `json:"TransactionCurrency"`
	PricingDate                     *string  `json:"PricingDate"`
	ProductTaxClassification        *string  `json:"ProductTaxClassification"`
	Project                         *string  `json:"Project"`
	OrderID                         *int     `json:"OrderID"`
	OrderItem                       *int     `json:"OrderItem"`
	OrderType                       *string  `json:"OrderType"`
	ContractType                    *string  `json:"ContractType"`
	OrderVaridityStartDate          *string  `json:"OrderVaridityStartDate"`
	OrderVaridityEndDate            *string  `json:"OrderVaridityEndDate"`
	InvoiceScheduleStartDate        *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate          *string  `json:"InvoiceScheduleEndDate"`
	DeliveryDocument                *int     `json:"DeliveryDocument"`
	DeliveryDocumentItem            *int     `json:"DeliveryDocumentItem"`
	OriginDocument                  *int     `json:"OriginDocument"`
	OriginDocumentItem              *int     `json:"OriginDocumentItem"`
	ReferenceDocument               *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem           *int     `json:"ReferenceDocumentItem"`
	ReferenceDocumentType           *string  `json:"ReferenceDocumentType"`
	ExternalReferenceDocument       *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem   *string  `json:"ExternalReferenceDocumentItem"`
	TaxCode                         *string  `json:"TaxCode"`
	TaxRate                         *float32 `json:"TaxRate"`
	CountryOfOrigin                 *string  `json:"CountryOfOrigin"`
	ItemPaymentRequisitionIsCreated *bool    `json:"ItemPaymentRequisitionIsCreated"`
}
type ItemPartner struct {
	InvoiceDocument     *int    `json:"InvoiceDocument"`
	InvoiceDocumentItem *int    `json:"InvoiceDocumentItem"`
	PartnerFunction     *string `json:"PartnerFunction"`
	BusinessPartner     *int    `json:"BusinessPartner"`
}

type ItemPricingElement struct {
	InvoiceDocument            *int     `json:"InvoiceDocument"`
	InvoiceDocumentItem        *int     `json:"InvoiceDocumentItem"`
	PricingProcedureCounter    *int     `json:"PricingProcedureCounter"`
	PricingProcedureStep       *int     `json:"PricingProcedureStep"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
