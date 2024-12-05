package main

import (
	"ABIS/api"
	"ABIS/config"
	"ABIS/database"
)

func main() {
	config.LoadENV()

	database.Connect()
	defer database.Disconnect()

	api.RouterStart()

	//parser.GetDataFromTable("./data/excel/test.xlsx")
}
