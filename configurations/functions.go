package configurations

// function to return the base path for endpoints

func baseEndpoint(env string) string {
	var endpoint string
	if env == "SANDBOX" {
		endpoint = "https://sandbox.safaricom.co.ke/mpesa/"
	} else if env == "LIVE" {
		endpoint = "https://sandbox.safaricom.co.ke/api/"
	}
	return endpoint
}
