package models

import (
	"time"
	"github.com/revel/revel"
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

func (c *Alternatif) Validate(v *revel.Validation) {
	v.Required(c.NamaAlternatif).Message("Nama Alternatif is Required")
	v.MinSize(c.NamaAlternatif, 2).Message("\"" + c.NamaAlternatif + "\" is not Long Enough")
}