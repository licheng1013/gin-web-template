package model

import (
    "time"
)

type Admin struct{
    // 管理员id
    AdminId int64 `json:"adminId" form:"adminId" gorm:"primaryKey" `
    // 账号
    UserName string `json:"userName" form:"userName" `
    // 密码
    Password string `json:"password" form:"password" `
    // 盐
    Salt string `json:"salt" form:"salt" `
    // 创建时间
    CreateTime time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime" `
    // 昵称
    NickName string `json:"nickName" form:"nickName" `
    
}
