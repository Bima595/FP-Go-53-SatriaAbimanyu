package models

import "gorm.io/gorm"

type GameType struct {
    ID       uint   `gorm:"primaryKey"`
    GameID   uint   `gorm:"not null"`
    Theme    string `gorm:"type:varchar(255);not null"`
}

// Fungsi ini digunakan untuk membuat atau menginisialisasi tabel GameTypes
func MigrateGameTypes(db *gorm.DB) {
    db.AutoMigrate(&GameType{})
}