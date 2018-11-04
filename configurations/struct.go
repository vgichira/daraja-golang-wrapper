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
