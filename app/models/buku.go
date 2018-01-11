package models

type Buku struct {
	Id        int    `gorm:"primary_key;auto_increment;not_null"`
	KdBuku    string `gorm:"size:10"`
	Judul     string `gorm:"size:50"`
	Pengarang string `gorm:"size:20"`
}
