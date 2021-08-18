package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)
/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
type Base struct {
	ID        string     `gorm:"type:char(36);primary_key;INDEX"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	u2 := uuid.NewV4().String()
	return scope.SetColumn("ID", u2)
}
