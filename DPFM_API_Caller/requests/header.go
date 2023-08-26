package requests

type Header struct {
	InvoiceDocument                   int      `json:"InvoiceDocument"`
	SupplyChainRelationshipID         int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID  int      `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID  int      `json:"SupplyChainRelationshipPaymentID"`
	Buyer							  int	   `json:"Buyer"`
	Seller							  int	   `json:"Seller"`
	BillToParty                       int      `json:"BillToParty"`
	BillFromParty                     int      `json:"BillFromParty"`
	BillToCountry                     string   `json:"BillToCountry"`
	BillFromCountry                   string   `json:"BillFromCountry"`
	Payer                             int      `json:"Payer"`
	Payee                             int      `json:"Payee"`
	InvoiceDocumentDate               string   `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime               string   `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate            string   `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              string   `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string  `json:"AccountingPostingDate"`
	IsExportImport                    *bool    `json:"IsExportImport"`
	HeaderBillingIsConfirmed          *bool    `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus           *string  `json:"HeaderBillingConfStatus"`
	TotalNetAmount                    float32  `json:"TotalNetAmount"`
	TotalTaxAmount                    float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                  float32  `json:"TotalGrossAmount"`
	TransactionCurrency               string   `json:"TransactionCurrency"`
	Incoterms                         *string  `json:"Incoterms"`
	PaymentTerms                      string   `json:"PaymentTerms"`
	DueCalculationBaseDate            *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                    *string  `json:"PaymentDueDate"`
	NetPaymentDays                    *int     `json:"NetPaymentDays"`
	PaymentMethod                     string   `json:"PaymentMethod"`
	ExternalReferenceDocument         *string  `json:"ExternalReferenceDocument"`
	DocumentHeaderText                *string  `json:"DocumentHeaderText"`
	HeaderIsCleared                   *bool    `json:"HeaderIsCleared"`
	HeaderPaymentBlockStatus          *bool    `json:"HeaderPaymentBlockStatus"`
	HeaderPaymentRequisitionIsCreated *bool    `json:"HeaderPaymentRequisitionIsCreated"`
	CreationDate                      string   `json:"CreationDate"`
	CreationTime                      string   `json:"CreationTime"`
	LastChangeDate                    string   `json:"LastChangeDate"`
	LastChangeTime                    string   `json:"LastChangeTime"`
	IsCancelled                       *bool    `json:"IsCancelled"`
}
