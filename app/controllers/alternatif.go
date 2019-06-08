package controllers

import (
	"strings"
	"github.com/revel/revel"
	"myapp/app/models"
	"errors"
	// "strconv"
	// "fmt"
	// "encoding/json"
)

type Alternatif struct {
	GormController
}

func (c Alternatif) Index() revel.Result {
	alternatifs := []models.Alternatif{}

	result := c.Txn.Find(&alternatifs)
	err := result.Error

	if err != nil {
		return c.RenderError(errors.New("Data Alternatif Not Found"))
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

	alternatif.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Alternatif.Add)
	}

	c.Txn.NewRecord(alternatif)
	ret := c.Txn.Create(&alternatif)

	if ret.Error != nil {
		return c.RenderError(errors.New("Record Failed to Create. " + ret.Error.Error()))
	}

	return c.Redirect(Alternatif.Index)
}

func (c Alternatif) Edit(id int) revel.Result {
	alternatif := []models.Alternatif{}

	result := c.Txn.First(&alternatif, id)
	err := result.Error

	if err != nil {
		return c.RenderError(errors.New("Data Alternatif Not Found"))
	}

	return c.Render(alternatif)
}

func (c Alternatif) Update() revel.Result {
	// id, err := strconv.Atoi(c.Params.Form.Get("id"))
	// if err != nil {
	// 	return c.Redirect(Alternatif.Index)
	// }

	// _, err = c.Txn.ExecUpdate(c.Gdb.SqlStatementBuilder.
	// 	Update("tbl_alternatif").Set("nama_alternatif", c.Params.Form.Get("nama_alternatif")).
	// 	Where("id=?", id))
	// if err != nil {
	// 	panic(err)
	// }

	// postData := models.Alternatif {
	// 	Id: id,
	// 	NamaAlternatif: strings.TrimSpace(c.Params.Form.Get("nama_alternatif")),
	// 	Keterangan: strings.TrimSpace(c.Params.Form.Get("keterangan")),
	// }
	// c.Txn.First(&alternatif, id)
	// alternatif.NamaAlternatif = strings.TrimSpace(c.Params.Form.Get("nama_alternatif"))
	// alternatif.Keterangan = strings.TrimSpace(c.Params.Form.Get("keterangan"))

	// alternatif.Validate(c.Validation)

	// if c.Validation.HasErrors() {
	// 	c.Validation.Keep()
	// 	c.FlashParams()
	// 	return c.Redirect("alternatif/")
	// }

	// ret := c.Txn.Save(&alternatif)

	// if ret.Error != nil {
	// 	return c.RenderError(errors.New("Record Failed to Update. " + ret.Error.Error()))
	// }

	return c.Redirect(Alternatif.Index)
}