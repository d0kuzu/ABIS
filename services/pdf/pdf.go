package pdf

import (
	. "ABIS/database/models"
	"ABIS/database/models/repos/book_repo"
	"errors"
	"fmt"
	"github.com/ledongthuc/pdf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func ParsePDF(filePath string) (Book, error) {
	info, err := GetInfo(filePath)
	if err != nil {
		return Book{}, err
	}

	isbn, err := CollectISBN(filePath)
	if err != nil && err.Error() != "isbn не найден" {
		return Book{}, err
	}

	isExist := book_repo.CheckIfExist(info.Title, filePath)

	if !isExist {
		book, err := book_repo.Create(info.Title, info.Author, isbn, filePath)
		if err != nil {
			return Book{}, err
		}

		return book, nil
	}

	return Book{}, errors.New("book already exists")
}

func GetInfo(filePath string) (*pdfcpu.PDFInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Не удалось открыть файл: " + err.Error())
	}
	defer file.Close()

	conf := model.NewDefaultConfiguration()

	info, err := api.PDFInfo(file, filePath, nil, conf)
	if err != nil {
		return nil, errors.New("pdfcpu.GetInfo: " + err.Error())
	}

	fmt.Println("Метаданные pdf:")
	fmt.Println("Title: ", info.Title)
	fmt.Println("Author: ", info.Author)
	fmt.Println("Subject: ", info.Subject)
	fmt.Println("Keywords: ", info.Keywords)
	fmt.Println("Creator: ", info.Creator)
	fmt.Println("Producer: ", info.Producer)
	fmt.Println("Creation Date: ", info.CreationDate)

	return info, nil
}

func CollectISBN(filePath string) (string, error) {
	f, r, err := pdf.Open(filePath)
	if err != nil {
		return "", errors.New("ошибка открытия PDF: %v" + err.Error())
	}
	defer f.Close()

	numPages := 20

	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page := r.Page(pageNum)

		pageText, err := page.GetPlainText(nil)

		if err != nil {
			log.Printf("ошибка извлечения текста с страницы %d: %v", pageNum, err)
			continue
		}
		isbn := FindISBN(pageText)
		if isbn != "" {
			return isbn, nil
		}
	}

	return "", errors.New("isbn не найден")
}

func FindISBN(text string) string {
	isbnPattern := `ISBN(?:-13)?:?\s*([\d-]+)`
	cleanDigits := regexp.MustCompile(`[^\d]`)

	isbnRegex := regexp.MustCompile(isbnPattern)
	match := isbnRegex.FindString(text)
	match = strings.ReplaceAll(match, "-", "")

	return cleanDigits.ReplaceAllString(match, "")
}
