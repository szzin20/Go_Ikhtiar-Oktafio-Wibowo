package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	//ID_user string `json:"id_user"`
	Judul  string `json:"judul"`
	Konten string `json:"konten"`
}
