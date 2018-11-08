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

type B2CRequestStruct struct {
	InitiatorName      string
	SecurityCredential string
	CommandID          string
	Amount             int
	PartyA             string
	PartyB             string
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
	Occasion           string
}

type B2BRequestStruct struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	SenderIdentifierType   int
	RecieverIdentifierType int
	Amount                 int
	PartyA                 string
	PartyB                 string
	AccountReference       string
	Remarks                string
	QueueTimeOutURL        string
	ResultURL              string
}
