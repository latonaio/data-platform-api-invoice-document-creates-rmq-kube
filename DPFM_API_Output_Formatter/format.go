package dpfm_api_output_formatter

import (
	"data-platform-api-invoice-document-creates-rmq-kube/sub_func_complementer"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) []HeaderCreates {
	var headerCreates []HeaderCreates

	for _, data := range subfuncSDC.Message.Header {
		headerCreates = append(headerCreates, HeaderCreates{
			InvoiceDocument:                   data.InvoiceDocument,
			CreationDate:                      data.CreationDate,
			LastChangeDate:                    data.LastChangeDate,
			BillFromParty:                     data.BillFromParty,
			BillToParty:                       data.BillToParty,
			BillFromCountry:                   data.BillFromCountry,
			BillToCountry:                     data.BillToCountry,
			Payer:                             data.Payer,
			Payee:                             data.Payee,
			InvoiceDocumentDate:               data.InvoiceDocumentDate,
			InvoiceDocumentTime:               data.InvoiceDocumentTime,
			InvoicePeriodStartDate:            data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:              data.InvoicePeriodEndDate,
			AccountingPostingDate:             data.AccountingPostingDate,
			InvoiceDocumentIsCancelled:        data.InvoiceDocumentIsCancelled,
			CancelledInvoiceDocument:          data.CancelledInvoiceDocument,
			IsExportImportDelivery:            data.IsExportImportDelivery,
			HeaderBillingIsConfirmed:          data.HeaderBillingIsConfirmed,
			HeaderBillingConfStatus:           data.HeaderBillingConfStatus,
			TotalNetAmount:                    data.TotalNetAmount,
			TotalTaxAmount:                    data.TotalTaxAmount,
			TotalGrossAmount:                  data.TotalGrossAmount,
			TransactionCurrency:               data.TransactionCurrency,
			Incoterms:                         data.Incoterms,
			PaymentTerms:                      data.PaymentTerms,
			DueCalculationBaseDate:            data.DueCalculationBaseDate,
			PaymentDueDate:                    data.PaymentDueDate,
			NetPaymentDays:                    data.NetPaymentDays,
			PaymentMethod:                     data.PaymentMethod,
			HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
			ExternalReferenceDocument:         data.ExternalReferenceDocument,
			DocumentHeaderText:                data.DocumentHeaderText,
			HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
		})
	}

	return headerCreates

}

// func ConvertToHeaderUpdates(headerUpdates dpfm_api_input_reader.HeaderUpdates) *HeaderUpdates {
// 	data := headerUpdates

// 	return &HeaderUpdates{
// 		InvoiceDocument:                   data.InvoiceDocument,
// 		CreationDate:                      data.CreationDate,
// 		LastChangeDate:                    data.LastChangeDate,
// 		BillToParty:                       data.BillToParty,
// 		BillFromParty:                     data.BillFromParty,
// 		BillToCountry:                     data.BillToCountry,
// 		BillFromCountry:                   data.BillFromCountry,
// 		InvoiceDocumentDate:               data.InvoiceDocumentDate,
// 		InvoiceDocumentTime:               data.InvoiceDocumentTime,
// 		InvoicePeriodStartDate:            data.InvoicePeriodStartDate,
// 		InvoicePeriodEndDate:              data.InvoicePeriodEndDate,
// 		AccountingPostingDate:             data.AccountingPostingDate,
// 		InvoiceDocumentIsCancelled:        data.InvoiceDocumentIsCancelled,
// 		CancelledInvoiceDocument:          data.CancelledInvoiceDocument,
// 		IsExportImportDelivery:            data.IsExportImportDelivery,
// 		HeaderBillingIsConfirmed:          data.HeaderBillingIsConfirmed,
// 		HeaderBillingConfStatus:           data.HeaderBillingConfStatus,
// 		TotalNetAmount:                    data.TotalNetAmount,
// 		TotalTaxAmount:                    data.TotalTaxAmount,
// 		TotalGrossAmount:                  data.TotalGrossAmount,
// 		TransactionCurrency:               data.TransactionCurrency,
// 		Incoterms:                         data.Incoterms,
// 		PaymentTerms:                      data.PaymentTerms,
// 		DueCalculationBaseDate:            data.DueCalculationBaseDate,
// 		NetPaymentDays:                    data.NetPaymentDays,
// 		PaymentMethod:                     data.PaymentMethod,
// 		HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
// 		ExternalReferenceDocument:         data.ExternalReferenceDocument,
// 		DocumentHeaderText:                data.DocumentHeaderText,
// 		HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
// 	}
// }

