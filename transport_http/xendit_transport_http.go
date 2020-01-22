package transporthttp

import (
	"net/http"

	"context"

	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	endpointXendit "xendit/endpoint"
	modelHttpXendit "xendit/model_http"
)

// MakeExtMemberHttpHandler - make http handler
func MakeXenditHttpHandler(ctx context.Context, endpointXendit endpointXendit.XenditEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(modelHttpXendit.EncodeError),
	}

	apiV1 := r.PathPrefix("/api/v1").Subrouter()

	//1. POST - Create disbursement
	apiV1.Methods("POST").Path("/disbursements").Handler(httpTransport.NewServer(
		endpointXendit.CreateDisbursementEndpoint,
		modelHttpXendit.DecodeCreateDisbursementRequest,
		modelHttpXendit.EncodeResponse,
		options...,
	))

	//2. GET - Get disbursement by id
	apiV1.Methods("GET").Path("/disbursements/GetDisbursementById").Handler(httpTransport.NewServer(
		endpointXendit.GetDisbursementByIdEndpoint,
		modelHttpXendit.DecodeGetDisbursementByIdRequest,
		modelHttpXendit.EncodeResponse,
		options...,
	))

	//3. GET - Get disbursement by external_id
	apiV1.Methods("GET").Path("/disbursements/GetDisbursementByExternalID").Handler(httpTransport.NewServer(
		endpointXendit.GetDisbursementByExternalIDEndpoint,
		modelHttpXendit.DecodeGetDisbursementByExternalIDRequest,
		modelHttpXendit.EncodeResponse,
		options...,
	))

	//4. POST - Disbursement callback
	apiV1.Methods("POST").Path("/disbursement_callback_url").Handler(httpTransport.NewServer(
		endpointXendit.DisbursementCallbackEndpoint,
		modelHttpXendit.DecodeDisbursementCallbackRequest,
		modelHttpXendit.EncodeResponse,
		options...,
	))

	//5. POST - Create batch disbursement
	apiV1.Methods("POST").Path("/batch_disbursements").Handler(httpTransport.NewServer(
		endpointXendit.CreateBulkDisbursementEndpoint,
		modelHttpXendit.DecodeCreateBulkDisbursementRequest,
		modelHttpXendit.EncodeResponse,
		options...,
	))

	//6. POST - Batch Disbursement callback
	apiV1.Methods("POST").Path("/batch_disbursement_callback_url").Handler(httpTransport.NewServer(
		endpointXendit.BulkDisbursementCallbackEndpoint,
		modelHttpXendit.DecodeBulkDisbursementCallbackRequest,
		modelHttpXendit.EncodeResponse,
		options...,
	))

	return apiV1
}
