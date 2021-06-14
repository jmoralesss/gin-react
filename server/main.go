package main

import (
	_ "morales-backend-1/database"
	"morales-backend-1/router"
)

func main() {
	router := router.CreateRouter()

	router.RunTLS(":443", "./.cert/cert.pem", "./.cert/key.pem")
}
