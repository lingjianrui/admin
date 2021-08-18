package api
import (
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
"fmt"
)

/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
type DeviceRouteResponse struct {
	Total int                  `json:"total"`
	Items []*model.DeviceRoute `json:"items"`
}

//获取商户列表
func (server *Server) GetDeviceRouteList(c *gin.Context) {
	var res map[string]interface{}
	page := c.Query("page")
	limit := c.Query("limit")
	pageN, _ := strconv.Atoi(page)
	limitN, _ := strconv.Atoi(limit)
	DeviceRoute := &model.DeviceRoute{}
	vlist, total, e := DeviceRoute.ListDeviceRoutes(server.DB, pageN, limitN)
	if e != nil {
		res = GenerateResponse(1, "Get Device List Error", e.Error(),nil)
	} else {
		res = GenerateResponse(20000, "", "Success",&DeviceRouteResponse{Total: total, Items: vlist})
	}
	c.JSON(http.StatusOK, res)
}


func (server *Server) GetRoutesByDeviceID(c *gin.Context) {
	var res map[string]interface{}
	page := c.Query("page")
	limit := c.Query("limit")
	deviceId := c.Query("deviceid")
	pageN, _ := strconv.Atoi(page)
	limitN, _ := strconv.Atoi(limit)
	DeviceRoute := &model.DeviceRoute{}
	vlist, total, e := DeviceRoute.FindDeviceRoutesByDeviceID(server.DB, deviceId, pageN, limitN)
	if e != nil {
		res = GenerateResponse(1, "Get Device List Error", e.Error(),nil)
	} else {
		res = GenerateResponse(20000, "", "Success",&DeviceRouteResponse{Total: total, Items: vlist})
	}
	c.JSON(http.StatusOK, res)
}

func (server *Server) GetRoutesByDeviceCode(c *gin.Context) {
	var res map[string]interface{}
	page := c.Query("page")
	limit := c.Query("limit")
	deviceId := c.Query("code")
	pageN, _ := strconv.Atoi(page)
	limitN, _ := strconv.Atoi(limit)
	DeviceRoute := &model.DeviceRoute{}
	fmt.Println("GetRoutesByDeviceCode")
	vlist, total, e := DeviceRoute.FindDeviceRoutesByDeviceCode(deviceId, pageN, limitN)
	if e != nil {
		res = GenerateResponse(1, "Get Device List Error", e.Error(),nil)
	} else {
		res = GenerateResponse(20000, "", "Success",&DeviceRouteResponse{Total: total, Items: vlist})
	}
	c.JSON(http.StatusOK, res)
}

//增加商户
func (server *Server) CreateDeviceRoute(c *gin.Context) {
	var res map[string]interface{}
	deviceId := c.PostForm("device_id")
	latitude := c.PostForm("latitude")
	longtitude := c.PostForm("longtitude")
	satelliteCount := c.PostForm("sitellite_count")
	altitude := c.PostForm("altitude")
	datetime := c.PostForm("datetime")
	Device := &model.DeviceRoute{DeviceID: deviceId,Latitude: latitude,Longtitude: longtitude,SatelliteCount: satelliteCount, Alitude: altitude,Datetime: datetime}
	v, e := Device.Insert(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Create Route Error", e.Error(),nil)
	} else {
		res = GenerateResponse(20000, "", "Success",v)
	}
	c.JSON(http.StatusOK, res)
}

//删除商户
func (server *Server) DeleteDeviceRoute(c *gin.Context) {
	var res map[string]interface{}
	id := c.Query("id")
	Device := &model.Device{}
	Device.ID = id
	v, e := Device.Delete(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Delete Device Error", e.Error(),nil)
	} else {
		res = GenerateResponse(20000, "", "Success" ,v)
	}
	c.JSON(http.StatusOK, res)
}
