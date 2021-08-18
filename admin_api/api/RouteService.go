package api

import (
	"backend/middleware"
	"net/http"
)

/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
func (server *Server) initializeRoutes() {
	server.Router.Use(middleware.Cors())
	server.Router.Use(middleware.LoggerToFile())

	server.Router.GET("/ping", server.Ping)
	server.Router.POST("/login", server.Login)
	server.Router.POST("/register", server.Register)

	v1 := server.Router.Group("/api/v1")
	v1.Use(middleware.JWT())
	{
		v1.POST("/post", server.ImagePost)
		v1.StaticFS("/upload", http.Dir("upload"))
		v1.GET("/home", server.GetDeviceAndRouteList)
		v1.GET("/userinfo", server.GetUserInfo)
		v1.GET("/user", server.GetUserList)
		v1.POST("/user", server.CreateUser)
		v1.DELETE("/user", server.DeleteUser)

		v1.GET("/project", server.GetProjectList)
		v1.POST("/project", server.CreateProject)
		v1.DELETE("/project", server.DeleteProject)

		v1.GET("/device", server.GetDeviceList)
		v1.GET("/earthquake/device", server.GetEarthquakeDeviceList)
		v1.POST("/device", server.CreateDevice)
		v1.DELETE("/device", server.DeleteDevice)

		v1.GET("/routelist", server.GetRoutesByDeviceID)
		v1.GET("/route", server.GetDeviceRouteList)
		v1.GET("/routehistory", server.GetRoutesByDeviceCode)
		v1.POST("/map/convert", server.BaiduLatLonConvert)
		v1.POST("/route", server.CreateDeviceRoute)
		v1.DELETE("/route", server.DeleteDeviceRoute)
	}
}
