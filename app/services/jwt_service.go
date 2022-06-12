package services

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"Eatiplan-Auth/app/externals"

	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"github.com/twinj/uuid"
)

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func CreateToken(userid uint64) (*externals.JwtToken, error) {
	td := &externals.JwtToken{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	err = createAuth(userid, td)
	if err != nil {
		return nil, err
	}

	return td, nil
}

func RefreshToken(refreshToken string) (*externals.JwtToken, error) {
	//Since token is valid, get the uuid:
	token, err := verifyRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims

	if !ok || !token.Valid {
		return nil, errors.New("Refresh Token expired")
	}

	refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
	if !ok {
		return nil, errors.New("Error occurred")
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}

	//Delete the previous Refresh Token
	deleted, delErr := deleteAuth(refreshUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		return nil, errors.New("Refresh Token deleted")
	}

	//Create new pairs of refresh and access tokens
	ts, createErr := CreateToken(userId)
	if createErr != nil {
		return nil, createErr
	}

	return ts, nil
}

func TokenValid(r *revel.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractUserIdFromRequest(r *revel.Request) (uint64, error) {
	tokenAuth, err := extractTokenMetadata(r)
	if err != nil {
		return 0, err
	}

	userId, err := fetchAuth(tokenAuth)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func DeleteAuthFromRequest(r *revel.Request) (int64, error) {
	au, err := extractTokenMetadata(r)

	if err != nil {
		return 0, nil
	}

	return deleteAuth(au.AccessUuid)
}

// Private
func verifyToken(r *revel.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func verifyRefreshToken(refreshToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		return nil, errors.New("Refresh Token expired")
	}

	return token, nil
}

func extractToken(r *revel.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func extractTokenMetadata(r *revel.Request) (*AccessDetails, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func fetchAuth(authD *AccessDetails) (uint64, error) {
	userid, err := externals.RedisClient.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func createAuth(userid uint64, td *externals.JwtToken) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := externals.RedisClient.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := externals.RedisClient.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func deleteAuth(givenUuid string) (int64, error) {
	deleted, err := externals.RedisClient.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
