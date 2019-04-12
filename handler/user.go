package handler

import (
	"net/http"
	"strconv"
	"time"

	"sekiro_echo/conf"
	"sekiro_echo/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//Signup a user
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

//Login a user
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

	claims["uid"] = strconv.FormatUint(user.UID, 10)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	user.Token, err = token.SignedString([]byte(conf.Conf.Jwt.Secret))
	if err != nil {
		return err
	}

	user.Password = "" // Don't send password
	return c.JSON(http.StatusOK, user)
}

func userIDFromToken(c echo.Context) uint64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	uid, err := strconv.ParseUint(claims["uid"].(string), 10, 64)
	if err != nil {
		panic(err)
	}
	return uid
}
