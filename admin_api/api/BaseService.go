package api

import (
	"backend/middleware"
	"backend/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
type Context struct {
	*gin.Context
}

//Server 服务模型
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

type FileVo struct {
	Fl string `json:"fileuid"`
}

func GenerateResponse(code int, e string, msg string, data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["error"] = e
	res["message"] = msg
	res["data"] = data
	return res
}

//Ping 连通性测试
func (server *Server) Ping(c *gin.Context) {
	var res map[string]interface{}
	middleware.Logger().Info("记录一下日志", "Info")
	res = GenerateResponse(20000, "", "success", "pong")
	c.JSON(http.StatusOK, res)
}

type Latlon struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type BaiduResponse struct {
	Status int      `json:"status"`
	Result []Latlon `json:"result"`
}

//百度地图地址转换
func (server *Server) BaiduLatLonConvert(c *gin.Context) {
	var res map[string]interface{}
	jData := make(map[string]interface{})
	c.BindJSON(&jData)
	dmaplist := jData["data"]
	fmt.Println(dmaplist)
	param := ""
	for i := 0; i < len(dmaplist.([]interface{})); i++ {
		p := dmaplist.([]interface{})[i].(map[string]interface{})["lon"].(string) + "," + dmaplist.([]interface{})[i].(map[string]interface{})["lat"].(string) + ";"
		param = param + p
	}
	req := fmt.Sprintf("http://api.map.baidu.com/geoconv/v1/?coords=%s&from=1&to=5&ak=bGbhDOp2aYUm49GTXG9jO51o195dudwT", param[0:len(param)-1])
	//req := fmt.Sprintf("http://api.map.baidu.com/geoconv/v1/?coords=116.809640,39.944340&from=1&to=5&ak=bGbhDOp2aYUm49GTXG9jO51o195dudwT")
	fmt.Println(req)
	resp, err := http.Get(req)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var baiduRes BaiduResponse
	fmt.Println(string(body))
	err = json.Unmarshal([]byte(string(body)), &baiduRes)
	if err != nil {
	}
	middleware.Logger().Info("记录一下日志", "Info")
	res = GenerateResponse(20000, "", "success", baiduRes.Result)
	c.JSON(http.StatusOK, res)
}

func (server *Server) ImagePost(c *gin.Context) {
	var res map[string]interface{}
	header, err := c.FormFile("file")
	if err != nil {
		res = GenerateResponse(1, "找不到文件", err.Error(), nil)
	}
	uid := uuid.NewV4()
	dst := "upload/" + uid.String()
	if err := c.SaveUploadedFile(header, dst); err != nil {
		res = GenerateResponse(1, "保存文件失败", err.Error(), nil)
	}
	fileVo := &FileVo{Fl: uid.String()}
	res = GenerateResponse(20000, "", "success", fileVo)
	c.JSON(http.StatusOK, res)
}

func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}

//Initialize 初始化数据库
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		middleware.Logger().Infof("数据库连接:%s", DBURL)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		server.DB.SingularTable(true)
		if err != nil {
			middleware.Logger().Errorf("连接数据库错误:%s", err)
		} else {
			middleware.Logger().Errorf("连接数据库成功:%s", Dbdriver)
		}
	} else {
		middleware.Logger().Error("Unknown Driver")
	}
	//数据库初始化修改
	server.DB.Debug().AutoMigrate(
		&model.Project{},
		&model.Device{},
		&model.DeviceRoute{},
		&model.User{},
	)
	server.Router = gin.Default()
	server.initializeRoutes()
}

//Run 系统运行入口方法
func (server *Server) Run(addr string) {
	middleware.Logger().Infof("service is runing")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
