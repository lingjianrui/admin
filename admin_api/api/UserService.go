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
type UserResponse struct {
	Total int           `json:"total"`
	Items []*model.User `json:"items"`
}

type UserVo struct {
	Roles         []string `json:"roles"`
	Instroduction string   `json:"introduction"`
	Avatar        string   `json:"avatar"`
	Name          string   `json:"name"`
}

//获取用户详情
func (server *Server) GetUserInfo(c *gin.Context) {
	var res map[string]interface{}
	d := &UserVo{
		Roles:         []string{"admin"},
		Instroduction: "test",
		Avatar:        "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:          "xiaohei",
	}
	res = GenerateResponse(20000, "", "success",d)
	c.JSON(http.StatusOK, res)
}

//获取用户列表
func (server *Server) GetUserList(c *gin.Context) {
	var res map[string]interface{}
	page := c.Query("page")
	limit := c.Query("limit")
	page_n, _ := strconv.Atoi(page)
	limit_n, _ := strconv.Atoi(limit)
	user := &model.User{}
	vlist, total, e := user.ListUsers(server.DB, page_n, limit_n)
	if e != nil {
		res = GenerateResponse(1, "Get User List Error", e.Error(),nil)
	} else {
		res = GenerateResponse(20000, "Get User List Error", "success",&UserResponse{Total: total, Items: vlist})
	}
	c.JSON(http.StatusOK, res)
}

//增加用户
func (server *Server) CreateUser(c *gin.Context) {
	var res map[string]interface{}
	username := c.PostForm("name")
	pwd := c.PostForm("password")
	roles := c.PostForm("roles")
	user := &model.User{Name: username, Password: pwd, Roles: roles}
	u, e := user.Insert(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Create User Error", e.Error(),nil)
	} else {
		res = GenerateResponse(1, "", "success",u)
	}
	c.JSON(http.StatusOK, res)
}

//删除用户
func (server *Server) DeleteUser(c *gin.Context) {
	var res map[string]interface{}
	id := c.Param("id")
	user := &model.User{}
	user.ID = id
	user, err := user.FindUserById(server.DB)
	if err != nil {
		res = GenerateResponse(1, "Not Able To Find The Device", err.Error(),nil)
		c.JSON(http.StatusOK, res)
		return
	}
	v, e := user.Delete(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Delete User Error", e.Error(),nil)
	} else {
		res = GenerateResponse(20000, "", "success",v)
	}
	c.JSON(http.StatusOK, res)
}
