package model

import "gorm.io/gorm"

type User struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    Username  string         `gorm:"unique;not null" json:"username"`  // 用户名
    Nickname  string         `gorm:"not null" json:"nickname"`         // 昵称
    Password  string         `gorm:"not null" json:"-"`                // 密码（不返回）
    Avatar    string         `json:"avatar"`                           // 头像 URL
    Signature string         `json:"signature"`                        // 个性签名
    CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"` // 创建时间
    UpdatedAt int64          `gorm:"autoUpdateTime" json:"updated_at"` // 更新时间
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                   // 软删除
}