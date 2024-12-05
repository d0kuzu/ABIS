package book_repo

import (
	"ABIS/database"
	. "ABIS/database/models"
	"errors"
)

func Create(title string, author string, isbn string, docPath string) (Book, error) {
	db := database.GetDB()

	var book Book
	book.Title = title
	book.ISBN = isbn
	book.BookPDFFile = docPath

	if err := db.Create(&book).Error; err != nil {
		return Book{}, errors.New("невозможно создать новый объект")
	}

	return book, nil
}

func CheckIfExist(title string, docPath string) bool {
	db := database.GetDB()
	var book Book
	field := "title = ?"
	value := title

	if title == "" {
		field = "book_pdf_file = ?"
		value = docPath
	}
	if err := db.Where(field, value).First(&book).Error; err != nil {
		return false
	}

	return true
}

func GetByTitle(title string) (Book, error) {
	db := database.GetDB()

	var book Book
	if err := db.First(&book, "Title = ?", title).Error; err != nil {
		return Book{}, errors.New("невозможно найти объект")
	}

	return book, nil
}

func Update(book Book) error {
	db := database.GetDB()

	if err := db.Save(&book).Error; err != nil {
		return errors.New("невозможно сохранить объект")
	}

	return nil
}
