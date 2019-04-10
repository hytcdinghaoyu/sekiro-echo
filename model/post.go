package model

import (
	"time"

	"github.com/hb-go/echo-web/module/log"
)

func (p *Post) GetPostById(id uint64) *Post {
	post := Post{}
	if err := DB().Where("id = ?", id).First(&post).Error; err != nil {
		log.Debugf("Get post error: %v", err)
		return nil
	}

	if err := DB().Model(&post).Related(&post.User).Error; err != nil {
		log.Debugf("Post user related error: %v", err)
		return &post
	}

	return &post
}

func (p *Post) GetUserPostsByUserId(userId uint64, page int, size int) *[]Post {
	posts := []Post{}
	if err := DB().Offset((page - 1) * size).Limit(size).Find(&posts).Error; err != nil {
		log.Debugf("Get user posts error: %v", err)
		return nil
	}

	for key, post := range posts {
		if err := DB().Model(&post).Related(&post.User).Error; err != nil {
			log.Debugf("Post user related error: %v", err)
		}
		posts[key] = post
	}

	return &posts
}

func (p *Post) PostSave() {
	if err := DB().Create(p).Error; err != nil {
		panic(err)
	}
}

type Post struct {
	Id        uint      `json:"id,omitempty"`
	UserId    uint64    `json:"user_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Context   string    `json:"context,omitempty"`
	CreatedAt time.Time `gorm:"column:created_time" json:"created_time,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_time" json:"updated_time,omitempty"`

	User User `gorm:"ForeignKey:UserId;AssociationForeignKey:UID" json:"user"`
}

func (p Post) TableName() string {
	return "post"
}
