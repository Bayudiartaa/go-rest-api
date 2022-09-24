package bookController

import (
	"github.com/Bayudiartaa/go-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)

	return c.JSON(fiber.Map{
		"message" : "Berhasil Menampilkan Data Buku",
		"data" : books,
	});
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message" : "Data Tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.JSON(book)
}

func Create(c *fiber.Ctx) error {
	var book models.Book
    if err := c.BodyParser(&book); err != nil {
	    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message" : "Data Berhasil Di tambahkan",
		"data" : book,
	})
}

func Update(c *fiber.Ctx) error {
	
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
	    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "Tidak Data Mengupdate Data",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "Data Berhasil Di Update",
	})
}

func Delete(c *fiber.Ctx) error {
	
	id := c.Params("id")

	var book models.Book
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "Data Tidak Data Di Hapus",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "Data Berhasil Di Hapus",
	})
}