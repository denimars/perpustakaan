package controllers

import (
	"log"
	"perpustakaan/app"
	"perpustakaan/app/models"
	"perpustakaan/app/routes"

	"github.com/revel/revel"
)

type Buku struct {
	App
}

func (c Buku) Index() revel.Result {
	var buku []models.Buku
	app.DB.Find(&buku)
	return c.Render(buku)
}

func (c Buku) Simpan(kdbuku string, judul string, pengarang string) revel.Result {
	var buku models.Buku
	buku.KdBuku = kdbuku
	buku.Judul = judul
	buku.Pengarang = pengarang
	err := app.DB.Save(&buku)
	if err.Error != nil {
		log.Println("________________")
		panic(err.Error)
	}
	return c.Redirect(routes.Buku.Index())
}
