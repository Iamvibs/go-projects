package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/iamvibs/golang-projects/fiber-crm/database"
	"github.com/iamvibs/golang-projects/fiber-crm/lead"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/lead", lead.GetLeads)
	app.Get("/api/lead/:id", lead.GetLead)
	app.Post("/api/lead/", lead.NewLead)
	app.Delete("/api/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to db")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen("3000")
	defer database.DBconn.Close()
}
