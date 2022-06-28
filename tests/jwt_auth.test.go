package tests

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/revel/revel/testing"
)

type JwtAuthTest struct {
	testing.TestSuite
}
type JwtAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t *JwtAuthTest) Before() {
}

func (t *JwtAuthTest) TestLoginWithValidParams() {
	values := map[string]string{"user_name": "testuser", "password": "123456"}
	jsonValue, _ := json.Marshal(values)

	t.Post("/login", "application/json", bytes.NewBuffer(jsonValue))

	t.AssertOk()
	t.AssertContains("access_token")
	t.AssertContains("refresh_token")
}

func (t *JwtAuthTest) TestLoginWithInValidParams() {
	values := map[string]string{"user_name": "testuser", "password": "12345"}
	jsonValue, _ := json.Marshal(values)

	t.Post("/login", "application/json", bytes.NewBuffer(jsonValue))

	t.AssertStatus(401)
}

func (t *JwtAuthTest) TestLogoutWithValidToken() {
	values := map[string]string{"user_name": "testuser", "password": "123456"}
	jsonValue, _ := json.Marshal(values)

	t.Post("/login", "application/json", bytes.NewBuffer(jsonValue))

	//Read the response body
	var result JwtAuthResponse
	json.Unmarshal(t.ResponseBody, &result)

	req := t.PostCustom(t.BaseUrl()+"/logout", "application/json", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", result.AccessToken))

	req.MakeRequest()
	t.AssertOk()
}

func (t *JwtAuthTest) TestLogoutWithInValidToken() {
	req := t.PostCustom(t.BaseUrl()+"/logout", "application/json", nil)
	req.Header.Set("Authorization", "abcd")

	req.MakeRequest()
	t.AssertStatus(401)
}

func (t *JwtAuthTest) TestRefreshTokenWithValidToken() {
	values := map[string]string{"user_name": "testuser", "password": "123456"}
	jsonValue, _ := json.Marshal(values)
	t.Post("/login", "application/json", bytes.NewBuffer(jsonValue))

	//Read the response body
	var result JwtAuthResponse
	json.Unmarshal(t.ResponseBody, &result)

	values = map[string]string{"refresh_token": result.RefreshToken}
	jsonValue, _ = json.Marshal(values)
	req := t.PostCustom(t.BaseUrl()+"/refresh_tokens", "application/json", bytes.NewBuffer(jsonValue))

	req.MakeRequest()
	t.AssertOk()
}

func (t *JwtAuthTest) TestRefreshTokenWithInValidToken() {
	req := t.PostCustom(t.BaseUrl()+"/refresh_tokens", "application/json", nil)
	req.Header.Set("Authorization", "Bearer abcd")

	req.MakeRequest()
	t.AssertStatus(422)
}

func (t *JwtAuthTest) TestSignup() {
	values := map[string]string{
		"user_name":  "testuser",
		"password":   "123456",
		"email":      "test@gmail.com",
		"first_name": "john",
		"last_name":  "nguyen",
	}
	jsonValue, _ := json.Marshal(values)

	t.Post("/signup", "application/json", bytes.NewBuffer(jsonValue))

	t.AssertOk()
	t.AssertContains("user")
	t.AssertContains("email")
	t.AssertContains("first_name")
	t.AssertContains("last_name")
	t.AssertContains("user_name")
	t.AssertContains("access_token")
	t.AssertContains("refresh_token")
}

func (t *JwtAuthTest) After() {
}
