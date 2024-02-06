package models

type Rating struct {
    ID       uint `gorm:"primaryKey"`
    ReviewID uint `gorm:"not null"`
    UserID   uint `gorm:"not null"`
    Rating   int  `gorm:"not null"`
}