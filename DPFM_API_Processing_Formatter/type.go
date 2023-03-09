package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InvoiceDocumentDate               string   `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime               string   `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate            string   `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              string   `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string  `json:"AccountingPostingDate"`
	HeaderBillingIsConfirmed          *bool    `json:"HeaderBillingIsConfirmed"`
	TotalNetAmount                    *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                    *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                  *float32 `json:"TotalGrossAmount"`
	PaymentTerms                      *string  `json:"PaymentTerms"`
	DueCalculationBaseDate            *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                    *string  `json:"PaymentDueDate"`
	NetPaymentDays                    *int     `json:"NetPaymentDays"`
	PaymentMethod                     *string  `json:"PaymentMethod"`
	ExternalReferenceDocument         *string  `json:"ExternalReferenceDocument"`
	DocumentHeaderText                *string  `json:"DocumentHeaderText"`
	HeaderIsCleared                   *bool    `json:"HeaderIsCleared"`
	HeaderPaymentBlockStatus          *bool    `json:"HeaderPaymentBlockStatus"`
	HeaderPaymentRequisitionIsCreated *bool    `json:"HeaderPaymentRequisitionIsCreated"`
}

type PartnerUpdates struct {
	ExternalDocumentID *string `json:"ExternalDocumentID"`
}

type AddressUpdates struct {
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}

type ItemUpdates struct {
	ItemBillingIsConfirmed        *bool    `json:"ItemBillingIsConfirmed"`
	NetAmount                     *float32 `json:"NetAmount"`
	TaxAmount                     *float32 `json:"TaxAmount"`
	GrossAmount                   *float32 `json:"GrossAmount"`
	ExternalReferenceDocument     *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem *string  `json:"ExternalReferenceDocumentItem"`
	ItemIsCleared                 *bool    `json:"ItemIsCleared"`
	ItemPaymentBlockStatus        *bool    `json:"ItemPaymentBlockStatus"`
}
