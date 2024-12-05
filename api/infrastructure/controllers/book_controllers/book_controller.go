package book_controllers

import (
	. "ABIS/api/infrastructure/response_models"
	. "ABIS/config"
	"ABIS/services/pdf"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBook(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "файл обязателен"})
		return
	}

	filePath := TempPath + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось сохранить файл"})
		return
	}

	book, err := pdf.ParsePDF(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)

	//if info.Title != "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не содержит метаданных", "isbn": isbn})
	//	return
	//}
	//
	////err = book_repo.UpdateData(info, isbn, filePath)
	//book, err := book_repo.GetByTitle(info.Title)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//book.Title = info.Title
	//book.ISBN = isbn
	//book.BookPDFFile = filePath
	//
	//err = book_repo.Update(book)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
}

func CreateBooks(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form"})
		return
	}

	files := form.File["file"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no files uploaded"})
		return
	}

	filesUploaded := map[string][]FileStatus{
		"uploaded":       {},
		"already_exists": {},
		"error":          {},
	}

	for _, file := range files {
		filePath := TempPath + file.Filename
		if err = c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to save file: %s", file.Filename)})
			return
		}

		_, err = pdf.ParsePDF(filePath)
		newFile := FileStatus{
			Path: filePath,
		}
		if err != nil {
			newFile.Status = err.Error()

			if err.Error() == "book already exists" {
				filesUploaded["already_exists"] = append(filesUploaded["already_exists"], newFile)
				continue
			}
			filesUploaded["error"] = append(filesUploaded["error"], newFile)
			continue
		}
		filesUploaded["uploaded"] = append(filesUploaded["uploaded"], newFile)
	}

	c.JSON(http.StatusOK, filesUploaded)
}
