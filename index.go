package main

import (
	"fmt"
	"hello/daraja-golang-wrapper/configurations"
)

func main() {
	fmt.Println(configurations.B2BRequest("SANDBOX", "BusinessPayBill", "600000", "test", "This is a test", 4, 4, 100))
}
