package model

import "time"

type User struct {
	ID        int         `gorm:"primaryKey;autoIncrement"`
	Username  string      `gorm:"type:VARCHAR(20);not null;" convert:"Username"`
	FirstName string      `gorm:"type:VARCHAR(70);not null;" convert:"FirstName"`
	LastName  string      `gorm:"type:VARCHAR(70)"`
	Profile   UserProfile `gorm:"foreignKey:ProfileID"`
	ProfileID int         // Foreign key
	Email     string      `gorm:"type:VARCHAR(70);not null;" convert:"Email"`
	Active    bool        `gorm:"type:bool;not null;"`
	Created   time.Time   `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated   time.Time   `gorm:"type:date;default:(CURRENT_DATE)"`
}

// DTO input
type AddUserIn struct {
	FirstName string `json:"firstname" binding:"required" convert:"FirstName"`
	LastName  string `json:"lastname"`
	Username  string `json:"username" binding:"required" convert:"Username"`
	Profile   int    `json:"profile" binding:"required" convert:"ProfileID"`
	Email     string `json:"email" binding:"required,email" convert:"Email"`
}

// DTO output
type GetUserOut struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Profile   string `json:"profile"`
	Email     string `json:"email"`
}
