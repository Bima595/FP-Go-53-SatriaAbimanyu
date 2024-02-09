package models

import "gorm.io/gorm"

type Review struct {
    ID      uint `gorm:"primaryKey"`
    GameID  uint // ID dari tabel game
    UserID  uint // ID dari tabel user
    Rating  int
    Comment string
}

// Fungsi ini digunakan untuk membuat atau menginisialisasi tabel Reviews
func MigrateReviews(db *gorm.DB) {
    db.AutoMigrate(&Review{})
}
