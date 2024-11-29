package parser

import (
	"fmt"
	"log"
	"regexp"

	"github.com/ledongthuc/pdf"
)

func Read(filePath string) {
	f, r, err := pdf.Open(filePath)
	if err != nil {
		log.Fatalf("Ошибка открытия PDF: %v", err)
	}
	defer f.Close()

	numPages := 10

	var fullText string
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page := r.Page(pageNum)

		pageText, err := page.GetPlainText(nil)
		if err != nil {
			log.Printf("Ошибка извлечения текста с страницы %d: %v", pageNum, err)
			continue
		}
		isbn := FindISBN(pageText)
		if isbn != "" {
			fmt.Println(isbn)
			break
		}
	}

	fmt.Println("Извлечённый текст с первых страниц:")
	fmt.Println(fullText)
}

func FindISBN(text string) string {
	isbnPattern := `(?:ISBN(?:-13)?:?\s*(\d{13})|(\d{10}))`
	isbnRegex := regexp.MustCompile(isbnPattern)

	return isbnRegex.FindString(text)
}
