package main

import (
	"SistemVaksinAPI/config"
	"SistemVaksinAPI/routes"
	//"SistemVaksinasiAPI/routes"
)

func main() {
	config.InitDB()
	//migrate.AutoMigrate()

	e := routes.New()

	e.Start(":8000")
}
