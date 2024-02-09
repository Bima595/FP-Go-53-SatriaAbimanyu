package models

import (
    "gorm.io/gorm"
)

type Game struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"not null"`
    Description string    
}


// Fungsi ini digunakan untuk membuat atau menginisialisasi tabel Games
func MigrateGames(db *gorm.DB) {
    db.AutoMigrate(&Game{})
}
