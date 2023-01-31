package dpfm_api_processing_formatter

func ConvertToHeaderUpdates(headerUpdates HeaderUpdates) *HeaderUpdates {
	data := headerUpdates

	return &HeaderUpdates{
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
		HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
		HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
	}
}

func ConvertToItemUpdates(itemUpdates ItemUpdates) *ItemUpdates {
	data := itemUpdates

	return &ItemUpdates{
		ItemBillingIsConfirmed:        data.ItemBillingIsConfirmed,
		NetAmount:                     data.NetAmount,
		TaxAmount:                     data.TaxAmount,
		GrossAmount:                   data.GrossAmount,
		ExternalReferenceDocument:     data.ExternalReferenceDocument,
		ExternalReferenceDocumentItem: data.ExternalReferenceDocumentItem,
		ItemPaymentBlockStatus:        data.ItemPaymentBlockStatus,
	}
}

func ConvertToPartnerUpdates(partnerUpdates PartnerUpdates) *PartnerUpdates {
	data := partnerUpdates

	return &PartnerUpdates{
		ExternalDocumentID: data.ExternalDocumentID,
	}
}
func ConvertToAddressUpdates(addressUpdates AddressUpdates) *AddressUpdates {
	data := addressUpdates

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
