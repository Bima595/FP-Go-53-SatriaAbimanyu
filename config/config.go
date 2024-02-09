package config

import (
	"BACKEND/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    
    // dbURL := "user:password@tcp(hostname:port)/databasename"
    dsn := fmt.Sprintf("root:edDbcAA5DG1f-eDhhg5B3Dd42dCFHDFC@tcp(monorail.proxy.rlwy.net:10973)/finalproject")

    var err error
    DB, err = MySQL(dsn)
    if err != nil {
        log.Fatal(err)
    }
    models.MigrateRole(DB)
    models.MigrateGames(DB)
    models.MigrateReviews(DB)
    models.MigrateRatings(DB)
}

func MySQL(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }

    log.Println("Connected to the database")
    return db, nil
}
