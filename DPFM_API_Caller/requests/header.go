package requests

type Header struct {
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
