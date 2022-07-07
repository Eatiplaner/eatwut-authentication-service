package controllers

import (
	"net/http"

	"github.com/revel/revel"

	"Eatiplan-Auth/app/grpc/client"
	"Eatiplan-Auth/app/grpc/rpc_pb"
	"Eatiplan-Auth/app/integrations"
	"Eatiplan-Auth/app/services"
)

type JwtAuth struct {
	*revel.Controller
}

func (c JwtAuth) Login() revel.Result {
	var userParams rpc_pb.FindUserRequest

	c.Params.BindJSON(&userParams)

	user, err := client.Service.FindUserByCredential(&userParams)

	if err != nil {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("Please provide correct credential")
	}

	token, err := services.CreateToken(uint64(user.Id))
	return renderUserAndToken(c, user, token, err)
}

func (c JwtAuth) Signup() revel.Result {
	var userParams rpc_pb.CreateRequest
	c.Params.BindJSON(&userParams)

	user, err := client.Service.CreateUser(&userParams)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(err.Error())
	}

	token, err := services.CreateToken(uint64(1))

	return renderUserAndToken(c, user, token, err)
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

func renderUserAndToken(
	c JwtAuth,
	user *rpc_pb.UserResponse,
	token *integrations.JwtToken,
	err error) revel.Result {
	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(err.Error())
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(map[string]interface{}{
		"user":          user,
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
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
