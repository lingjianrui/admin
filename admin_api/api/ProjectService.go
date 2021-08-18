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
type VoResponse struct {
	Total int              `json:"total"`
	Items []*model.Project `json:"items"`
}

//获取所有记录
func (server *Server) GetProjectList(c *gin.Context) {
	var res map[string]interface{}
	page := c.Query("page")
	limit := c.Query("limit")
	pageN, _ := strconv.Atoi(page)
	limitN, _ := strconv.Atoi(limit)
	Project := &model.Project{}
	vlist, total, e := Project.ListItems(server.DB, pageN, limitN)
	if e != nil {
		res = GenerateResponse(1, "get device list error", e.Error(), nil)
	} else {
		res = GenerateResponse(20000, "", "success", &VoResponse{Total: total, Items: vlist})
	}
	c.JSON(http.StatusOK, res)
}

//增加项目
func (server *Server) CreateProject(c *gin.Context) {
	var res map[string]interface{}
	projectName := c.PostForm("project_name")
	projectDisplayName := c.PostForm("project_display_name")
	projectManagerName := c.PostForm("project_manager_name")
	description := c.PostForm("description")
	Project := &model.Project{ProjectName: projectName, ProjectDisplayName: projectDisplayName, ProjectManagerName: projectManagerName, Description: description}
	v, e := Project.Insert(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Create Object Error", e.Error(), nil)
	} else {
		res = GenerateResponse(20000, "", "Success", v)
	}
	c.JSON(http.StatusOK, res)
}

//删除项目
func (server *Server) DeleteProject(c *gin.Context) {
	var res map[string]interface{}
	id := c.Query("id")
	project := &model.Project{}
	project.ID = id
	project, err := project.FindById(server.DB)
	if err != nil {
		res = GenerateResponse(1, "Can't Find The Project", err.Error(), nil)
		c.JSON(http.StatusOK, res)
		return
	}
	v, e := project.Delete(server.DB)
	if e != nil {
		res = GenerateResponse(1, "Delete Object Error", e.Error(), nil)
	} else {
		res = GenerateResponse(20000, "", "Success", v)
	}
	c.JSON(http.StatusOK, res)
}
