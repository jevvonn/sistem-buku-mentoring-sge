package entity

import "github.com/google/uuid"

type Book struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Author      string    `gorm:"type:varchar(255);not null"`
	Publisher   string    `gorm:"type:varchar(255);not null"`
	PublishYear int       `gorm:"type:int;not null"`
}
