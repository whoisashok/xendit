package service

import (
	entity "xendit/xendit_services/entity"
	vmXendit "xendit/xendit_services/view_model"
)

type XenditService interface {
	// disbursements
	CreateDisbursementHandler(entity.DisbursementRequestBody) (*vmXendit.XenditResponse, error)
	// Get disbursement by id
	GetDisbursementByIdHandler(entity.GetDisbursementByIdRequest) (*vmXendit.XenditResponse, error)
	// Get disbursement by external_id
	GetDisbursementByExternalIDHandler(entity.GetDisbursementByExternalIDRequest) (*vmXendit.XenditResponse, error)
	// Disbursement callback
	DisbursementCallbackHandler(entity.DisbursementCallBackRequest) (*vmXendit.XenditResponse, error)
	// Create batch disbursement
	CreateBulkDisbursementHandler(entity.DisbursementBulkRequest) (*vmXendit.XenditResponse, error)
	// Batch Disbursement callback
	BulkDisbursementCallbackHandler(entity.BulkDisbursementCallBackRequest) (*vmXendit.XenditResponse, error)
}
