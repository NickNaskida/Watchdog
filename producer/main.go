package main

import (
	"fmt"
	"github.com/NickNaskida/Watchdog/producer/services"
)

func main() {
	fmt.Println("Starting the alert producer ...")
	randomAlert := services.NewAlert()

	fmt.Println("Alert Message: ", randomAlert.Message)
	fmt.Println("Alert Category: ", randomAlert.Category)
}
