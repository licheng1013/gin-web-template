package api

import (
	"gin-web-template/common"
	"gin-web-template/dto"
	"gin-web-template/model"
	"gin-web-template/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func init() {
	c := &AdminApi{}
	c.init(common.R) //这里需要引入你的gin框架的实例
}

func (t AdminApi) init(g *gin.Engine) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("/admin")
	group.GET("/list", t.list) //不设置限制条件的画默认查询所有
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
}

// AdminApi 控制器
type AdminApi struct {
	Page int   `form:"page"`
	Size int   `form:"size"`
	Ids  []int `form:"ids"`
}

func (t AdminApi) db() *gorm.DB {
	return common.Db
}

// 分页列表
func (t AdminApi) list(c *gin.Context) {
	_ = c.Bind(&t)
	v := model.Admin{}
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, dto.OkData(service.AdminService.List(t.Page, t.Size, &v)))
}

// 根据主键Id查询记录
func (t AdminApi) one(c *gin.Context) {
	var v model.Admin
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, dto.OkData(service.AdminService.One(v.AdminId)))
}

// 修改记录
func (t AdminApi) update(c *gin.Context) {
	var v model.Admin
	_ = c.ShouldBindJSON(&v)
	service.AdminService.Update(v)
	c.JSON(http.StatusOK, dto.OkMsg("修改成功！"))
}

// 插入记录
func (t AdminApi) insert(c *gin.Context) {
	var v model.Admin
	_ = c.ShouldBindJSON(&v)
	service.AdminService.Insert(v)
	c.JSON(http.StatusOK, dto.OkMsg("插入成功！"))
}

// 根据主键删除
func (t AdminApi) delete(c *gin.Context) {
	_ = c.Bind(&t)
	service.AdminService.Delete(t.Ids)
	c.JSON(http.StatusOK, dto.OkMsg("删除成功！"))
}
