package models

type Review struct {
    ID       uint   `gorm:"primaryKey"`
    GameID   uint   `gorm:"not null"`
    UserID   uint   `gorm:"not null"`
    Rating   int    `gorm:"not null"`
    Comment  string `gorm:"type:text"`
}