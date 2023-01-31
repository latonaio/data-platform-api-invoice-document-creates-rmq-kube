package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-creates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-invoice-document-creates-rmq-kube/config"
	"data-platform-api-invoice-document-creates-rmq-kube/existence_conf"
	"data-platform-api-invoice-document-creates-rmq-kube/sub_func_complementer"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient

	configure    *existence_conf.ExistenceConf
	complementer *sub_func_complementer.SubFuncComplementer
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

	confirmor *existence_conf.ExistenceConf,
	complementer *sub_func_complementer.SubFuncComplementer,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:          context.Background(),
		conf:         conf,
		rmq:          rmq,
		configure:    confirmor,
		complementer: complementer,
	}
}

func (c *DPFMAPICaller) AsyncInvoiceDocumentCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	exconfAllExist := false

	exconfFin := make(chan error)
	subFuncFin := make(chan error)

	subfuncSDC := &sub_func_complementer.SDC{}

	// 他PODへ問い合わせ
	wg.Add(1)
	go c.exconfProcess(&mtx, &wg, exconfFin, input, output, &exconfAllExist, &errs, log)
	if input.APIType == "creates" {
		go c.subfuncProcess(&mtx, &wg, subFuncFin, input, output, subfuncSDC, accepter, &errs, log)
	} else if input.APIType == "updates" {
		go func() { subFuncFin <- nil }()
	} else {
		go func() { subFuncFin <- nil }()
	}

	// 処理待ち
	ticker := time.NewTicker(10 * time.Second)
	if err := c.finWait(&mtx, exconfFin, ticker); err != nil || len(errs) != 0 {
		if err != nil {
			errs = append(errs, err)
		}
		return subfuncSDC, errs
	}
	if !exconfAllExist {
		mtx.Lock()
		return subfuncSDC, nil
	}
	wg.Wait()
	for range accepter {
		if err := c.finWait(&mtx, subFuncFin, ticker); err != nil || len(errs) != 0 {
			if err != nil {
				errs = append(errs, err)
			}
			return subfuncSDC, errs
		}
	}

	var response interface{}
	// SQL処理
	if input.APIType == "creates" {
		response = c.createSqlProcess(nil, &mtx, input, output, subfuncSDC, accepter, &errs, log)
	} else if input.APIType == "updates" {
		response = c.updateSqlProcess(nil, &mtx, input, output, accepter, &errs, log)
	}

	return response, nil
}

func (c *DPFMAPICaller) exconfProcess(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	exconfFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	exconfAllExist *bool,
	errs *[]error,
	log *logger.Logger,
) {
	defer wg.Done()
	var e []error
	*exconfAllExist, e = c.configure.Conf(input, output, log)
	if len(e) != 0 {
		mtx.Lock()
		*errs = append(*errs, e...)
		mtx.Unlock()
		exconfFin <- xerrors.New("exconf error")
		return
	}
	exconfFin <- nil
}

func (c *DPFMAPICaller) subfuncProcess(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	subFuncFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) {
	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			c.headerCreate(mtx, wg, subFuncFin, input, output, subfuncSDC, errs, log)
		case "Item":
			c.itemCreate(mtx, wg, subFuncFin, input, output, subfuncSDC, errs, log)
		default:
			wg.Done()
		}
	}
}

func (c *DPFMAPICaller) finWait(
	mtx *sync.Mutex,
	finChan chan error,
	ticker *time.Ticker,
) error {
	select {
	case e := <-finChan:
		if e != nil {
			mtx.Lock()
			return e
		}
	case <-ticker.C:
		return xerrors.New("time out")
	}
	return nil
}

func (c *DPFMAPICaller) headerCreate(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	errFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	var err error = nil
	defer func() {
		errFin <- err
	}()
	defer wg.Done()
	err = c.complementer.ComplementHeader(input, subfuncSDC, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
		return
	}
	output.SubfuncResult = getBoolPtr(true)
	if subfuncSDC.SubfuncResult == nil || !*subfuncSDC.SubfuncResult {
		output.SubfuncResult = getBoolPtr(false)
		output.SubfuncError = subfuncSDC.SubfuncError
		err = xerrors.New(output.SubfuncError)
		return
	}
	return
}

func (c *DPFMAPICaller) itemCreate(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	errFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	var err error = nil
	defer func() {
		errFin <- err
	}()
	defer wg.Done()
	err = c.complementer.ComplementItem(input, subfuncSDC, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
		return
	}
	output.SubfuncResult = getBoolPtr(true)
	if subfuncSDC.SubfuncResult == nil || !*subfuncSDC.SubfuncResult {
		output.SubfuncResult = getBoolPtr(false)
		output.SubfuncError = subfuncSDC.SubfuncError
		err = xerrors.New(output.SubfuncError)
		return
	}
	return
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	d, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := d.(string)
	if !ok {
		return false
	}
	return result == "success"
}

func getBoolPtr(b bool) *bool {
	return &b
}
