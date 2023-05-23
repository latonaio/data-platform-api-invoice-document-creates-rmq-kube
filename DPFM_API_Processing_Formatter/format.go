package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		InvoiceDocument:                   data.InvoiceDocument,
		InvoiceDocumentDate:               data.InvoiceDocumentDate,
		InvoiceDocumentTime:               data.InvoiceDocumentTime,
		InvoicePeriodStartDate:            data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:              data.InvoicePeriodEndDate,
		AccountingPostingDate:             data.AccountingPostingDate,
		HeaderBillingIsConfirmed:          data.HeaderBillingIsConfirmed,
		TotalNetAmount:                    data.TotalNetAmount,
		TotalTaxAmount:                    data.TotalTaxAmount,
		TotalGrossAmount:                  data.TotalGrossAmount,
		PaymentTerms:                      data.PaymentTerms,
		DueCalculationBaseDate:            data.DueCalculationBaseDate,
		PaymentDueDate:                    data.PaymentDueDate,
		NetPaymentDays:                    data.NetPaymentDays,
		PaymentMethod:                     data.PaymentMethod,
		ExternalReferenceDocument:         data.ExternalReferenceDocument,
		DocumentHeaderText:                data.DocumentHeaderText,
		HeaderIsCleared:                   data.HeaderIsCleared,
		HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
		HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item

	return &ItemUpdates{
		InvoiceDocument:               dataHeader.InvoiceDocument,
		InvoiceDocumentItem:           data.InvoiceDocumentItem,
		ItemBillingIsConfirmed:        data.ItemBillingIsConfirmed,
		NetAmount:                     data.NetAmount,
		TaxAmount:                     data.TaxAmount,
		GrossAmount:                   data.GrossAmount,
		ExternalReferenceDocument:     data.ExternalReferenceDocument,
		ExternalReferenceDocumentItem: data.ExternalReferenceDocumentItem,
		ItemIsCleared:                 data.ItemIsCleared,
		ItemPaymentBlockStatus:        data.ItemPaymentBlockStatus,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		InvoiceDocument:    dataHeader.InvoiceDocument,
		PartnerFunction:    data.PartnerFunction,
		BusinessPartner:    data.BusinessPartner,
		ExternalDocumentID: data.ExternalDocumentID,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		InvoiceDocument: dataHeader.InvoiceDocument,
		AddressID:       data.AddressID,
		PostalCode:      data.PostalCode,
		LocalRegion:     data.LocalRegion,
		Country:         data.Country,
		District:        data.District,
		StreetName:      data.StreetName,
		CityName:        data.CityName,
		Building:        data.Building,
		Floor:           data.Floor,
		Room:            data.Room,
	}
}
