package controllers

import (
	"net/http"

	"github.com/revel/revel"

	"Eatiplan-Auth/app/database"
	"Eatiplan-Auth/app/integrations"
	"Eatiplan-Auth/app/models"
	"Eatiplan-Auth/app/services"
)

type JwtAuth struct {
	*revel.Controller
}
type UserSt struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (c JwtAuth) Login() revel.Result {
	var userParams UserSt
	c.Params.BindJSON(&userParams)

	var current_user models.User
	database.DB.Where("user_name = ? ", userParams.UserName).First(&current_user)

	if !models.CheckPasswordHash(userParams.Password, current_user.Password) {
		c.Response.Status = http.StatusUnauthorized

		return c.RenderJSON("Your Username or Password is not correct")
	}

	token, err := services.CreateToken(uint64(current_user.ID))
	return renderToken(c, token, err)
}

func (c JwtAuth) Logout() revel.Result {
	deleted, delErr := services.DeleteAuthFromRequest(c.Request)

	if delErr != nil || deleted == 0 { //if any goes wrong
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("Can not process")
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON("Successfully logged out")
}

func (c JwtAuth) Refresh() revel.Result {
	mapToken := map[string]string{}
	if err := c.Params.BindJSON(&mapToken); err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(err.Error())
	}
	refreshToken := mapToken["refresh_token"]

	token, err := services.RefreshToken(refreshToken)
	return renderToken(c, token, err)
}

func renderToken(c JwtAuth, token *integrations.JwtToken, err error) revel.Result {
	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(err.Error())
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}
