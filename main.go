package main

import (
	"SistemVaksinAPI/config"
	"SistemVaksinAPI/migrate"
	"SistemVaksinAPI/routes"
	"fmt"
)

func main() {
	config.InitDB()
	migrate.AutoMigrate()
	fmt.Println("")

	e := routes.New()

	e.Start(":8000")
}
