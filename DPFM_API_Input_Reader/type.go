package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string          `json:"connection_key"`
	Result           bool            `json:"result"`
	RedisKey         string          `json:"redis_key"`
	Filepath         string          `json:"filepath"`
	APIStatusCode    int             `json:"api_status_code"`
	RuntimeSessionID string          `json:"runtime_session_id"`
	BusinessPartner  *int            `json:"business_partner"`
	ServiceLabel     string          `json:"service_label"`
	APIType          string          `json:"api_type"`
	InputParameters  InputParameters `json:"InputParameters"`
	Header           Header          `json:"InvoiceDocument"`
	APISchema        string          `json:"api_schema"`
	Accepter         []string        `json:"accepter"`
	Deleted          bool            `json:"deleted"`
}

type InputParameters struct {
	InvoiceDocumentDate       *string    `json:"InvoiceDocumentDate"`
	BillFromParty             *[]*int    `json:"BillFromParty"`
	BillFromPartyTo           *int       `json:"BillFromPartyTo"`
	BillFromPartyFrom         *int       `json:"BillFromPartyFrom"`
	BillToParty               *[]*int    `json:"BillToParty"`
	BillToPartyTo             *int       `json:"BillToPartyTo"`
	BillToPartyFrom           *int       `json:"BillToPartyFrom"`
	ConfirmedDeliveryDate     *[]*string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateTo   *string    `json:"ConfirmedDeliveryDateTo"`
	ConfirmedDeliveryDateFrom *string    `json:"ConfirmedDeliveryDateFrom"`
	ActualGoodsIssueDate      *[]*string `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueDateTo    *string    `json:"ActualGoodsIssueDateTo"`
	ActualGoodsIssueDateFrom  *string    `json:"ActualGoodsIssueDateFrom"`
	ActualGoodsIssueTime      *[]*string `json:"ActualGoodsIssueTime"`
	ActualGoodsIssueTimeTo    *string    `json:"ActualGoodsIssueTimeTo"`
	ActualGoodsIssueTimeFrom  *string    `json:"ActualGoodsIssueTimeFrom"`
	ReferenceDocument         *int       `json:"ReferenceDocument"`
}

type Header struct {
	InvoiceDocument                   int      `json:"InvoiceDocument"`
	SupplyChainRelationshipID         *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID  *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID  *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer							  *int	   `json:"Buyer"`
	Seller							  *int	   `json:"Seller"`
	BillToParty                       *int     `json:"BillToParty"`
	BillFromParty                     *int     `json:"BillFromParty"`
	BillToCountry                     *string  `json:"BillToCountry"`
	BillFromCountry                   *string  `json:"BillFromCountry"`
	Payer                             *int     `json:"Payer"`
	Payee                             *int     `json:"Payee"`
	InvoiceDocumentDate               *string  `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime               *string  `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate            *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              *string  `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string  `json:"AccountingPostingDate"`
	IsExportImport                    *bool    `json:"IsExportImport"`
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
	DocumentHeaderText                *string  `json:"DocumentHeaderText"`
	HeaderIsCleared                   *bool    `json:"HeaderIsCleared"`
	HeaderPaymentBlockStatus          *bool    `json:"HeaderPaymentBlockStatus"`
	HeaderPaymentRequisitionIsCreated *bool    `json:"HeaderPaymentRequisitionIsCreated"`
	ExternalReferenceDocument         *string  `json:"ExternalReferenceDocument"`
	CreationDate                      *string  `json:"CreationDate"`
	CreationTime                      *string  `json:"CreationTime"`
	LastChangeDate                    *string  `json:"LastChangeDate"`
	LastChangeTime                    *string  `json:"LastChangeTime"`
	IsCancelled                       *bool    `json:"IsCancelled"`
	Item                              []Item      `json:"Item"`
	HeaderDoc                         []HeaderDoc `json:"HeaderDoc"`
	Partner                           []Partner   `json:"Partner"`
	Address                           []Address   `json:"Address"`
}

