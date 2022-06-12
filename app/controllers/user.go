package controllers

import (
	"Eatiplan-Auth/app/database"
	"Eatiplan-Auth/app/models"
	"Eatiplan-Auth/app/serializers"
	"net/http"

	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) CreateUser() revel.Result {
	var user models.User
	c.Params.BindJSON(&user)

	hashPassword, err := models.HashPassword(user.Password)
	user.Password = hashPassword

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity

		return c.RenderJSON(err.Error())
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.Response.Status = http.StatusUnprocessableEntity

		return c.RenderJSON(result.Error.Error())
	}

	c.Response.Status = http.StatusCreated
	return c.RenderJSON(serializers.UserSerializer(user))
}
