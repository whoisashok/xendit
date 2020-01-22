package usecase

import (
	"fmt"
	"io/ioutil"
	"net/http"
	cfg "xendit/config"
	serviceXendit "xendit/xendit_services"
	entity "xendit/xendit_services/entity"
	httpClient "xendit/xendit_services/http_client"
	vmXendit "xendit/xendit_services/view_model"

	helperXendit "xendit/xendit_services/helper"
	helperRandomStringXendit "xendit/xendit_services/helper/random_string"

	logging "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var config cfg.Config

type XenditUsecaseImpl struct {
	//XenditRegisterRepo repository.XenditRepository
	logger logging.Logger
}

func NewXenditUsecaseImpl(logger logging.Logger) XenditUsecase {
	return &XenditUsecaseImpl{logger}
}

// Create Disbursement Request
func (u *XenditUsecaseImpl) CreateDisbursementHandler(req entity.DisbursementRequestBody) (*vmXendit.XenditResponse, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	req.ExternalID = helperXendit.DisbursementPrefix + helperRandomStringXendit.GetRandomString(10)

	resp, err := clientService.XenditService.CreateDisbursementHandler(req)
	if err != nil {
		level.Error(u.logger).Log("function", "XenditUsecaseImpl createDisbursementHandler", "error", err)
		return nil, err
	}

	CreateDisbursementHandler, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("CreateDisbursementHandler	=", string(CreateDisbursementHandler))

	vmResHTTPJSON := string(CreateDisbursementHandler)
	xenditResponse := vmXendit.XenditResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceXendit.SuccessCodeType,
		Message:  serviceXendit.MessageRetrievedSuccess,
		Data:     vmResHTTPJSON,
	}

	return &xenditResponse, nil
}

// Get disbursement by id
func (u *XenditUsecaseImpl) GetDisbursementByIdHandler(req entity.GetDisbursementByIdRequest) (*vmXendit.XenditResponse, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	resp, err := clientService.XenditService.GetDisbursementByIdHandler(req)
	if err != nil {
		level.Error(u.logger).Log("function", "XenditUsecaseImpl GetDisbursementByIdHandler", "error", err)
		return nil, err
	}

	GetDisbursementByIdHandler, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("GetDisbursementById	=", string(GetDisbursementByIdHandler))

	vmResHTTPJSON := string(GetDisbursementByIdHandler)
	xenditResponse := vmXendit.XenditResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceXendit.SuccessCodeType,
		Message:  serviceXendit.MessageRetrievedSuccess,
		Data:     vmResHTTPJSON,
	}

	return &xenditResponse, nil
}

// Get disbursement by external_id
func (u *XenditUsecaseImpl) GetDisbursementByExternalIDHandler(req entity.GetDisbursementByExternalIDRequest) (*vmXendit.XenditResponse, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	resp, err := clientService.XenditService.GetDisbursementByExternalIDHandler(req)
	if err != nil {
		level.Error(u.logger).Log("function", "XenditUsecaseImpl GetDisbursementByExternalIDHandler", "error", err)
		return nil, err
	}

	GetDisbursementByExternalIDHandler, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("GetDisbursementByExternalIDHandler	=", string(GetDisbursementByExternalIDHandler))

	vmResHTTPJSON := string(GetDisbursementByExternalIDHandler)
	xenditResponse := vmXendit.XenditResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceXendit.SuccessCodeType,
		Message:  serviceXendit.MessageRetrievedSuccess,
		Data:     vmResHTTPJSON,
	}

	return &xenditResponse, nil
}

// Disbursement callback
func (u *XenditUsecaseImpl) DisbursementCallbackHandler(entity.DisbursementCallBackRequest) (*vmXendit.XenditResponse, error) {

	vmResHTTPJSON := "result"
	xenditResponse := vmXendit.XenditResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceXendit.SuccessCodeType,
		Message:  serviceXendit.MessageRetrievedSuccess,
		Data:     vmResHTTPJSON,
	}

	return &xenditResponse, nil
}

// Create batch disbursement
func (u *XenditUsecaseImpl) CreateBulkDisbursementHandler(req entity.DisbursementBulkRequest) (*vmXendit.XenditResponse, error) {

	req.Reference = helperXendit.BatchReferencePrefix + helperRandomStringXendit.GetRandomString(10)

	var disbursementBulkRequestLists []entity.DisbursementBulkRequestLists
	for _, disb := range req.DisbursementBulkRequestLists {

		disbursementBulkRequestLists = append(disbursementBulkRequestLists, entity.DisbursementBulkRequestLists{
			Amount:            disb.Amount,
			BankCode:          disb.BankCode,
			BankAccountName:   disb.BankAccountName,
			BankAccountNumber: disb.BankAccountNumber,
			Description:       disb.Description,
			ExternalID:        helperXendit.BatchDisbursementPrefix + helperRandomStringXendit.GetRandomString(10),
			EmailTo:           disb.EmailTo,
			EmailCc:           disb.EmailCc,
			EmailBcc:          disb.EmailBcc,
		})
	}
	req.DisbursementBulkRequestLists = disbursementBulkRequestLists

	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)
	resp, err := clientService.XenditService.CreateBulkDisbursementHandler(req)
	if err != nil {
		level.Error(u.logger).Log("function", "XenditUsecaseImpl createDisbursementHandler", "error", err)
		return nil, err
	}

	CreateBulkDisbursementHandler, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("CreateBulkDisbursementHandler	=", string(CreateBulkDisbursementHandler))

	vmResHTTPJSON := string(CreateBulkDisbursementHandler)
	xenditResponse := vmXendit.XenditResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceXendit.SuccessCodeType,
		Message:  serviceXendit.MessageRetrievedSuccess,
		Data:     vmResHTTPJSON,
	}

	return &xenditResponse, nil
}

// Batch Disbursement callback
func (u *XenditUsecaseImpl) BulkDisbursementCallbackHandler(entity.BulkDisbursementCallBackRequest) (*vmXendit.XenditResponse, error) {

	vmResHTTPJSON := "result"
	xenditResponse := vmXendit.XenditResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceXendit.SuccessCodeType,
		Message:  serviceXendit.MessageRetrievedSuccess,
		Data:     vmResHTTPJSON,
	}

	return &xenditResponse, nil
}
