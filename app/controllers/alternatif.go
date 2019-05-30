package controllers

import (
	"strings"
	"github.com/revel/revel"
	"myapp/app/models"
	"errors"
)

type Alternatif struct {
	*revel.Controller
}

func (c Alternatif) Index() revel.Result {
	alternatifs := []models.Alternatif{}

	result := DB.Find(&alternatifs)
	err := result.Error

	if err != nil {
		return c.RenderError(errors.New("Data Alternatif Tidak Ditemukan"))
	}

	return c.Render(alternatifs)
}

func (c Alternatif) Add() revel.Result {

	return c.Render()
}

func (c Alternatif) Create() revel.Result {
	alternatif := models.Alternatif {
		NamaAlternatif: strings.TrimSpace(c.Params.Form.Get("nama_alternatif")),
		Keterangan: strings.TrimSpace(c.Params.Form.Get("keterangan")),
	}
	c.Validation.Required(alternatif.NamaAlternatif).Message("Nama Alternatif is Required")
	c.Validation.MinSize(alternatif.NamaAlternatif, 2).Message("\"" + alternatif.NamaAlternatif + "\" is not Long Enough")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Alternatif.Add)
	}

	ret := DB.Create(&alternatif)

	if ret.Error != nil {
		return c.RenderError(errors.New("Record Failed to Create. " + ret.Error.Error()))
	}

	return c.Redirect(Alternatif.Index)
}