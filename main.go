package main

import (
	"platform/placeholder"
	"platform/services"
)

func main() {
	services.RegisterDefaultServices()
	placeholder.Start()
}
