package parser

import (
	"ABIS/database"
	"ABIS/database/models"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"time"
)

var fields = map[string]int{
	"NN":           0,
	"ID_BOOK":      1,
	"Bk_GUID":      2,
	"ID_Catalog":   3,
	"N_Label":      4,
	"ID_Librarian": 5,
	"FirstDate":    6,
	"LastDate":     7,
	"RecDate":      8,
	"CopyCnt":      9,
	"42.1":         10,
	"44.1":         11,
	"5.1":          12,
	"5.3":          13,
}

func GetDataFromTable(path string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer f.Close()

	sheetName := "Sheet1"
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Ошибка чтения строк: %v", err)
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 62 {
			row = append(row, make([]string, 62-len(row))...)
		}
		book := models.Book{}

		//uint64Value, err := strconv.ParseUint(row[fields["ID_BOOK"]], 10, 0)
		//if err != nil {
		//	fmt.Println("Error converting string to uint:", err)
		//	return
		//}
		//book.ID = uint(uint64Value)

		int64Value, err := strconv.ParseInt(row[fields["ID_BOOK"]], 10, 0)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		book.NumOfCopies = int(int64Value)

		book.Title = row[fields["5.1"]]
		book.Language = row[fields["44.1"]]

		layout := "02.01.2006"

		book.FirstDate, _ = time.Parse(row[fields["FirstDate"]], layout)
		book.LastDate, _ = time.Parse(row[fields["LastDate"]], layout)
		book.RecDate, _ = time.Parse(row[fields["RecDate"]], layout)

		db := database.GetDB()

		if err := db.Create(&book).Error; err != nil {
			fmt.Println("Could not create object", err)
			return
		}
	}
}
