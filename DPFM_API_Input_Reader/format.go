package dpfm_api_input_reader

import (
	"data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBpExistenceConf() {

}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.InvoiceDocument
	return &requests.Header{
		InvoiceDocument:            data.InvoiceDocument,
		CreationDate:               data.CreationDate,
		LastChangeDate:             data.LastChangeDate,
		BillToParty:                data.BillToParty,
		BillFromParty:              data.BillFromParty,
		BillToCountry:              data.BillToCountry,
		BillFromCountry:            data.BillFromCountry,
		InvoiceDocumentDate:        data.InvoiceDocumentDate,
		InvoiceDocumentTime:        data.InvoiceDocumentTime,
		InvoicePeriodStartDate:     data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:       data.InvoicePeriodEndDate,
		AccountingPostingDate:      data.AccountingPostingDate,
		InvoiceDocumentIsCancelled: data.InvoiceDocumentIsCancelled,
		CancelledInvoiceDocument:   data.CancelledInvoiceDocument,
		IsExportImportDelivery:     data.IsExportImportDelivery,
		HeaderBillingIsConfirmed:   data.HeaderBillingIsConfirmed,
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
		HeaderPaymentBlockStatus:   data.HeaderPaymentBlockStatus,
		ExternalReferenceDocument:  data.ExternalReferenceDocument,
		DocumentHeaderText:         data.DocumentHeaderText,
	}
}

func (sdc *SDC) ConvertToHeaderPartner(num int) *requests.HeaderPartner {
	dataHeader := sdc.InvoiceDocument
	data := sdc.InvoiceDocument.HeaderPartner[num]
	return &requests.HeaderPartner{
		InvoiceDocument:         dataHeader.InvoiceDocument,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
}

func (sdc *SDC) ConvertToHeaderPartnerContact(hpNum, hpcNum int) *requests.HeaderPartnerContact {
	dataHeader := sdc.InvoiceDocument
	dataHeaderPartner := sdc.InvoiceDocument.HeaderPartner[hpNum]
	data := dataHeaderPartner.HeaderPartnerContact[hpcNum]
	return &requests.HeaderPartnerContact{
		InvoiceDocument:   dataHeader.InvoiceDocument,
		PartnerFunction:   dataHeaderPartner.PartnerFunction,
		BusinessPartner:   dataHeaderPartner.BusinessPartner,
		ContactID:         data.ContactID,
		ContactPersonName: data.ContactPersonName,
		EmailAddress:      data.EmailAddress,
		PhoneNumber:       data.PhoneNumber,
		MobilePhoneNumber: data.MobilePhoneNumber,
		FaxNumber:         data.FaxNumber,
		ContactTag1:       data.ContactTag1,
		ContactTag2:       data.ContactTag2,
		ContactTag3:       data.ContactTag3,
		ContactTag4:       data.ContactTag4,
	}
}
