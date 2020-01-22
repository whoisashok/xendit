package service

import (
	"time"
	entity "xendit/xendit_services/entity"
	usecase "xendit/xendit_services/usecase"
	vmXendit "xendit/xendit_services/view_model"

	raven "github.com/getsentry/raven-go"
	logging "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
)

type XenditServiceImpl struct {
	usecaseXendit usecase.XenditUsecase
	logger        logging.Logger
}

func NewXenditServiceImpl(usecaseXendit usecase.XenditUsecase, logger logging.Logger) XenditService {
	return &XenditServiceImpl{usecaseXendit, logger}
}

// Create Disbursement Request
func (s XenditServiceImpl) CreateDisbursementHandler(req entity.DisbursementRequestBody) (*vmXendit.XenditResponse, error) {
	level.Info(s.logger).Log("function", "NewXenditServiceImpl CreateDisbursementHandler", "result", "Entry")

	res, err := s.usecaseXendit.CreateDisbursementHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewXenditServiceImpl CreateDisbursementHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewXenditServiceImpl CreateDisbursementHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewXenditServiceImpl CreateDisbursementHandler", "result", "Exit")
	return res, nil
}

// Get disbursement by id
func (s XenditServiceImpl) GetDisbursementByIdHandler(req entity.GetDisbursementByIdRequest) (*vmXendit.XenditResponse, error) {
	level.Info(s.logger).Log("function", "NewXenditServiceImpl CreateDisbursementHandler", "result", "Entry")

	res, err := s.usecaseXendit.GetDisbursementByIdHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementById", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewXenditServiceImpl GetDisbursementById",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementById", "result", "Exit")
	return res, nil
}

// Get disbursement by external_id
func (s XenditServiceImpl) GetDisbursementByExternalIDHandler(req entity.GetDisbursementByExternalIDRequest) (*vmXendit.XenditResponse, error) {
	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "result", "Entry")

	res, err := s.usecaseXendit.GetDisbursementByExternalIDHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "result", "Exit")
	return res, nil
}

// Disbursement callback
func (s XenditServiceImpl) DisbursementCallbackHandler(req entity.DisbursementCallBackRequest) (*vmXendit.XenditResponse, error) {
	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "result", "Entry")

	res, err := s.usecaseXendit.DisbursementCallbackHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "result", "Exit")
	return res, nil
}

// Create batch disbursement
func (s XenditServiceImpl) CreateBulkDisbursementHandler(req entity.DisbursementBulkRequest) (*vmXendit.XenditResponse, error) {
	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "result", "Entry")

	res, err := s.usecaseXendit.CreateBulkDisbursementHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "result", "Exit")
	return res, nil
}

// Batch Disbursement callback
func (s XenditServiceImpl) BulkDisbursementCallbackHandler(req entity.BulkDisbursementCallBackRequest) (*vmXendit.XenditResponse, error) {
	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "result", "Entry")

	res, err := s.usecaseXendit.BulkDisbursementCallbackHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewXenditServiceImpl GetDisbursementByExternalIDHandler", "result", "Exit")
	return res, nil
}
