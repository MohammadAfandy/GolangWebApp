package controllers

import (
	"net/http"
	"strconv"
	"github.com/revel/revel"
)

type Person struct {
	*revel.Controller
}

type PersonResource struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Job string `json:"job"`
}

func (c Person) Index() revel.Result {
	welcome := "Welcome !"
	return c.Render(welcome)
}

func (c Person) Show() revel.Result {
	var person PersonResource

	id := c.Params.Route.Get("id")
	
	person.ID, _ = strconv.Atoi(id)
	person.Name = "Afandy"
	person.Job = "Programmer"

	c.Response.Status = http.StatusOK

	return c.RenderJSON(person)
}