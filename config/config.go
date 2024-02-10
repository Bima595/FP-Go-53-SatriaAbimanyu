package config

import (
 "BACKEND/models"
 "BACKEND/utils"
 "fmt"

 "gorm.io/driver/mysql"
 "gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
  username := utils.Getenv("DATABASE_USERNAME", "root")
  password := utils.Getenv("DATABASE_PASSWORD", "Bima123yayang")
  host := utils.Getenv("DATABASE_HOST", "localhost")
  port := utils.Getenv("DATABASE_PORT", "3306")
  database := utils.Getenv("DATABASE_NAME", "finalproject")

  dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic(err.Error())
  }

  db.AutoMigrate(&models.User{}, &models.Game{}, &models.GameType{}, &models.Review{}, &models.Role{})

  return db
}