func ConvertToHeaderPartner(subfuncSDC *sub_func_complementer.SDC) []HeaderPartner {
	var headerPartner []HeaderPartner

	for _, data := range subfuncSDC.Message.HeaderPartner {
		headerPartner = append(headerPartner, HeaderPartner{
			InvoiceDocument:         data.InvoiceDocument,
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
		})
	}

	return headerPartner
}

// func ConvertToHeaderPartnerContact(subfuncSDC *sub_func_complementer.SDC) []HeaderPartnerContact {
// 	var headerPartnerContact []HeaderPartnerContact

// 	for _, data := range subfuncSDC.Message.HeaderPartnerContact {
// 		headerPartnerContact = append(headerPartnerContact, HeaderPartnerContact{
// 			InvoiceDocument:   data.InvoiceDocument,
// 			PartnerFunction:   data.PartnerFunction,
// 			BusinessPartner:   data.BusinessPartner,
// 			ContactID:         data.ContactID,
// 			ContactPersonName: data.ContactPersonName,
// 			EmailAddress:      data.EmailAddress,
// 			PhoneNumber:       data.PhoneNumber,
// 			MobilePhoneNumber: data.MobilePhoneNumber,
// 			FaxNumber:         data.FaxNumber,
// 			ContactTag1:       data.ContactTag1,
// 			ContactTag2:       data.ContactTag2,
// 			ContactTag3:       data.ContactTag3,
// 			ContactTag4:       data.ContactTag4,
// 		})
// 	}

// 	return headerPartnerContact
// }

// func ConvertToHeaderPDF(subfuncSDC *sub_func_complementer.SDC) []HeaderPDF {
// 	var headerPDF []HeaderPDF

// 	for _, data := range subfuncSDC.Message.HeaderPDF {
// 		headerPDF = append(headerPDF, HeaderPDF{
// 			InvoiceDocument:          data.InvoiceDocument,
// 			DocType:                  data.DocType,
// 			DocVersionID:             data.DocVersionID,
// 			DocID:                    data.DocID,
// 			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
// 			FileName:                 data.FileName,
// 		})
// 	}

// 	return headerPDF
// }

// func ConvertToAddress(subfuncSDC *sub_func_complementer.SDC) []Address {
// 	var address []Address

// 	for _, data := range subfuncSDC.Message.Address {
// 		address = append(address, Address{
// 			InvoiceDocument: data.InvoiceDocument,
// 			AddressID:       data.AddressID,
// 			PostalCode:      data.PostalCode,
// 			LocalRegion:     data.LocalRegion,
// 			Country:         data.Country,
// 			District:        data.District,
// 			StreetName:      data.StreetName,
// 			CityName:        data.CityName,
// 			Building:        data.Building,
// 			Floor:           data.Floor,
// 			Room:            data.Room,
// 		})
// 	}

// 	return address
// }

