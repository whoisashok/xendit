package modelhttp

import (
	"context"
	"encoding/json"
	"net/http"
	entity "xendit/xendit_services/entity"
	helperInputValidateFaspay "xendit/xendit_services/helper/input_validate"
)

// Create Disbursement Request
func DecodeCreateDisbursementRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t entity.DisbursementRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}
	return t, nil
}

// Get disbursement by id
func DecodeGetDisbursementByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t entity.GetDisbursementByIdRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}
	return t, nil
}

// Get disbursement by external_id
func DecodeGetDisbursementByExternalIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t entity.GetDisbursementByExternalIDRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}
	return t, nil
}

// Disbursement callback
func DecodeDisbursementCallbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t entity.DisbursementCallBackRequest
	/*decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}*/
	return t, nil
}

// Create batch disbursement
func DecodeCreateBulkDisbursementRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t entity.DisbursementBulkRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}
	return t, nil
}

// Batch Disbursement callback
func DecodeBulkDisbursementCallbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t entity.BulkDisbursementCallBackRequest
	/*decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}*/
	return t, nil
}
