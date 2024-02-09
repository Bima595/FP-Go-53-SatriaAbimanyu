package models

import "gorm.io/gorm"

type Rating struct {
    ID        uint   `gorm:"primaryKey"`
    ReviewID  uint   `gorm:"not null"`
    Rating    int    `gorm:"not null"`
    UserID    uint   `gorm:"not null"`
}

// Fungsi ini digunakan untuk membuat atau menginisialisasi tabel Ratings
func MigrateRatings(db *gorm.DB) {
    db.AutoMigrate(&Rating{})
}
