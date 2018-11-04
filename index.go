package main

import (
	"fmt"
	"hello/daraja-golang-wrapper/configurations"
)

func main() {
	fmt.Println(configurations.BaseEndpoint(configurations.Enviroment))
}
