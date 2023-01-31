package dpfm_api_output_formatter

import (
	"data-platform-api-invoice-document-creates-rmq-kube/sub_func_complementer"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) *HeaderCreates {
	data := subfuncSDC.Message.Header

	headerCreates := &HeaderCreates{
		InvoiceDocument:                   data.InvoiceDocument,
		CreationDate:                      data.CreationDate,
		CreationTime:                      data.CreationTime,
		LastChangeDate:                    data.LastChangeDate,
		LastChangeTime:                    data.LastChangeTime,
		SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID:  data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID:  data.SupplyChainRelationshipPaymentID,
		BillToParty:                       data.BillToParty,
		BillFromParty:                     data.BillFromParty,
		BillToCountry:                     data.BillToCountry,
		BillFromCountry:                   data.BillFromCountry,
		Payer:                             data.Payer,
		Payee:                             data.Payee,
		InvoiceDocumentDate:               data.InvoiceDocumentDate,
		InvoiceDocumentTime:               data.InvoiceDocumentTime,
		InvoicePeriodStartDate:            data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:              data.InvoicePeriodEndDate,
		AccountingPostingDate:             data.AccountingPostingDate,
		IsExportImport:                    data.IsExportImport,
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
		ExternalReferenceDocument:         data.ExternalReferenceDocument,
		DocumentHeaderText:                data.DocumentHeaderText,
		HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
		HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
		InvoiceDocumentIsCancelled:        data.InvoiceDocumentIsCancelled,
		CancelledInvoiceDocument:          data.CancelledInvoiceDocument,
	}

	return headerCreates

}

func ConvertToPartner(subfuncSDC *sub_func_complementer.SDC) *[]Partner {
	var partner []Partner

	for _, data := range *subfuncSDC.Message.Partner {
		partner = append(partner, Partner{
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

	return &partner
}

func ConvertToHeaderDoc(subfuncSDC *sub_func_complementer.SDC) *[]HeaderDoc {
	var headerDoc []HeaderDoc

	for _, data := range *subfuncSDC.Message.HeaderDoc {
		headerDoc = append(headerDoc, HeaderDoc{
			InvoiceDocument:          data.InvoiceDocument,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}

	return &headerDoc
}

func ConvertToAddress(subfuncSDC *sub_func_complementer.SDC) *[]Address {
	var address []Address

	for _, data := range *subfuncSDC.Message.Address {
		address = append(address, Address{
			InvoiceDocument: data.InvoiceDocument,
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
		})
	}

	return &address
}

func ConvertToItem(subfuncSDC *sub_func_complementer.SDC) *[]Item {
	var item []Item

	for _, data := range *subfuncSDC.Message.Item {
		item = append(item, Item{
			InvoiceDocument:                         data.InvoiceDocument,
			InvoiceDocumentItem:                     data.InvoiceDocumentItem,
			InvoiceDocumentItemCategory:             data.InvoiceDocumentItemCategory,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:       data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:  data.SupplyChainRelationshipDeliveryPlantID,
			InvoiceDocumentItemText:                 data.InvoiceDocumentItemText,
			InvoiceDocumentItemTextByBuyer:          data.InvoiceDocumentItemTextByBuyer,
			InvoiceDocumentItemTextBySeller:         data.InvoiceDocumentItemTextBySeller,
			Product:                                 data.Product,
			ProductGroup:                            data.ProductGroup,
			ProductStandardID:                       data.ProductStandardID,
			CreationDate:                            data.CreationDate,
			CreationTime:                            data.CreationTime,
			LastChangeDate:                          data.LastChangeDate,
			LastChangeTime:                          data.LastChangeTime,
			ItemBillingIsConfirmed:                  data.ItemBillingIsConfirmed,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			DeliverToParty:                          data.DeliverToParty,
			DeliverFromParty:                        data.DeliverFromParty,
			DeliverToPlant:                          data.DeliverToPlant,
			DeliverToPlantStorageLocation:           data.DeliverToPlantStorageLocation,
			DeliverFromPlant:                        data.DeliverFromPlant,
			DeliverFromPlantStorageLocation:         data.DeliverFromPlantStorageLocation,
			ProductionPlantBusinessPartner:          data.ProductionPlantBusinessPartner,
			ProductionPlant:                         data.ProductionPlant,
			ProductionPlantStorageLocation:          data.ProductionPlantStorageLocation,
			ServicesRenderedDate:                    data.ServicesRenderedDate,
			InvoiceQuantity:                         data.InvoiceQuantity,
			InvoiceQuantityUnit:                     data.InvoiceQuantityUnit,
			InvoiceQuantityInBaseUnit:               data.InvoiceQuantityInBaseUnit,
			BaseUnit:                                data.BaseUnit,
			ActualGoodsIssueDate:                    data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:                    data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:                  data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:                  data.ActualGoodsReceiptTime,
			ItemGrossWeight:                         data.ItemGrossWeight,
			ItemNetWeight:                           data.ItemNetWeight,
			ItemWeightUnit:                          data.ItemWeightUnit,
			NetAmount:                               data.NetAmount,
			TaxAmount:                               data.TaxAmount,
			GrossAmount:                             data.GrossAmount,
			GoodsIssueOrReceiptSlipNumber:           data.GoodsIssueOrReceiptSlipNumber,
			TransactionCurrency:                     data.TransactionCurrency,
			PricingDate:                             data.PricingDate,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                data.DefinedTaxClassification,
			Project:                                 data.Project,
			OrderID:                                 data.OrderID,
			OrderItem:                               data.OrderItem,
			OrderType:                               data.OrderType,
			ContractType:                            data.ContractType,
			OrderValidityStartDate:                  data.OrderValidityStartDate,
			OrderValidityEndDate:                    data.OrderValidityEndDate,
			InvoicePeriodStartDate:                  data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:                    data.InvoicePeriodEndDate,
			DeliveryDocument:                        data.DeliveryDocument,
			DeliveryDocumentItem:                    data.DeliveryDocumentItem,
			OriginDocument:                          data.OriginDocument,
			OriginDocumentItem:                      data.OriginDocumentItem,
			ReferenceDocument:                       data.ReferenceDocument,
			ReferenceDocumentItem:                   data.ReferenceDocumentItem,
			ExternalReferenceDocument:               data.ExternalReferenceDocument,
			ExternalReferenceDocumentItem:           data.ExternalReferenceDocumentItem,
			TaxCode:                                 data.TaxCode,
			TaxRate:                                 data.TaxRate,
			CountryOfOrigin:                         data.CountryOfOrigin,
			CountryOfOriginLanguage:                 data.CountryOfOriginLanguage,
			ItemPaymentRequisitionIsCreated:         data.ItemPaymentRequisitionIsCreated,
			ItemPaymentBlockStatus:                  data.ItemPaymentBlockStatus,
		})
	}

	return &item
}

func ConvertToItemPricingElement(subfuncSDC *sub_func_complementer.SDC) *[]ItemPricingElement {
	var itemPricingElement []ItemPricingElement

	for _, data := range *subfuncSDC.Message.ItemPricingElement {
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			InvoiceDocument:            data.InvoiceDocument,
			InvoiceDocumentItem:        data.InvoiceDocumentItem,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			ConditionQuantityUnit:      data.ConditionQuantityUnit,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
		})
	}

	return &itemPricingElement
}
