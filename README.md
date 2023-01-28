# GinWebTemplate
## 介绍
- 一个简单Gin框架Web模板

## 功能
- 跨域中间件
- 异常中间件
- 路由注册
- Gorm插件模板和自定义示例
- Mysql示例

## 启动项目
- main.go
- 需要创建数据库t_gorm并导入sql目录的t_gorm.sql文件
- 然后修改common目录下mysql.go文件

## 其他
- gorm目录下为插件自定义模板示例

## 启动后访问
- http://localhost:8088/admin/list?size=10
- http://localhost:8088/admin/one?adminId=3
