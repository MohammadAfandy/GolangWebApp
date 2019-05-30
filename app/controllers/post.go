package controllers

import (
	"strings"
	"github.com/revel/revel"
	"myapp/app/models"
	"errors"
)

type Post struct {
	*revel.Controller
}

func (c Post) Index() revel.Result {
	posts := []models.Post{}

	result := DB.Order("id desc").Find(&posts)
	err := result.Error

	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}

	// return c.RenderJSON(posts)
	return c.Render(posts)
}

func (c Post) Create() revel.Result {
	post := models.Post {
		Body: strings.TrimSpace(c.Params.Form.Get("body")),
	}
	c.Validation.Required(post.Body).Message("Todo is Required")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Post.Index)
	}

	ret := DB.Create(&post)

	if ret.Error != nil {
		return c.RenderError(errors.New("Record Failed to Create. " + ret.Error.Error()))
	}

	return c.Redirect(Post.Index)
}

func (c Post) Delete(id uint64) revel.Result {
	ret := DB.Delete(&models.Post{Id: id})

	if ret.Error != nil {
		return c.RenderError(errors.New("Record Failed to Delete. " + ret.Error.Error()))
	}

	return c.Redirect(Post.Index)
	// return method
}