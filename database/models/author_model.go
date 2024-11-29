package models

type Author struct {
	ID    uint    `gorm:"primaryKey;autoIncrement"`
	Name  string  `gorm:"size:255;not null"`
	Books []*Book `gorm:"many2many:book_authors;"`
}
