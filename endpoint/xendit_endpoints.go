package endpoint

import (
	endpointHttp "github.com/go-kit/kit/endpoint"
)

type XenditEndpoints struct {
	CreateDisbursementEndpoint          endpointHttp.Endpoint
	GetDisbursementByIdEndpoint         endpointHttp.Endpoint
	GetDisbursementByExternalIDEndpoint endpointHttp.Endpoint
	DisbursementCallbackEndpoint        endpointHttp.Endpoint
	CreateBulkDisbursementEndpoint      endpointHttp.Endpoint
	BulkDisbursementCallbackEndpoint    endpointHttp.Endpoint
}
