package model

import (
	"github.com/jinzhu/gorm"
)

/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
type Project struct {
	Base
	ProjectName        string `json:"project_name"`
	ProjectDisplayName string `json:"project_display_name"`
	ProjectManagerName string `json:"project_manager_name"`
	ProjectType        string `json:"project_type"`
	Description        string `gorm:"size:255" json:"description"`
}

func (object *Project) Insert(db *gorm.DB) (*Project, error) {
	var err error
	if e := db.Create(object).Error; e != nil {
		err = e
	}
	return object, err
}

func (object *Project) Delete(db *gorm.DB) (*Project, error) {
	var err error
	if e := db.Delete(object).Error; e != nil {
		err = e
	}
	return object, err
}

func (object *Project) FindById(db *gorm.DB) (*Project, error) {
	var err error
	var d Project
	if e := db.Debug().Model(object).Where("id = ?", object.ID).Find(&d).Error; e != nil {
		err = e
	}
	return &d, err
}

func (object *Project) ListItems(db *gorm.DB, page int, pageSize int) ([]*Project, int, error) {
	var err error
	Projects := make([]*Project, 0)
	var total int = 0
	db.Model(&Project{}).Count(&total)
	if e := db.Debug().Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&Projects).Error; e != nil {
		err = e
	}
	return Projects, total, err
}
