package main

import (
	"log"
	"skn-go-mail/src/pkgs/routes"
	"skn-go-mail/src/pkgs/utils"
)

func main() {
	routes.GetRoutes()

	var base string = utils.ViperEnvVariable("ROUTES.MAIL.BASIC")

	log.Fatalf("Base: %s", base)

}
