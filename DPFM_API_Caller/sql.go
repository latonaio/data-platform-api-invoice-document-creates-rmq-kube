package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-invoice-document-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemPricingElement *[]dpfm_api_output_formatter.ItemPricingElement
	var address *[]dpfm_api_output_formatter.Address
	var partner *[]dpfm_api_output_formatter.Partner
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Item":
			item = c.itemCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemPricingElement":
			itemPricingElement = c.itemPricingElementCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Partner":
			partner = c.partnerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Address":
			address = c.addressCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Item:               item,
		ItemPricingElement: itemPricingElement,
		Address:            address,
		Partner:            partner,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var partner *[]dpfm_api_output_formatter.Partner
	var address *[]dpfm_api_output_formatter.Address
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Item":
			item = c.itemUpdateSql(mtx, input, output, errs, log)
		case "Partner":
			partner = c.partnerUpdateSql(mtx, input, output, errs, log)
		case "Address":
			address = c.addressUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:  header,
		Item:    item,
		Partner: partner,
		Address: address,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_invoice_document_header_dataの更新
	for _, headerData := range *subfuncSDC.Message.Header {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "InvoiceDocumentHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot insert"
			return nil
		}

		if output.SQLUpdateResult == nil {
			output.SQLUpdateResult = getBoolPtr(true)
		}
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_invoice_document_item_dataの更新
	for _, itemData := range *subfuncSDC.Message.Item {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "InvoiceDocumentItem", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemPricingElementCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPricingElement {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_invoice_document_item_pricing_element_dataの更新
	for _, itemPricingElementData := range *subfuncSDC.Message.ItemPricingElement {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemPricingElementData, "function": "InvoiceDocumentItemPricingElement", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Pricing Element Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElementCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_invoice_document_partner_dataの更新
	for _, partnerData := range *subfuncSDC.Message.Partner {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "InvoiceDocumentPartner", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Partner Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) addressCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_invoice_document_address_dataの更新
	for _, addressData := range *subfuncSDC.Message.Address {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": addressData, "function": "InvoiceDocumentAddress", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Address Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAddressCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "InvoiceDocumentHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	req := make([]dpfm_api_processing_formatter.ItemUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		itemData := *dpfm_api_processing_formatter.ConvertToItemUpdates(header, item)

		if itemIsUpdate(&itemData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "InvoiceDocumentItem", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Item Data cannot insert"
				return nil
			}
		}
		req = append(req, itemData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	req := make([]dpfm_api_processing_formatter.PartnerUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, partner := range header.Partner {
		partnerData := *dpfm_api_processing_formatter.ConvertToPartnerUpdates(header, partner)

		if partnerIsUpdate(&partnerData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "InvoiceDocumentPartner", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Partner Data cannot insert"
				return nil
			}
		}
		req = append(req, partnerData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) addressUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	req := make([]dpfm_api_processing_formatter.AddressUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, address := range header.Address {
		addressData := *dpfm_api_processing_formatter.ConvertToAddressUpdates(header, address)

		if addressIsUpdate(&addressData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": addressData, "function": "InvoiceDocumentAddress", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Address Data cannot insert"
				return nil
			}
		}
		req = append(req, addressData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAddressUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	invoiceDocument := header.InvoiceDocument

	return !(invoiceDocument == 0)
}

func itemIsUpdate(item *dpfm_api_processing_formatter.ItemUpdates) bool {
	invoiceDocument := item.InvoiceDocument
	invoiceDocumentItem := item.InvoiceDocumentItem

	return !(invoiceDocument == 0 || invoiceDocumentItem == 0)
}

func partnerIsUpdate(partner *dpfm_api_processing_formatter.PartnerUpdates) bool {
	invoiceDocument := partner.InvoiceDocument
	partnerFunction := partner.PartnerFunction
	businessPartner := partner.BusinessPartner

	return !(invoiceDocument == 0 || partnerFunction == "" || businessPartner == 0)
}

func addressIsUpdate(address *dpfm_api_processing_formatter.AddressUpdates) bool {
	invoiceDocument := address.InvoiceDocument
	addressID := address.AddressID

	return !(invoiceDocument == 0 || addressID == 0)
}
