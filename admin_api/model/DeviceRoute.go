package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/taosdata/driver-go/taosSql"
	"database/sql"
	"fmt"
)
/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
type DeviceRoute struct {
	Base
	DeviceID       string  `gorm:"size:255;not null;INDEX" json:"deivce_id"`
	Latitude       string  `gorm:"size:255" json:"latitude"`
	Longtitude     string  `json:"longtitude"`
	SatelliteCount string  `json:"satellite_count"`
	Alitude        string  `json:"alitude"`
	Datetime       string  `json:"datetime"`
}

func (deviceRoute *DeviceRoute) Insert(db *gorm.DB) (*DeviceRoute, error) {
	var err error
	if e := db.Create(deviceRoute).Error; e != nil {
		err = e
	}
	return deviceRoute, err
}

func (deviceRoute *DeviceRoute) Delete(db *gorm.DB) (*DeviceRoute, error) {
	var err error
	if e := db.Delete(deviceRoute).Error; e != nil {
		err = e
	}
	return deviceRoute, err
}

func (deviceRoute *DeviceRoute) FindDeviceRoutes(db *gorm.DB) ([]*DeviceRoute, error) {
	var err error
	var deviceRouteList []*DeviceRoute
	if e := db.Debug().Model(deviceRoute).Find(&deviceRouteList).Error; e != nil {
		err = e
	}
	return deviceRouteList, err
}

func (deviceRoute *DeviceRoute) FindDeviceRoutesByDeviceID(db *gorm.DB,deviceid string, page int, pageSize int) ([]*DeviceRoute,int, error) {
	var err error
	var deviceRouteList []*DeviceRoute
	var total int
	db.Model(deviceRoute).Where("device_id = ?", deviceRoute.DeviceID).Count(&total)
	if e := db.Debug().Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Where("device_id = ?", deviceid).Find(&deviceRouteList).Error; e != nil {
		err = e
	}
	return deviceRouteList, total, err
}

func (deviceRoute *DeviceRoute) FindDeviceRoutesByDeviceCode(devicecode string, page int, pageSize int) ([]*DeviceRoute,int, error) {
	var err error
	var deviceRouteList []*DeviceRoute
	var total int
	db, err := sql.Open("taosSql", "root:taosdata@tcp(td1:6030)/test_db")
	fmt.Println("FindDeviceRoutesByDeviceCode")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	row, err := db.Query("select count(time) from gnss_data where device_code = '?'", devicecode)
	if err != nil {
		panic(err)
	}
	if row.Next() {
		row.Scan(&total)
	}
	fmt.Println(total)
	rows, _ := db.Query("select time, device_code, latitude, longtitude, alitude from gnss_data where device_code='?' order by time desc limit ? offset ?", devicecode,pageSize,page-1)
	for  rows.Next() {
		var t string
		var device_code string
		var latitude string
		var longtitude string
		var alitude string
		rows.Scan(&t, &device_code, &latitude,&longtitude, &alitude)
        route := &DeviceRoute{DeviceID:device_code,Latitude:latitude,Longtitude:longtitude,Alitude:alitude, Datetime:t}
		deviceRouteList = append(deviceRouteList, route)
	}
	return deviceRouteList, total, err
}

func (deviceRoute *DeviceRoute) ListDeviceRoutes(db *gorm.DB, page int, pageSize int) ([]*DeviceRoute, int, error) {
	var err error
	DeviceRoutes := make([]*DeviceRoute, 0)
	var total int = 0
	db.Model(&DeviceRoute{}).Count(&total)
	if e := db.Debug().Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&DeviceRoutes).Error; e != nil {
		err = e
	}
	return DeviceRoutes, total, err
}
