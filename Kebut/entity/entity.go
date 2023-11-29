package entity

import "gorm.io/gorm"

type Penyakit struct {
	gorm.Model
	Nama       string `gorm:"type:varchar(20); NOT NULL" json:"nama"`
	Gejala     string `gorm:"type:string; not null" json:"gejala"`
	Penanganan string `gorm:"string; not null" json:"penanganan"`
	Links      []Link
}

type Link struct {
	gorm.Model
	Link       string `gorm:"type:string; not null" json:"link"`
	PenyakitId uint
}
