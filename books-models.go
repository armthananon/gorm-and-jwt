package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getBook(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(book)
}

func getBooks(db *gorm.DB) []Book {
	var book []Book
	result := db.Find(&book)

	if result.Error != nil {
		log.Fatalf("Unable to find book: %v", result.Error)
	}

	return book
}

func updateBook(db *gorm.DB, book *Book) error {
	result := db.Save(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func deleteBook(db *gorm.DB, id uint) error {
	var book Book
	result := db.Delete(&book, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
