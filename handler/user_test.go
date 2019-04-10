package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "sekiro_echo/conf"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"password":"123456","email":"jon@labstack.com"}`
)

func init() {
	// init conf
	if err := InitConfig("../conf/conf.toml"); err != nil {
		log.Panic(err)
	}

}

func TestSignup(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, Signup(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
