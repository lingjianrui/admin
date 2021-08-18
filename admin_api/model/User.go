package model

import (
	"github.com/jinzhu/gorm"
)
/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
type User struct {
	Base
	Name      string    `gorm:"size:255;not null;unique" json:"user_name"`
	Password  string    `gorm:"size:255" json:"user_password"`
	Devices   []Device  `json:"devices"`
	Roles     string    `json:"roles"`
}

func (user *User) Insert(db *gorm.DB) (*User, error) {
	var err error
	if e := db.Create(user).Error; e != nil {
		err = e
	}
	return user, err
}

func (user *User) Delete(db *gorm.DB) (*User, error) {
	var err error
	if e := db.Delete(user).Error; e != nil {
		err = e
	}
	return user, err
}

func (user *User) FindUsers(db *gorm.DB) ([]*User, error) {
	var err error
	var userList []*User
	if e := db.Debug().Model(user).Find(&userList).Error; e != nil {
		err = e
	}
	return userList, err
}
func (user *User) FindUserById(db *gorm.DB) (*User, error) {
	var err error
	var u User
	if e := db.Debug().Model(user).Where("id = ?",user.ID).Find(&u).Error; e != nil {
		err = e
	}
	return &u, err
}
func (user *User) FindUserByCredential(db *gorm.DB) ([]*User, error) {
	var err error
	var userList []*User
	if e := db.Debug().Where(&User{Name: user.Name, Password: user.Password}).First(&userList).Error; e != nil {
		err = e
	}
	return userList, err
}
func (user *User) ListUsers(db *gorm.DB, page int, pageSize int) ([]*User, int, error) {
	var err error
	users := make([]*User, 0)
	var total int = 0
	db.Model(&User{}).Count(&total)
	if e := db.Debug().Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&users).Error; e != nil {
		err = e
	}
	return users, total, err
}
