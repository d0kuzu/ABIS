package pdf

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	_ "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func GetInfo() {
	filePath := "test3.pdf"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	conf := model.NewDefaultConfiguration()

	info, err := api.PDFInfo(file, filePath, nil, conf)
	if err != nil {
		log.Fatalf("Ошибка при извлечении метаданных: %v", err)
	}

	fmt.Println("Метаданные pdf:")
	fmt.Println("Title: ", info.Title)
	fmt.Println("Author: ", info.Author)
	fmt.Println("Subject: ", info.Subject)
	fmt.Println("Keywords: ", info.Keywords)
	fmt.Println("Creator: ", info.Creator)
	fmt.Println("Producer: ", info.Producer)
	fmt.Println("Creation Date: ", info.CreationDate)
}
