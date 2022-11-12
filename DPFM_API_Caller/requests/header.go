package requests

type Header struct {
	D struct {
		InvoiceDocumentstring            string `json:"InvoiceDocument"`
		CreationDatestring               string `json:"CreationDate"`
		LastChangeDatestring             string `json:"LastChangeDate"`
		BillToPartystring                string `json:"BillToParty"`
		BillFromPartystring              string `json:"BillFromParty"`
		InvoiceDocumentDatestring        string `json:"InvoiceDocumentDate"`
		InvoiceDocumentTimestring        string `json:"InvoiceDocumentTime"`
		InvoicePeriodStartDatestring     string `json:"InvoicePeriodStartDate"`
		InvoicePeriodEndDatestring       string `json:"InvoicePeriodEndDate"`
		AccountingPostingDatestring      string `json:"AccountingPostingDate"`
		InvoiceDocumentIsCancelledstring string `json:"InvoiceDocumentIsCancelled"`
		CancelledInvoiceDocumentstring   string `json:"CancelledInvoiceDocument"`
		IsExportDeliverystring           string `json:"IsExportDelivery"`
		HeaderBillingConfStatusstring    string `json:"HeaderBillingConfStatus"`
		TotalNetAmountstring             string `json:"TotalNetAmount"`
		TotalTaxAmountstring             string `json:"TotalTaxAmount"`
		TotalGrossAmountstring           string `json:"TotalGrossAmount"`
		TransactionCurrencystring        string `json:"TransactionCurrency"`
		Incotermsstring                  string `json:"Incoterms"`
		PaymentTermsstring               string `json:"PaymentTerms"`
		DueCalculationBaseDatestring     string `json:"DueCalculationBaseDate"`
		NetPaymentDaysstring             string `json:"NetPaymentDays"`
		PaymentMethodstring              string `json:"PaymentMethod"`
		PaymentBlockingReasonstring      string `json:"PaymentBlockingReason"`
		ExternalReferenceDocumentstring  string `json:"ExternalReferenceDocument"`
		DocumentHeaderTextstring         string `json:"DocumentHeaderText"`
	} `json:"d"`
}
