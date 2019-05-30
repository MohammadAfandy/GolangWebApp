package models

import (
	"time"
)

type Alternatif struct {
	Id int `gorm:"primary_key" json:"id"`
	NamaAlternatif string `sql:"size:100" json:"nama_alternatif"`
	Keterangan string `sql:"size:255" json:"keterangan"`
	IdSpk int `json:"id_spk"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

func (Alternatif) TableName() string {
	return "tbl_alternatif"
}