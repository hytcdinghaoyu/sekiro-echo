package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"sekiro_echo/model"
	. "sekiro_echo/conf"
	"fmt"
	"reflect"
)

func Signup(c echo.Context) (err error) {
	// Bind
	User := model.User{}
	if err = c.Bind(&User); err != nil {
		return
	}

	// Validate
	if User.Email == "" || User.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	User.CreateUser()

	return c.JSON(http.StatusCreated, User)
}

func Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(&u); err != nil {
		return
	}

	// Find user
	user := u.GetUserByEmailPwd(u.Email, u.Password)
	if user == nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "wrong email or password"}
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	
	claims["uid"] = user.UID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	fmt.Println(reflect.TypeOf(claims["exp"]))
	fmt.Println(reflect.TypeOf(claims["uid"]))

	// Generate encoded token and send it as response
	user.Token, err = token.SignedString([]byte(Conf.Jwt.Secret))
	if err != nil {
		return err
	}

	user.Password = "" // Don't send password
	return c.JSON(http.StatusOK, user)
}


func userIDFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	fmt.Println(user)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(reflect.TypeOf(claims["uid"]))
	fmt.Println(reflect.TypeOf(claims["exp"]))
	return claims["uid"].(uint)
}