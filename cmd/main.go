package main

import (
	"attendance-system/internal/infra"
	"attendance-system/internal/router"
)

func main() {

	infra.NewMongoClient()
	defer infra.DisconnectMongo()

	e := router.NewRouter()
	e.Logger.Fatal(e.Start(":1323"))
}
