package main

import (
	"ABIS/services/parser"
)

func main() {
	//config.LoadENV()

	//database.Connect()
	//defer database.Disconnect()
	//
	//api.RouterStart()

	parser.Read("static/pdf/test.pdf")
}
