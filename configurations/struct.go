package configurations

type RegisterUrl struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}

type simulateC2BStruct struct {
	ShortCode   string
	CommandId   string
	PhoneNumber string
	Amount      int
	AccountNo   string
}

type MobileCheckoutRequest struct {
	BusinessShortCode string
	Password          string
	Timestamp         string
	TransactionType   string
	Amount            int
	PartyA            string
	PartyB            string
	PhoneNumber       string
	CallBackURL       string
	AccountReference  string
	TransactionDesc   string
}
