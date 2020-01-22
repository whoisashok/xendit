package endpoint

import (
	"context"

	serviceXendit "xendit/service"
	entity "xendit/xendit_services/entity"

	endpointGRPC "github.com/go-kit/kit/endpoint"
)

// Create Disbursement Request
func MakeCreateDisbursementEndpoint(s serviceXendit.XenditService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.DisbursementRequestBody)
		r, err := s.CreateDisbursementHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

// Get disbursement by id
func MakeGetDisbursementByIdEndpoint(s serviceXendit.XenditService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.GetDisbursementByIdRequest)
		r, err := s.GetDisbursementByIdHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

// Get disbursement by external_id
func MakeGetDisbursementByExternalIDEndpoint(s serviceXendit.XenditService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.GetDisbursementByExternalIDRequest)
		r, err := s.GetDisbursementByExternalIDHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

// Disbursement callback
func MakeDisbursementCallbackEndpoint(s serviceXendit.XenditService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.DisbursementCallBackRequest)
		r, err := s.DisbursementCallbackHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

// Create batch disbursement
func MakeCreateBulkDisbursementEndpoint(s serviceXendit.XenditService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.DisbursementBulkRequest)
		r, err := s.CreateBulkDisbursementHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

// Batch Disbursement callback
func MakeBulkDisbursementCallbackEndpoint(s serviceXendit.XenditService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.BulkDisbursementCallBackRequest)
		r, err := s.BulkDisbursementCallbackHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
