package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"sekiro_echo/model"
)

func CreatePost(c echo.Context) (err error) {

	var p model.Post
	if err = c.Bind(&p); err != nil {
		return
	}

	p.UserId = userIDFromToken(c)
	p.PostSave()

	return c.JSON(http.StatusCreated, p)
}

func FetchPost(c echo.Context) (err error) {
	userId := userIDFromToken(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// Retrieve posts from database
	var Post model.Post
	posts := Post.GetUserPostsByUserId(userId, page, limit)

	c.JSON(http.StatusOK, map[string]interface{}{
		"title": "Post",
		"posts": posts,
	})

	return nil
}