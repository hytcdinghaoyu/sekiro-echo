package handler

import (
	"net/http"
	"strconv"
	"time"

	"sekiro_echo/model"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

//FetchMatches fetch matches
func FetchMatches(c echo.Context) (err error) {
	date := c.QueryParam("date")
	status := c.QueryParam("status")

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}
	if status == "" {
		status = "SCHEDULED"
	}
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	// Retrieve matches from database
	matches := []*model.Match{}
	db := mongodb.Clone()
	if err = db.DB("football_data").C("matches").
		Find(bson.M{"Status": status, "matchdate": date}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&matches); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, matches)
}
