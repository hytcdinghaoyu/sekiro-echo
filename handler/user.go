package handler

import (
	"net/http"
	"strconv"
	"time"

	"sekiro_echo/conf"
	"sekiro_echo/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Signup a user
func Signup(c *gin.Context) {
	// Bind
	User := model.User{}
	if err := c.Bind(&User); err != nil {
		return
	}

	// Validate
	if User.Email == "" || User.Password == "" {
		//return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
		c.JSON(http.StatusOK, gin.H{"code": LoginError, "message": StatusText(LoginError)})
	}

	User.CreateUser()

	c.JSON(http.StatusCreated, User)
}

//Login a user
func Login(c *gin.Context) {
	// Bind
	u := new(model.User)
	if err := c.Bind(&u); err != nil {
		return
	}

	// Find user
	user := u.GetUserByEmailPwd(u.Email, u.Password)
	if user == nil {
		//return &c.HTTPError{Code: http.StatusBadRequest, Message: "wrong email or password"}
		c.JSON(http.StatusOK, gin.H{"code": LoginError, "message": StatusText(LoginError)})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)

	claims["uid"] = strconv.FormatUint(user.UID, 10)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	var err error
	user.Token, err = token.SignedString([]byte(conf.Conf.Jwt.Secret))
	if err != nil {
		//return err
	}

	user.Password = "" // Don't send password
	c.JSON(http.StatusOK, user)
}