type Item struct {
	InvoiceDocument                         int      `json:"InvoiceDocument"`
	InvoiceDocumentItem                     int      `json:"InvoiceDocumentItem"`
	InvoiceDocumentItemCategory             *string  `json:"InvoiceDocumentItemCategory"`
	SupplyChainRelationshipID               *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID       *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID  *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	InvoiceDocumentItemText                 *string  `json:"InvoiceDocumentItemText"`
	InvoiceDocumentItemTextByBuyer          *string  `json:"InvoiceDocumentItemTextByBuyer"`
	InvoiceDocumentItemTextBySeller         *string  `json:"InvoiceDocumentItemTextBySeller"`
	Product                                 *string  `json:"Product"`
	ProductGroup                            *string  `json:"ProductGroup"`
	ProductStandardID                       *string  `json:"ProductStandardID"`
	ItemBillingIsConfirmed                  *bool    `json:"ItemBillingIsConfirmed"`
	Buyer                                   *int     `json:"Buyer"`
	Seller                                  *int     `json:"Seller"`
	DeliverToParty                          *int     `json:"DeliverToParty"`
	DeliverFromParty                        *int     `json:"DeliverFromParty"`
	DeliverToPlant                          *string  `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation           *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlant                        *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation         *string  `json:"DeliverFromPlantStorageLocation"`
	ProductionPlantBusinessPartner          *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                         *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation          *string  `json:"ProductionPlantStorageLocation"`
	ServicesRenderedDate                    *string  `json:"ServicesRenderedDate"`
	InvoiceQuantity                         *float32 `json:"InvoiceQuantity"`
	InvoiceQuantityUnit                     *string  `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit               *float32 `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                                *string  `json:"BaseUnit"`
	DeliveryUnit                            *string  `json:"DeliveryUnit"`
	ActualGoodsIssueDate                    *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                    *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                  *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                  *string  `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight                         *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                           *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                          *string  `json:"ItemWeightUnit"`
	NetAmount                               *float32 `json:"NetAmount"`
	TaxAmount                               *float32 `json:"TaxAmount"`
	GrossAmount                             *float32 `json:"GrossAmount"`
	GoodsIssueOrReceiptSlipNumber           *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency                     *string  `json:"TransactionCurrency"`
	PricingDate                             *string  `json:"PricingDate"`
	TransactionTaxClassification            *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                *string  `json:"DefinedTaxClassification"`
	Project                                 *int     `json:"Project"`
	WBSElement                              *int     `json:"WBSElement"`
	OrderID                                 *int     `json:"OrderID"`
	OrderItem                               *int     `json:"OrderItem"`
	OrderType                               *string  `json:"OrderType"`
	ContractType                            *string  `json:"ContractType"`
	OrderVaridityStartDate                  *string  `json:"OrderVaridityStartDate"`
	OrderVaridityEndDate                    *string  `json:"OrderVaridityEndDate"`
	InvoicePeriodStartDate                  *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                    *string  `json:"InvoicePeriodEndDate"`
	DeliveryDocument                        *int     `json:"DeliveryDocument"`
	DeliveryDocumentItem                    *int     `json:"DeliveryDocumentItem"`
	OriginDocument                          *int     `json:"OriginDocument"`
	OriginDocumentItem                      *int     `json:"OriginDocumentItem"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
	CountryOfOrigin                         *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                 *string  `json:"CountryOfOriginLanguage"`
	Equipment                        		*int     `json:"Equipment"`
	ItemPaymentRequisitionIsCreated         *bool    `json:"ItemPaymentRequisitionIsCreated"`
	ItemIsCleared                           *bool    `json:"ItemIsCleared"`
	ItemPaymentBlockStatus                  *bool    `json:"ItemPaymentBlockStatus"`
	ExternalReferenceDocument               *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem           *string  `json:"ExternalReferenceDocumentItem"`
	CreationDate                            *string  `json:"CreationDate"`
	CreationTime                            *string  `json:"CreationTime"`
	LastChangeDate                          *string  `json:"LastChangeDate"`
	LastChangeTime                          *string  `json:"LastChangeTime"`
	IsCancelled                             *bool    `json:"IsCancelled"`
	ItemPricingElement                      []ItemPricingElement `json:"ItemPricingElement"`
}

type ItemPricingElement struct {
	InvoiceDocument            int      `json:"InvoiceDocument"`
	InvoiceDocumentItem        int      `json:"InvoiceDocumentItem"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	SupplyChainRelationshipID  int		`json:"SupplyChainRelationshipID"`
	Buyer					   int		`json:"Buyer"`
	Seller					   int		`json:"Seller"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     *int  	`json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     *int  	`json:"ConditionScaleQuantity"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
	CreationDate               *string  `json:"CreationDate"`
	CreationTime               *string  `json:"CreationTime"`
	LastChangeDate             *string  `json:"LastChangeDate"`
	LastChangeTime             *string  `json:"LastChangeTime"`
	IsCancelled                *bool    `json:"IsCancelled"`
}

type HeaderDoc struct {
	InvoiceDocument          int     `json:"InvoiceDocument"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type Partner struct {
	InvoiceDocument         int     `json:"InvoiceDocument"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type Address struct {
	InvoiceDocument int     `json:"InvoiceDocument"`
	AddressID       int     `json:"AddressID"`
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
