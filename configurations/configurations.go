package configurations

// DARAJA ENVIRONMENT
// CAN EITHER BE SANDBOX OR LIVE

var enviroment = "SANDBOX"

// Daraja basic configurations

var daraja = map[string]map[string]string{
	"SANDBOX_CONFIGS": {
		"APP_CONSUMER_KEY":        "pyhfLWi17bMPs3gchEnEAY2wb6S9Wj9n",
		"APP_SECRET_KEY":          "VCORc4rrhPGP3SRj",
		"B2B_B2C_SHORTCODE":       "603034", // this is the shortcode to be used for c2b transactions
		"C2B_SHORTCODE":           "600000", // this is the shortcode to be used for b2b and b2c transactions
		"INITIATOR_NAME":          "apiop31",
		"TEST_MSISDN":             "254708374149",
		"SECURITY_CREDENTIAL":     "4bfTPxTk",
		"LIPA_NA_MPESA_SHORTCODE": "174379",
		"LIPA_NA_MPESA_PASSKEY":   "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
		"CERT_LOCATION":           "",
	},

	"LIVE_CONFIGS": {
		"APP_CONSUMER_KEY":        "",
		"APP_SECRET_KEY":          "",
		"B2C_SHORTCODE":           "",
		"C2B_SHORTCODE":           "",
		"INITIATOR_NAME":          "",
		"SECURITY_CREDENTIAL":     "",
		"LIPA_NA_MPESA_SHORTCODE": "",
		"LIPA_NA_MPESA_PASSKEY":   "",
		"PRODUCTION_CERT":         "",
	},
}

// daraja post endpoints

var endpoints = map[string]map[string]string{
	"SANDBOX_ENDPOINTS": {
		"AUTHENTICATION":    "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials",
		"SIMULATE_C2B":      baseEndpoint(enviroment) + "c2b/v1/simulate",
		"C2B_REGISTER_URLS": baseEndpoint(enviroment) + "c2b/v1/registerurl",
		"B2C":               baseEndpoint(enviroment) + "b2c/v1/paymentrequest",
		"B2B":               baseEndpoint(enviroment) + "b2b/v1/paymentrequest",
		"ACCOUNT_BALANCE":   baseEndpoint(enviroment) + "accountbalance/v1/query",
		"REVERSAL":          baseEndpoint(enviroment) + "reversal/v1/request",
		"CHECKOUT":          baseEndpoint(enviroment) + "stkpush/v1/processrequest",
	},
	"LIVE_ENDPOINTS": {
		"AUTHENTICATION":    "https://api.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials",
		"C2B_REGISTER_URLS": baseEndpoint(enviroment) + "c2b/v1/registerurl",
		"B2C":               baseEndpoint(enviroment) + "b2c/v1/paymentrequest",
		"B2B":               baseEndpoint(enviroment) + "b2b/v1/paymentrequest",
		"ACCOUNT_BALANCE":   baseEndpoint(enviroment) + "accountbalance/v1/query",
		"REVERSAL":          baseEndpoint(enviroment) + "reversal/v1/request",
		"CHECKOUT":          baseEndpoint(enviroment) + "stkpush/v1/processrequest",
	},
}

// daraja callback urls

var callbacks = map[string]string{
	"CONFIRMATION_URL":               "https://callmevincent.com/daraja/confirmation/",
	"VALIDATION_URL":                 "https://callmevincent.com/daraja/validation/",
	"B2C_RESULT_URL":                 "https://callmevincent.com/daraja/b2cnotifs/",
	"B2C_TIMEOUT_URL":                "https://callmevincent.com/daraja/timeout/",
	"B2B_RESULT_URL":                 "https://callmevincent.com/daraja/b2bnotifs/",
	"B2B_TIMEOUT_URL":                "https://callmevincent.com/daraja/b2btimeout/",
	"ACCOUNT_BALANCE_RESULT_URL":     "https://callmevincent.com/daraja/accountbalance/",
	"ACCOUNT_BALANCE_TIMEOUT_URL":    "https://callmevincent.com/daraja/balancetimeout/",
	"TRANSACTION_STATUS_RESULT_URL":  "https://callmevincent.com/daraja/",
	"TRANSACTION_STATUS_TIMEOUT_URL": "https://callmevincent.com/daraja/",
	"REVERSAL_RESULT_URL":            "https://callmevincent.com/daraja/reversal/",
	"REVERSAL_TIMEOUT_URL":           "https://callmevincent.com/daraja/reversaltimeout/",
	"MOBILE_CHECKOUT_URL":            "https://callmevincent.com/daraja/checkout/",
}
