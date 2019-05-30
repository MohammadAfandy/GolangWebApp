package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	welcome := "Selamat Datang !"
	return c.Render(welcome)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Name is Required")
	c.Validation.MinSize(myName, 5).Message("Name \"" + myName + "\" is not Long Enough")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}