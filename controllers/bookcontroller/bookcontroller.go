package bookcontroller

import (
	"github.com/gofiber/fiber/v3"
)

func Index(c *fiber.Ctx) error{
	var book []models.Book
	models.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(books);
}
func Show(c *fiber.Ctx) error{
	//
}
func Store(c *fiber.Ctx) error{
	var book models.Book
	if err := c.BodyParser(&book); err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	if err := models.DB.Create(&book).Error;err !=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":err.Error(),
		})

}
func Update(c *fiber.Ctx) error{
	//
}
func Delete(c *fiber.Ctx) error{
	//
}