func ConvertToItem(subfuncSDC *sub_func_complementer.SDC) []Item {
	var item []Item

	for _, data := range subfuncSDC.Message.Item {
		item = append(item, Item{
			InvoiceDocument:                 data.InvoiceDocument,
			InvoiceDocumentItem:             data.InvoiceDocumentItem,
			InvoiceDocumentItemCategory:     data.InvoiceDocumentItemCategory,
			InvoiceDocumentItemText:         data.InvoiceDocumentItemText,
			InvoiceDocumentItemTextByBuyer:  data.InvoiceDocumentItemTextByBuyer,
			InvoiceDocumentItemTextBySeller: data.InvoiceDocumentItemTextBySeller,
			CreationDate:                    data.CreationDate,
			CreationTime:                    data.CreationTime,
			ItemBillingIsConfirmed:          data.ItemBillingIsConfirmed,
			ItemBillingConfStatus:           data.ItemBillingConfStatus,
			Buyer:                           data.Buyer,
			Seller:                          data.Seller,
			DeliverFromParty:                data.DeliverFromParty,
			DeliverToParty:                  data.DeliverToParty,
			Product:                         data.Product,
			ProductGroup:                    data.ProductGroup,
			ProductStandardID:               data.ProductStandardID,
			ProductionPlantPartnerFunction:  data.ProductionPlantPartnerFunction,
			ProductionPlantBusinessPartner:  data.ProductionPlantBusinessPartner,
			ProductionPlant:                 data.ProductionPlant,
			ProductionPlantStorageLocation:  data.ProductionPlantStorageLocation,
			IssuingPlantPartnerFunction:     data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:     data.IssuingPlantBusinessPartner,
			IssuingPlant:                    data.IssuingPlant,
			IssuingPlantStorageLocation:     data.IssuingPlantStorageLocation,
			ReceivingPlantPartnerFunction:   data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:   data.ReceivingPlantBusinessPartner,
			ReceivingPlant:                  data.ReceivingPlant,
			ReceivingPlantStorageLocation:   data.ReceivingPlantStorageLocation,
			ServicesRenderedDate:            data.ServicesRenderedDate,
			InvoiceQuantity:                 data.InvoiceQuantity,
			InvoiceQuantityUnit:             data.InvoiceQuantityUnit,
			InvoiceQuantityInBaseUnit:       data.InvoiceQuantityInBaseUnit,
			BaseUnit:                        data.BaseUnit,
			ActualGoodsIssueDate:            data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:            data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:          data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:          data.ActualGoodsReceiptTime,
			ItemGrossWeight:                 data.ItemGrossWeight,
			ItemNetWeight:                   data.ItemNetWeight,
			ItemWeightUnit:                  data.ItemWeightUnit,
			NetAmount:                       data.NetAmount,
			TaxAmount:                       data.TaxAmount,
			GrossAmount:                     data.GrossAmount,
			ItemPaymentBlockStatus:          data.ItemPaymentBlockStatus,
			GoodsIssueOrReceiptSlipNumber:   data.GoodsIssueOrReceiptSlipNumber,
			TransactionCurrency:             data.TransactionCurrency,
			PricingDate:                     data.PricingDate,
			ProductTaxClassification:        data.ProductTaxClassification,
			Project:                         data.Project,
			OrderID:                         data.OrderID,
			OrderItem:                       data.OrderItem,
			OrderType:                       data.OrderType,
			ContractType:                    data.ContractType,
			OrderVaridityStartDate:          data.OrderVaridityStartDate,
			OrderVaridityEndDate:            data.OrderVaridityEndDate,
			InvoiceScheduleStartDate:        data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:          data.InvoiceScheduleEndDate,
			DeliveryDocument:                data.DeliveryDocument,
			DeliveryDocumentItem:            data.DeliveryDocumentItem,
			OriginDocument:                  data.OriginDocument,
			OriginDocumentItem:              data.OriginDocumentItem,
			ReferenceDocument:               data.ReferenceDocument,
			ReferenceDocumentItem:           data.ReferenceDocumentItem,
			ReferenceDocumentType:           data.ReferenceDocumentType,
			ExternalReferenceDocument:       data.ExternalReferenceDocument,
			ExternalReferenceDocumentItem:   data.ExternalReferenceDocumentItem,
			TaxCode:                         data.TaxCode,
			TaxRate:                         data.TaxRate,
			CountryOfOrigin:                 data.CountryOfOrigin,
			ItemPaymentRequisitionIsCreated: data.ItemPaymentRequisitionIsCreated,
		})
	}

	return item
}

// func ConvertToItemPartner(subfuncSDC *sub_func_complementer.SDC) []ItemPartner {
// 	var itemPartner []ItemPartner

// 	for _, data := range subfuncSDC.Message.ItemPartner {
// 		itemPartner = append(itemPartner, ItemPartner{
// 			InvoiceDocument:     data.InvoiceDocument,
// 			InvoiceDocumentItem: data.InvoiceDocumentItem,
// 			PartnerFunction:     data.PartnerFunction,
// 			BusinessPartner:     data.BusinessPartner,
// 		})
// 	}

// 	return itemPartner
// }

// func ConvertToItemPricingElement(subfuncSDC *sub_func_complementer.SDC) []ItemPricingElement {
// 	var itemPricingElement []ItemPricingElement

// 	for _, data := range subfuncSDC.Message.ItemPricingElement {
// 		itemPricingElement = append(itemPricingElement, ItemPricingElement{
// 			InvoiceDocument:            data.InvoiceDocument,
// 			InvoiceDocumentItem:        data.InvoiceDocumentItem,
// 			PricingProcedureStep:       data.PricingProcedureStep,
// 			PricingProcedureCounter:    data.PricingProcedureCounter,
// 			ConditionType:              data.ConditionType,
// 			PricingDate:                data.PricingDate,
// 			ConditionRateValue:         data.ConditionRateValue,
// 			ConditionCurrency:          data.ConditionCurrency,
// 			ConditionQuantity:          data.ConditionQuantity,
// 			ConditionQuantityUnit:      data.ConditionQuantityUnit,
// 			ConditionRecord:            data.ConditionRecord,
// 			ConditionSequentialNumber:  data.ConditionSequentialNumber,
// 			TaxCode:                    data.TaxCode,
// 			ConditionAmount:            data.ConditionAmount,
// 			TransactionCurrency:        data.TransactionCurrency,
// 			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
// 		})
// 	}

// 	return itemPricingElement
// }
