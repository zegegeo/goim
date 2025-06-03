package main

import (
    "goim/model"
    "goim/router"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, _ := gorm.Open(sqlite.Open("goim.db"), &gorm.Config{})
    _ = db.AutoMigrate(&model.User{})

    r := router.SetupRouter(db)
    r.Run(":8080")
}