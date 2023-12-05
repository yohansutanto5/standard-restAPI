package model

import "time"

type User struct {
	ID        int         `gorm:"primaryKey;autoIncrement"`
	Username  string      `gorm:"type:VARCHAR(20);not null;" convert:"Username"`
	FirstName string      `gorm:"type:VARCHAR(70);not null;" convert:"FirstName"`
	LastName  string      `gorm:"type:VARCHAR(70)" convert:"LastName"`
	Profile   UserProfile `gorm:"foreignKey:ProfileID" convert:"Profile"`
	ProfileID int         // Foreign key
	Email     string      `gorm:"type:VARCHAR(70);not null;" convert:"Email"`
	Active    bool        `gorm:"type:bool;not null;"`
	Created   time.Time   `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated   time.Time   `gorm:"type:date;default:(CURRENT_DATE)"`
}
type AddUserIn struct {
	FirstName string `json:"firstname" binding:"required" convert:"FirstName"`
	LastName  string `json:"lastname"`
	Username  string `json:"username" binding:"required" convert:"Username"`
	Profile   int    `json:"profile" binding:"required" convert:"ProfileID"`
	Email     string `json:"email" binding:"required,email" convert:"Email"`
}

type GetUserOut struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Profile   string `json:"profile"`
	Email     string `json:"email"`
}

func (m *User) PopulateFromDTOInput(input AddUserIn) {
	m.FirstName = input.FirstName
	m.LastName = input.LastName
	m.Username = input.Username
	m.ProfileID = input.Profile
	m.Email = input.Email
}
func (m *User) ConstructGetUserOut() (res GetUserOut) {
	res.FirstName = m.FirstName
	res.LastName = m.LastName
	res.Username = m.Username
	res.Email = m.Email
	res.Profile = m.Profile.Name
	return
}
