package config

import (
    "fmt"
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

const (
    username = "root"
    password = "Bima123yayang"
    database = "finalproject"
)

var dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = MySQL()
    if err != nil {
        log.Fatal(err)
    }
}

func MySQL() (*gorm.DB, error) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }

    log.Println("Connected to the database")
    return db, nil
}
