package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		InvoiceDocumentDate:               *data.InvoiceDocumentDate,
		InvoiceDocumentTime:               *data.InvoiceDocumentTime,
		InvoicePeriodStartDate:            *data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:              *data.InvoicePeriodEndDate,
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

func ConvertToItemUpdates(item dpfm_api_input_reader.Item) *ItemUpdates {
	data := item

	return &ItemUpdates{
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

func ConvertToPartnerUpdates(partnerUpdates PartnerUpdates) *PartnerUpdates {
	data := partnerUpdates

	return &PartnerUpdates{
		ExternalDocumentID: data.ExternalDocumentID,
	}
}

func ConvertToAddressUpdates(address dpfm_api_input_reader.Address) *AddressUpdates {
	data := address

	return &AddressUpdates{
		PostalCode:  data.PostalCode,
		LocalRegion: data.LocalRegion,
		Country:     data.Country,
		District:    data.District,
		StreetName:  data.StreetName,
		CityName:    data.CityName,
		Building:    data.Building,
		Floor:       data.Floor,
		Room:        data.Room,
	}
}
