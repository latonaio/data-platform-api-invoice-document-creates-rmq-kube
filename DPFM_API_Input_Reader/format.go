package dpfm_api_input_reader

import (
	"data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.DeliveryDocument
	return &requests.Header{
		InvoiceDocument:            data.InvoiceDocument,
		CreationDate:               data.CreationDate,
		LastChangeDate:             data.LastChangeDate,
		BillToParty:                data.BillToParty,
		BillFromParty:              data.BillFromParty,
		InvoiceDocumentDate:        data.InvoiceDocumentDate,
		InvoiceDocumentTime:        data.InvoiceDocumentTime,
		InvoicePeriodStartDate:     data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:       data.InvoicePeriodEndDate,
		AccountingPostingDate:      data.AccountingPostingDate,
		InvoiceDocumentIsCancelled: data.InvoiceDocumentIsCancelled,
		CancelledInvoiceDocument:   data.CancelledInvoiceDocument,
		IsExportDelivery:           data.IsExportDelivery,
		HeaderBillingConfStatus:    data.HeaderBillingConfStatus,
		TotalNetAmount:             data.TotalNetAmount,
		TotalTaxAmount:             data.TotalTaxAmount,
		TotalGrossAmount:           data.TotalGrossAmount,
		TransactionCurrency:        data.TransactionCurrency,
		Incoterms:                  data.Incoterms,
		PaymentTerms:               data.PaymentTerms,
		DueCalculationBaseDate:     data.DueCalculationBaseDate,
		NetPaymentDays:             data.NetPaymentDays,
		PaymentMethod:              data.PaymentMethod,
		PaymentBlockingReason:      data.PaymentBlockingReason,
		ExternalReferenceDocument:  data.ExternalReferenceDocument,
		DocumentHeaderText:         data.DocumentHeaderText,
	}
}
