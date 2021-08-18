package api

import (
	"backend/model"
	"backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */

//Token token对象
type Token struct {
	Token string `json:"token"`
}

//Login 登录请求
func (server *Server) Login(c *gin.Context) {
	var res  map[string]interface{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	if ! (len(username) > 0 && len(password) > 0) {
		res = GenerateResponse(1, "登录失败", "请输入用户名和密码",nil)
		c.JSON(http.StatusOK, res)
		return
	}
	u := &model.User{Name: username, Password: password}
	ul, e := u.FindUserByCredential(server.DB)
	if e != nil {
		res = GenerateResponse(1, "查找用户失败", e.Error(), nil)
		c.JSON(http.StatusOK, res)
		return
	}
	if len(ul) != 1 {
		res = GenerateResponse(1, "登录失败", "用户名不存在", nil)
		c.JSON(http.StatusOK, res)
		return
	}
	signedToken, err := util.GenerateToken(username, password)
	if err != nil {
		res = GenerateResponse(1, "generate token error", err.Error(), nil)
		c.JSON(http.StatusOK, res)
		return
	}
	token := &Token{Token: signedToken}
	res = GenerateResponse(20000, "", "success", token)
	c.JSON(http.StatusOK, res)
}

func (server *Server) Register(c *gin.Context) {
	var res map[string]interface{}
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	user := &model.User{Name: username, Password: pwd, Roles: "admin"}
	u, e := user.Insert(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Register Error", e.Error(),nil)
	} else {
		res = GenerateResponse(20000, "", "success",u)
	}
	c.JSON(http.StatusOK, res)
}
