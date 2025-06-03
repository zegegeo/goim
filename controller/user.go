package controller

import (
    "github.com/gin-gonic/gin"
    "goim/model"
    "goim/utils"
    "gorm.io/gorm"
    "net/http"
)

type UserController struct {
    DB *gorm.DB
}

// 用户注册
func (u *UserController) Register(c *gin.Context) {
    var req model.User
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
        return
    }

    // hash 密码
    req.Password = utils.HashPassword(req.Password)

    if err := u.DB.Create(&req).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// 用户登录
func (u *UserController) Login(c *gin.Context) {
    var req model.User
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
        return
    }

    var user model.User
    if err := u.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
        return
    }

    if !utils.CheckPassword(req.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "登录成功",
        "token":   token,
        "user_id": user.ID,
    })
}