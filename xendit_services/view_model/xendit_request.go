package viewmodel

type EmptyVM struct {
}

type DisbursementRequest struct {
	ExternalID              string
	Amount                  int64
	BankCode                string
	AccountHolderName       string
	AccountNumber           string
	DisbursementDescription string
	EmailTo                 string
	EmailCc                 string
	EmailBcc                string
}
