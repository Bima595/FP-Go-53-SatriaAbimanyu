package models

import "time"

type Game struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"not null"`
    Description string    `gorm:"type:text"`
    ReleaseDate time.Time 
}