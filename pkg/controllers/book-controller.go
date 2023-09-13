package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/udodinho/bookstore/pkg/models"
)

func CreateBook(context *fiber.Ctx) error {
	newBook := &models.Books{}

	err := context.BodyParser(&newBook)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
			return err
	}

	bk, err := newBook.CreateBook()

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Could not create book"})
			return err
	}

	context.Status(http.StatusCreated).JSON(&fiber.Map{
		"message":"Book has been created successfully",
		"data":bk,
	})
		
	return nil

}

func GetAllBooks(context *fiber.Ctx) error {
	bks, err := models.GetAllBooks()

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Could not fetch books"})
			return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Books fetched successfully",
		"data":bks,})

		return nil
}

func GetBook(context *fiber.Ctx) error {
	id := context.Params("id")

	bkID, err := strconv.Atoi(id)

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"Id does not exist"})
			return nil
	}

	bks, _, err := models.GetBookbyID(int64(bkID))
	
	if bkID != int(bks.ID) || int(bks.ID) < 1 {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "No book with id", "data":bkID})
			return err
		}
		
		if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Could not fetch book"})
			return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Book fetched successfully",
		"data":bks,})

		return nil
}

func UpdateBook(context *fiber.Ctx) error {
	updateBook := &models.Books{}

	err := context.BodyParser(&updateBook)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
			return err
	}

	id := context.Params("id")

	bkID, err := strconv.Atoi(id)

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"Id does not exist"})
			return nil
	}

	updatedBook, db, err := models.GetBookbyID(int64(bkID))

	if bkID != int(updatedBook.ID) || int(updatedBook.ID) < 1 {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "No book with id", "data":bkID})
			return err
		}

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Could not fetch book"})
			return err
	}

	if *updateBook.Author != "" {
		updatedBook.Author = updateBook.Author
	}

	if *updateBook.Title != "" {
		updatedBook.Title = updateBook.Title
	}

	if *updateBook.Publisher != "" {
		updatedBook.Publisher = updateBook.Publisher
	}

	db.Save(&updatedBook)

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Book updated successfully",
		"data":updatedBook,})

		return nil
}

func DeleteBook(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"Id cannot be empty", "data":id})
			return nil
	}

	bkID, err := strconv.Atoi(id)

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"Id does not exist"})
			return nil
	}
	
	bks, err := models.DeleteBook(int64(bkID))

	if bkID != int(bks.ID) || int(bks.ID) < 1 {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "No book with id", "data":bkID})
			return err
		}

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Could not delete book"})
			return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Book deleted successfully",
		"data":bks,})

		return nil
}
