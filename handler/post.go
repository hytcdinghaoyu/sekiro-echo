package handler

import (
	"net/http"
	"strconv"

	"sekiro_echo/model"

	"github.com/gin-gonic/gin"
)

//CreatePost create user post
func CreatePost(c *gin.Context) {

	var p model.Post
	if err := c.Bind(&p); err != nil {
		return
	}

	p.UserId = 1
	p.PostSave()

	c.JSON(http.StatusCreated, p)
}

//FetchPost fetch user posts
func FetchPost(c *gin.Context) {
	var userID uint64
	userID = 1
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// Retrieve posts from database
	var Post model.Post
	posts := Post.GetUserPostsByUserId(userID, page, limit)

	c.JSON(http.StatusOK, map[string]interface{}{
		"title": "Post",
		"posts": posts,
	})

}
