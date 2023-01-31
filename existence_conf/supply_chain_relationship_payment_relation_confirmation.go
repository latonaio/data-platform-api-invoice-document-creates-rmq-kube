package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) supplyChainRelationshipPaymentRelationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0
	headers := make([]dpfm_api_input_reader.Header, 0, 1)
	headers = append(headers, input.Header)
	for _, header := range headers {
		for _, item := range header.Item {
			supplyChainRelationshipID, supplyChainRelationshipBillingID, supplyChainRelationshipPaymentID, buyer, seller, billToParty, billFromParty, payer, payee, err := getSupplyChainRelationshipPaymentRelationExistenceConfKey(mapper, &header, &item, exconfErrMsg)
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
				res, err := c.supplyChainRelationshipPaymentRelationExistenceConfRequest(supplyChainRelationshipID, supplyChainRelationshipBillingID, supplyChainRelationshipPaymentID, buyer, seller, billToParty, billFromParty, payer, payee, queueName, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) supplyChainRelationshipPaymentRelationExistenceConfRequest(supplyChainRelationshipID int, supplyChainRelationshipBillingID int, supplyChainRelationshipPaymentID int, buyer int, seller int, billToParty int, billFromParty int, payer int, payee int, queueName string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"SupplyChainRelationshipID":        supplyChainRelationshipID,
		"SupplyChainRelationshipBillingID": supplyChainRelationshipBillingID,
		"Buyer":                            buyer,
		"Seller":                           seller,
		"BillToParty":                      billToParty,
		"BillFromParty":                    billFromParty,
		"Payer":                            payer,
		"Payee":                            payee,
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

	data := &SupplyChainRelationshipPaymentRelationReturn{
		SupplyChainRelationshipID:        supplyChainRelationshipID,
		SupplyChainRelationshipBillingID: supplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID: supplyChainRelationshipPaymentID,
		Buyer:                            buyer,
		Seller:                           seller,
		BillToParty:                      billToParty,
		BillFromParty:                    billFromParty,
		Payer:                            payer,
		Payee:                            payee,
	}
	req.SupplyChainRelationshipPaymentRelationReturn = data

	exist, err = c.exconfRequest(req, queueName, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getSupplyChainRelationshipPaymentRelationExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (int, int, int, int, int, int, int, int, int, error) {
	var supplyChainRelationshipID, supplyChainRelationshipBillingID, supplyChainRelationshipPaymentID, buyer, seller, billToParty, billFromParty, payer, payee int
	var err error

	if header.SupplyChainRelationshipID == nil ||
		header.SupplyChainRelationshipBillingID == nil ||
		header.SupplyChainRelationshipPaymentID == nil ||
		item.Buyer == nil ||
		item.Seller == nil ||
		header.BillToParty == nil ||
		header.BillFromParty == nil ||
		header.Payer == nil ||
		header.Payee == nil {
		err = xerrors.Errorf("cannot specify null keys")
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, err
	}
	supplyChainRelationshipID = *header.SupplyChainRelationshipID
	supplyChainRelationshipBillingID = *header.SupplyChainRelationshipBillingID
	supplyChainRelationshipPaymentID = *header.SupplyChainRelationshipPaymentID
	buyer = *item.Buyer
	seller = *item.Seller
	billToParty = *header.BillToParty
	billFromParty = *header.BillFromParty
	payer = *header.Payer
	payee = *header.Payee

	return supplyChainRelationshipID, supplyChainRelationshipBillingID, supplyChainRelationshipPaymentID, buyer, seller, billToParty, billFromParty, payer, payee, nil
}
