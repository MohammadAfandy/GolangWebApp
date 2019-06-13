package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
	"errors"
)

type ApiAlternatif struct {
	GormController
}

func (c ApiAlternatif) Index() revel.Result {
	alternatifs := []models.Alternatif{}

	result := c.Txn.Find(&alternatifs)
	err := result.Error

	if err != nil {
		return c.RenderError(errors.New("Data Alternatif Not Found"))
	}

	return c.RenderJSON(alternatifs)
}

func (c ApiAlternatif) Create() revel.Result {
	res := make(map[string]interface{})
	alternatif := models.Alternatif{}
    c.Params.BindJSON(&alternatif)
	alternatif.Validate(c.Validation)

	if c.Validation.HasErrors() {
		res["status"] = 422
		res["error"] = c.Validation.Errors
        return c.RenderJSON(res)
	}

	c.Txn.NewRecord(alternatif)
	ret := c.Txn.Create(&alternatif)

	if ret.Error != nil {
		res["status"] = 409
		res["error"] = "Record Failed to Create. " + ret.Error.Error()
		return c.RenderJSON(res)
	}
	
	res["status"] = 201
	res["message"] = "Success"

	return c.RenderJSON(res)
}

func (c ApiAlternatif) Update(id int) revel.Result {
	res := make(map[string]interface{})
	if err := c.Txn.First(&models.Alternatif{}, id).Error; err != nil {
		res["status"] = 404
		res["error"] = "Resource Not Found"
		return c.RenderJSON(res)
	}

	alternatif := models.Alternatif{}
    c.Params.BindJSON(&alternatif)
	alternatif.Validate(c.Validation)

	if c.Validation.HasErrors() {
		res["status"] = 422
		res["error"] = c.Validation.Errors
        return c.RenderJSON(res)
	}

	ret := c.Txn.Save(&alternatif)

	if ret.Error != nil {
		res["status"] = 409
		res["error"] = "Record Failed to Update. " + ret.Error.Error()
		return c.RenderJSON(res)
	}

	res["status"] = 200
	res["message"] = "Success"

	return c.RenderJSON(res)
}

func (c ApiAlternatif) Delete(id int) revel.Result {
	res := make(map[string]interface{})
	if err := c.Txn.First(&models.Alternatif{}, id).Error; err != nil {
		res["status"] = 404
		res["error"] = "Resource Not Found"
		return c.RenderJSON(res)
	}
	ret := c.Txn.Delete(&models.Alternatif{Id: id})
	
	if ret.Error != nil {
		res["status"] = 409
		res["error"] = "Record Failed to Delete. " + ret.Error.Error()
		return c.RenderJSON(res)
	}

	res["status"] = 200
	res["message"] = "Success"

	return c.RenderJSON(res)
}