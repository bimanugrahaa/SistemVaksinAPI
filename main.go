package main

import (
	"SistemVaksinAPI/config"
	"SistemVaksinAPI/migrate"
	"SistemVaksinAPI/routes"
)

func main() {
	config.InitDB()
	migrate.AutoMigrate()

	e := routes.New()

	e.Start(":8000")
}
