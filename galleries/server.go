package main

import (
	"galleries/packages/database"
	"galleries/routes"
)

func main() {
	database.Init()
	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))
}
