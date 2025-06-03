package router

import (
    "github.com/gin-gonic/gin"
    "goim/controller"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    userController := controller.UserController{DB: db}

    r.POST("/register", userController.Register)
    r.POST("/login", userController.Login)

    return r
}