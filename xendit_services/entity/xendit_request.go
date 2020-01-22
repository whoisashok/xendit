package entity

type DisbursementRequestBody struct {
	ExternalID              string   `json:"external_id"`
	Amount                  int64    `json:"amount"`
	BankCode                string   `json:"bank_code"`
	AccountHolderName       string   `json:"account_holder_name"`
	AccountNumber           string   `json:"account_number"`
	DisbursementDescription string   `json:"description"`
	EmailTo                 []string `json:"email_to,omitempty"`
	EmailCc                 []string `json:"email_cc,omitempty"`
	EmailBcc                []string `json:"email_bcc,omitempty"`
}

type GetDisbursementByIdRequest struct {
	DisbursementID string `json:"disbursement_id"`
}

type GetDisbursementByExternalIDRequest struct {
	ExternalID string `json:"external_id"`
}

type DisbursementCallBackRequest struct {
	ID                      string   `json:"id"`
	UserID                  string   `json:"user_id"`
	ExternalID              string   `json:"external_id"`
	Amount                  string   `json:"amount"`
	BankCode                string   `json:"bank_code"`
	AccountHolderName       string   `json:"account_holder_name"`
	DisbursementDescription string   `json:"description"`
	FailureCode             string   `json:"failure_code"`
	IsInstant               string   `json:"is_instant"`
	Status                  string   `json:"status"`
	Updated                 string   `json:"updated"`
	Created                 string   `json:"created"`
	EmailTo                 []string `json:"email_to,omitempty"`
	EmailCc                 []string `json:"email_cc,omitempty"`
	EmailBcc                []string `json:"email_bcc,omitempty"`
}

type DisbursementBulkRequest struct {
	Reference                    string                         `json:"reference"`
	DisbursementBulkRequestLists []DisbursementBulkRequestLists `json:"disbursements"`
}

type DisbursementBulkRequestLists struct {
	Amount            int64    `json:"amount"`
	BankCode          string   `json:"bank_code"`
	BankAccountName   string   `json:"bank_account_name"`
	BankAccountNumber string   `json:"bank_account_number"`
	Description       string   `json:"description"`
	ExternalID        string   `json:"external_id"`
	EmailTo           []string `json:"email_to,omitempty"`
	EmailCc           []string `json:"email_cc,omitempty"`
	EmailBcc          []string `json:"email_bcc,omitempty"`
}

type BulkDisbursementCallBackRequest struct {
	Created             string `json:"created"`
	Reference           string `json:"reference"`
	TotalUploadedAmount int64  `json:"total_uploaded_amount"`
	TotalUploadedCount  string `json:"total_uploaded_count"`
	Status              string `json:"status"`
	ID                  string `json:"id"`
}
