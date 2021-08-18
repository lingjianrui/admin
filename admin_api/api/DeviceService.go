package api

import (
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
type DeviceResponse struct {
	Total int             `json:"total"`
	Items []*model.Device `json:"items"`
}

//获取设备列表
func (server *Server) GetDeviceList(c *gin.Context) {
	var res map[string]interface{}
	page := c.Query("page")
	limit := c.Query("limit")
	pageN, _ := strconv.Atoi(page)
	limitN, _ := strconv.Atoi(limit)
	Device := &model.Device{}
	vlist, total, e := Device.ListDevices(server.DB, pageN, limitN)
	if e != nil {
		res = GenerateResponse(1, "get device list error", e.Error(), nil)
	} else {
		res = GenerateResponse(20000, "", "success", &DeviceResponse{Total: total, Items: vlist})
	}
	c.JSON(http.StatusOK, res)
}

//获取地震设备列表
func (server *Server) GetEarthquakeDeviceList(c *gin.Context) {
	var res map[string]interface{}
	Device := &model.Device{}
	vlist, total, e := Device.ListEarthquakeDevices(server.DB)
	if e != nil {
		res = GenerateResponse(1, "get device list error", e.Error(), nil)
	} else {
		res = GenerateResponse(20000, "", "success", &DeviceResponse{Total: total, Items: vlist})
	}
	c.JSON(http.StatusOK, res)
}

//增加设备
func (server *Server) CreateDevice(c *gin.Context) {
	var res map[string]interface{}
	deviceCode := c.PostForm("device_code")
	deviceName := c.PostForm("device_name")
	deviceType := c.PostForm("device_type")
	description := c.PostForm("description")
	projectName := c.PostForm("project_name")
	Device := &model.Device{DeviceCode: deviceCode, DeviceName: deviceName, DeviceType: deviceType, Description: description, ProjectName: projectName}
	v, e := Device.Insert(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Create Device Error", e.Error(), nil)
	} else {
		res = GenerateResponse(20000, "", "Success", v)
	}
	c.JSON(http.StatusOK, res)
}

//删除设备
func (server *Server) DeleteDevice(c *gin.Context) {
	var res map[string]interface{}
	id := c.Query("id")
	device := &model.Device{}
	device.ID = id
	device, err := device.FindDeviceById(server.DB)
	if err != nil {
		res = GenerateResponse(1, "Can't Find The Device", err.Error(), nil)
		c.JSON(http.StatusOK, res)
		return
	}
	v, e := device.Delete(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Delete Device Error", e.Error(), nil)
	} else {
		res = GenerateResponse(20000, "", "Success", v)
	}
	c.JSON(http.StatusOK, res)
}

type HomeDeviceVo struct {
	DeviceCode     string `json:"device_code"`
	Datetime       string `json:"datetime"`
	Latitude       string `json:"latitude"`
	Longtitude     string `json:"longtitude"`
	DeviceType     string `json:"type"`
	SatelliteCount string `json:"satellite_count"`
	Alitude        string `json:"alitude"`
	ProjectName    string `json:"project_name"`
	Showflag       bool   `json:"showflag"`
}

//获取设备历史轨迹
func (server *Server) GetDeviceAndRouteList(c *gin.Context) {
	var res map[string]interface{}
	Device := &model.Device{}
	dList, e := Device.FindDevices(server.DB)
	ddlist := []*HomeDeviceVo{}
	for i := 0; i < len(dList); i++ {
		lat, lon, al := dList[i].FindDeviceAndLatestRouteList()
		vo := &HomeDeviceVo{}
		vo.DeviceCode = dList[i].DeviceCode
		vo.Longtitude = lon
		vo.Latitude = lat
		vo.Alitude = al
		vo.ProjectName = dList[i].ProjectName
		ddlist = append(ddlist, vo)
	}
	if e != nil {
		res = GenerateResponse(1, "get device list error", e.Error(), nil)
	} else {
		res = GenerateResponse(20000, "", "success", ddlist)
	}
	c.JSON(http.StatusOK, res)
}
