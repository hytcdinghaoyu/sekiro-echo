package model

import (
	"time"
	"github.com/hb-go/echo-web/module/log"
	"github.com/hb-go/gorm"
	"github.com/satori/go.uuid"
)

func (u *User) CreateUser() {
	if err := DB().Create(u).Error; err != nil {
		log.Debugf("Create user error: %v", err)
		panic(err)
	}
}

func (u *User) GetUserByEmailPwd(email string, pwd string) *User {
	user := User{}
	if err := DB().Where("email = ? AND password = ?", email, pwd).First(&user).Error; err != nil {
		log.Debugf("GetUserByEmailPwd error: %v", err)
		return nil
	}
	return &user
}

type (
	User struct {
		UID       uint 			`json:"uid" gorm:"primary_key"`
		UUID      string 		`json:"uuid"`
		Email     string        `json:"email"`
		Password  string        `json:"password,omitempty"`
		Token     string        `json:"token,omitempty" gorm:"-"`
		CreatedAt time.Time 	`gorm:"column:created_time" json:"created_time,omitempty"`
		UpdatedAt time.Time 	`gorm:"column:updated_time" json:"updated_time,omitempty"`
	}
)

func (user *User) BeforeCreate(scope *gorm.Scope) error {
  scope.SetColumn("UUID", uuid.NewV4())
  return nil
}

func (u User) TableName() string {
	return "user"
}


