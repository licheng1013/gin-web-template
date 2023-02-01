package api

import (
	"gin-web-template/common"
	"gin-web-template/dto"
	"gin-web-template/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 通过包初始化注入
func init() {
	api := IndexApi{}
	group := common.R.Group("/index")
	group.GET("", api.index)          // http://localhost:8088/index
	group.GET("/err", api.err)        // http://localhost:8088/index/err
	group.POST("/upload", api.upload) //
}

type IndexApi struct {
	// 页 ->
	Page int64 `form:"page"`
	// 数量
	Size int64 `form:"size"`
}

// Page 演示
type Page struct {
	// 页 ->
	Page int64 `form:"page"`
	// 数量
	Size int64 `form:"size"`
}

// 接受参数返回示例 http://localhost:8088/index?page=1&size=10
func (t IndexApi) index(c *gin.Context) {
	_ = c.ShouldBindQuery(&t) //对于简单的参数绑定，异常可以不做处理，我们可以通过对绑定的对象值进行判断处理在决定信息的返回
	// 当然你也能定义另一个结构体进行绑定例如下面两行，然后上面一行就可以注释掉了
	//var page Page
	//_ = c.ShouldBindQuery(&page)
	log.Println("收到请求: ", t)
	c.JSON(http.StatusOK, dto.OkData("HelloWorld"))
}

// 模拟异常示例
func (t IndexApi) err(c *gin.Context) {
	panic(middleware.NewServiceError("测试异常"))
}

// 文件上传示例 -> image/upload_img.png -> 图片演示
func (t IndexApi) upload(c *gin.Context) {

	_ = c.ShouldBind(&t) //绑定表单参数示例,有时候上传文件会带点额外参数就需要这样做
	log.Println(t)

	// 单文件
	file, _ := c.FormFile("file")
	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	err := c.SaveUploadedFile(file, dst)
	if err != nil { // 这里一般可以不用处理，
		panic(err)
	}
	c.JSON(http.StatusOK, dto.OkData(file.Filename)) //我们把名字返回回去
}
