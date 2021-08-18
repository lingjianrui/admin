package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
)

/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
type Device struct {
	Base
	DeviceCode  string        `gorm:"size:255;not null" json:"deivce_code"`
	DeviceName  string        `json:"device_name"`
	ProjectName string        `json:"project_name"`
	DeviceType  string        `json:"device_type"`
	Description string        `gorm:"size:255" json:"description"`
	UserId      uint          `json:"user_id"`
	DeviceRoute []DeviceRoute `gorm:"FOREIGNKEY:DeviceID;ASSOCIATION_FOREIGNKEY:ID"`
}

func (device *Device) Insert(db *gorm.DB) (*Device, error) {
	var err error
	if e := db.Create(device).Error; e != nil {
		err = e
	}
	return device, err
}

func (device *Device) Delete(db *gorm.DB) (*Device, error) {
	var err error
	if e := db.Delete(device).Error; e != nil {
		err = e
	}
	return device, err
}

func (device *Device) FindDeviceById(db *gorm.DB) (*Device, error) {
	var err error
	var d Device
	if e := db.Debug().Model(device).Where("id = ?", device.ID).Find(&d).Error; e != nil {
		err = e
	}
	return &d, err
}

func (device *Device) FindDevices(db *gorm.DB) ([]*Device, error) {
	var err error
	var DeviceList []*Device
	if e := db.Debug().Model(device).Find(&DeviceList).Error; e != nil {
		err = e
	}
	return DeviceList, err
}

func (device *Device) ListDevices(db *gorm.DB, page int, pageSize int) ([]*Device, int, error) {
	var err error
	Devices := make([]*Device, 0)
	var total int = 0
	db.Model(&Device{}).Count(&total)
	if e := db.Debug().Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&Devices).Error; e != nil {
		err = e
	}
	return Devices, total, err
}
func (device *Device) ListEarthquakeDevices(db *gorm.DB) ([]*Device, int, error) {
	var err error
	Devices := make([]*Device, 0)
	var total int = 0
	db.Model(&Device{}).Where("device_type = ?", "earthquake").Count(&total)
	if e := db.Debug().Where("device_type = ?", "earthquake").Find(&Devices).Error; e != nil {
		err = e
	}
	return Devices, total, err
}
func (device *Device) FindDeviceAndLatestRouteList() (string, string, string) {
	var err error
	//var deviceRouteList []*DeviceRoute
	//var total int
	db, _ := sql.Open("taosSql", "root:taosdata@tcp(td1:6030)/test_db")
	fmt.Println("FindDeviceRoutesByDeviceCode")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	row, _ := db.Query("select latitude, longtitude, alitude from gnss_data where device_code='?' order by time desc limit 1", device.DeviceCode)
	if err != nil {
		panic(err)
	}
	var latitude string
	var longtitude string
	var alitude string

	if row.Next() {
		row.Scan(&latitude, &longtitude, alitude)
	}
	return latitude, longtitude, alitude
}
