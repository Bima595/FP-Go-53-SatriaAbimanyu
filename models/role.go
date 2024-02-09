package models

import "gorm.io/gorm"

type Role struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"unique;not null"`
}

// Fungsi ini digunakan untuk membuat atau menginisialisasi tabel Roles
func MigrateRole(db *gorm.DB) {
    db.AutoMigrate(&Role{})
}