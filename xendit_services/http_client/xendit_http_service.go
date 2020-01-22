package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	entity "xendit/xendit_services/entity"
)

// methods of the HTTP REST API.
type XenditService struct {
	client *ClientService
}

// Create Disbursement Request
func (s *XenditService) CreateDisbursementHandler(opt entity.DisbursementRequestBody) (*http.Response, error) {
	// build the URL
	u, err := url.Parse(config.GetString(`xendit.base_url`) + config.GetString(`xendit.disbursements`))
	if err != nil {
		return nil, err
	}

	json_body, _ := json.Marshal(opt)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(json_body))
	if err != nil {
		panic(err)
	}
	//fmt.Println(req)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-IDEMPOTENCY-KEY", time.Now().Format("2006-01-02 15:04:05"))
	req.Header.Add("Authorization", config.GetString(`xendit_keys.Authorization`))

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Get disbursement by id
func (s *XenditService) GetDisbursementByIdHandler(opt entity.GetDisbursementByIdRequest) (*http.Response, error) {
	// build the URL
	u, err := url.Parse(config.GetString(`xendit.base_url`) + config.GetString(`xendit.disbursements`) + "/" + opt.DisbursementID)
	if err != nil {
		return nil, err
	}

	json_body, _ := json.Marshal(opt)
	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(json_body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", config.GetString(`xendit_keys.Authorization`))

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Get disbursement by external_id
func (s *XenditService) GetDisbursementByExternalIDHandler(opt entity.GetDisbursementByExternalIDRequest) (*http.Response, error) {
	// build the URL
	u, err := url.Parse(config.GetString(`xendit.base_url`) + config.GetString(`xendit.disbursements`) + "?external_id=" + opt.ExternalID)
	if err != nil {
		return nil, err
	}

	json_body, _ := json.Marshal(opt)

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer(json_body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", config.GetString(`xendit_keys.Authorization`))

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Create batch disbursement
func (s *XenditService) CreateBulkDisbursementHandler(opt entity.DisbursementBulkRequest) (*http.Response, error) {
	// build the URL
	u, err := url.Parse(config.GetString(`xendit.base_url`) + config.GetString(`xendit.batch_disbursements`))
	if err != nil {
		return nil, err
	}

	json_body, _ := json.Marshal(opt)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(json_body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-IDEMPOTENCY-KEY", time.Now().Format("2006-01-02 15:04:05"))
	req.Header.Add("Authorization", config.GetString(`xendit_keys.Authorization`))

	fmt.Println(req)

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}
