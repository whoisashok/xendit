package viewmodel

type XenditResponse struct {
	Status   string      `json:"status"`
	Code     int32       `json:"code"`
	CodeType string      `json:"code_type"`
	Message  string      `json:"code_message"`
	Data     interface{} `json:"data"`
}
