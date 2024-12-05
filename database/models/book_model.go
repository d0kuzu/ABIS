package models

import (
	"time"
)

type Book struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"` //ID_BOOK
	Title            string    `gorm:"size:255;not null"`        //5.1
	Description      string    `gorm:"size:1000"`                //
	Language         string    `gorm:"default:'EN'"`             //44.1
	Photo            string    `gorm:"size:255"`
	BookPDFFile      string    `gorm:"size:255"` //module
	Barcode          string    `gorm:"size:8"`
	BarcodeImagePath string    `gorm:"size:255"`
	FirstDate        time.Time //FirstDate
	LastDate         time.Time //LastDate
	RecDate          time.Time //RecDate
	ISBN             string    `gorm:"size:255"`      //module
	Format           string    `gorm:"default:'PDF'"` //
	NumOfCopies      int       //CopyCnt
	DocumentCode     string    `gorm:"size:255"`
	Keywords         []string  `gorm:"type:text"`
	Categories       []string  `gorm:"type:text"`
	Marcode          string    `gorm:"size:255"`
	Udk              string    `gorm:"size:255"`
	Bbk              string    `gorm:"size:255"`
	MarcStandard     string    `gorm:"size:20"`
	MarcRecordXml    string    `gorm:"type:text"`
	MarcFilePath     string    `gorm:"size:255"`
}
