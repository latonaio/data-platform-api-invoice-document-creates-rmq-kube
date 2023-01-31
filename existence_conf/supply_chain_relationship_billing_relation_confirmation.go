package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) supplyChainRelationshipBillingRelationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0
	headers := make([]dpfm_api_input_reader.Header, 0, 1)
	headers = append(headers, input.Header)
	for _, header := range headers {
		for _, item := range header.Item {
			supplyChainRelationshipID, supplyChainRelationshipBillingID, buyer, seller, billToParty, billFromParty, err := getSupplyChainRelationshipBillingRelationExistenceConfKey(mapper, &header, &item, exconfErrMsg)
			if err != nil {
				*errs = append(*errs, err)
				return
			}
			queueName, err := getQueueName(mapper)
			if err != nil {
				*errs = append(*errs, err)
				return
			}
			wg2.Add(1)
			exReqTimes++
			go func() {
				res, err := c.supplyChainRelationshipBillingRelationExistenceConfRequest(supplyChainRelationshipID, supplyChainRelationshipBillingID, buyer, seller, billToParty, billFromParty, queueName, input, existenceMap, mtx, log)
				if err != nil {
					mtx.Lock()
					*errs = append(*errs, err)
					mtx.Unlock()
				}
				if res != "" {
					*exconfErrMsg = res
				}
				wg2.Done()
			}()
		}
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) supplyChainRelationshipBillingRelationExistenceConfRequest(supplyChainRelationshipID int, supplyChainRelationshipBillingID int, buyer int, seller int, billToParty int, billFromParty int, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"SupplyChainRelationshipID":        supplyChainRelationshipID,
		"SupplyChainRelationshipBillingID": supplyChainRelationshipBillingID,
		"Buyer":                            buyer,
		"Seller":                           seller,
		"BillToParty":                      billToParty,
		"BillFromParty":                    billFromParty,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}

	data := &SupplyChainRelationshipBillingRelationReturn{
		SupplyChainRelationshipID:        supplyChainRelationshipID,
		SupplyChainRelationshipBillingID: supplyChainRelationshipBillingID,
		Buyer:                            buyer,
		Seller:                           seller,
		BillToParty:                      billToParty,
		BillFromParty:                    billFromParty,
	}
	req.SupplyChainRelationshipBillingRelationReturn = data

	exist, err = c.exconfRequest(req, queueName, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getSupplyChainRelationshipBillingRelationExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (int, int, int, int, int, int, error) {
	var supplyChainRelationshipID, supplyChainRelationshipBillingID, buyer, seller, billToParty, billFromParty int
	var err error

	if header.SupplyChainRelationshipID == nil ||
		header.SupplyChainRelationshipBillingID == nil ||
		item.Buyer == nil ||
		item.Seller == nil ||
		header.BillToParty == nil ||
		header.BillFromParty == nil {
		err = xerrors.Errorf("cannot specify null keys")
		return 0, 0, 0, 0, 0, 0, err
	}
	supplyChainRelationshipID = *header.SupplyChainRelationshipID
	supplyChainRelationshipBillingID = *header.SupplyChainRelationshipBillingID
	buyer = *item.Buyer
	seller = *item.Seller
	billToParty = *header.BillToParty
	billFromParty = *header.BillFromParty

	return supplyChainRelationshipID, supplyChainRelationshipBillingID, buyer, seller, billToParty, billFromParty, nil
}
