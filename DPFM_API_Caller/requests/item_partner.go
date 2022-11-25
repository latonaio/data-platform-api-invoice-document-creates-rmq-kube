package requests

type ItemPartner struct {
	InvoiceDocument     *int    `json:"InvoiceDocument"`
	InvoiceDocumentItem *int    `json:"InvoiceDocumentItem"`
	PartnerFunction     *string `json:"PartnerFunction"`
	BusinessPartner     *int    `json:"BusinessPartner"`
}
