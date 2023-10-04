package lead

import (
	"github.com/gofiber/fiber"
	"github.com/iamvibs/golang-projects/fiber-crm/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBconn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBconn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBconn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	db := database.DBconn
	id := c.Params("id")
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500)
		return
	}
	db.Delete(&lead)
}
