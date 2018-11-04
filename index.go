package main

import (
	"fmt"
	"hello/daraja-golang-wrapper/configurations"
)

func main() {
	fmt.Println(configurations.SimulateC2B("CustomerBuyGoodsOnline", "254708374149", "vgichira", 10000))
}
