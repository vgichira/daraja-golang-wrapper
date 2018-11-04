package main

import (
	"fmt"
	"hello/daraja-golang-wrapper/configurations"
)

func main() {
	fmt.Println(configurations.MobileCheckout("SANDBOX", "254725089232", "vgichira", "This is a test", 100))
}
