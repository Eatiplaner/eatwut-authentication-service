package controllers

import (
	"net/http"

	"github.com/revel/revel"

	"Eatiplan-Auth/app/integrations"
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

	// TODO: replace password with grpc valid user
	if userParams.Password != "123456" {
		c.Response.Status = http.StatusUnauthorized

		return c.RenderJSON("Your Username or Password is not correct")
	}

	// TODO: replace 1 with real user id
	token, err := services.CreateToken(uint64(1))
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
