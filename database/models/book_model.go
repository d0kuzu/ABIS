package models

import (
	"time"
)

type Book struct {
	ID               uint   `gorm:"primaryKey;autoIncrement"`
	Title            string `gorm:"size:255;not null"`
	Description      string `gorm:"size:1000"`
	Language         string `gorm:"type:language;default:'EN'"`
	Photo            string `gorm:"size:255"`
	BookPDFFile      string `gorm:"size:255"`
	Barcode          string `gorm:"size:8"`
	BarcodeImagePath string `gorm:"size:255"`
	PublishDate      time.Time
	ISBN             string `gorm:"size:255"`
	Format           string `gorm:"type:format;default:'PDF'"`
	NumOfCopies      int
	DocumentCode     string   `gorm:"size:255"`
	Keywords         []string `gorm:"type:text"`
	Categories       []string `gorm:"type:text"`
	Marcode          string   `gorm:"size:255"`
	Udk              string   `gorm:"size:255"`
	Bbk              string   `gorm:"size:255"`
	MarcStandard     string   `gorm:"size:20"`
	MarcRecordXml    string   `gorm:"type:text"`
	MarcFilePath     string   `gorm:"size:255"`
}
