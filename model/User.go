package model

import "time"

type User struct {
	ID        int         `gorm:"primaryKey;autoIncrement"`
	Username  string      `gorm:"type:VARCHAR(20);not null;unique"`
	FirstName string      `gorm:"type:VARCHAR(70);not null;"`
	LastName  string      `gorm:"type:VARCHAR(70);"`
	ProfileID int         `gorm:"not null;"`
	Profile   UserProfile `gorm:"foreignKey:ProfileID"`
	Email     string      `gorm:"type:VARCHAR(70);not null;"`
	Active    bool        `gorm:"type:bool;not null;"`
	Created   time.Time   `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated   time.Time   `gorm:"type:date;default:(CURRENT_DATE)"`
}

// DTO input and func to populate it
type UserInput struct {
	Username  string    `json:"Username" binding:"required"`
	FirstName string    `json:"FirstName" binding:"required"`
	LastName  string    `json:"LastName"`
	ProfileID int       `json:"ProfileID"`
	Email     string    `json:"Email" binding:"email"`
	Updated   time.Time `json:"Updated"`
}

func (m *User) PopulateFromDTOInput(input UserInput) {
	m.Username = input.Username
	m.FirstName = input.FirstName
	m.LastName = input.LastName
	m.ProfileID = input.ProfileID
	m.Email = input.Email
	m.Updated = input.Updated

}

// DTO out and func to populate it
type UserOutput struct {
	ID        int       `json:"ID"`
	Username  string    `json:"Username"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	ProfileID int       `json:"ProfileID"`
	Email     string    `json:"Email"`
	Active    bool      `json:"Active"`
	Created   time.Time `json:"Created"`
}

func (m *User) PopulateDTOOutput() (output UserOutput) {
	output.ID = m.ID
	output.Username = m.Username
	output.FirstName = m.FirstName
	output.LastName = m.LastName
	output.ProfileID = m.ProfileID
	output.Email = m.Email
	output.Active = m.Active
	output.Created = m.Created

	return
}
