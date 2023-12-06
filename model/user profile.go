package model

type UserProfile struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:VARCHAR(20);not null;" convert:"Profile.Name"`
}

// DTO input

type AddUserProfileIn struct {
	Name string `json:"Name" binding:"required"`
}

func (m *UserProfile) PopulateFromDTOInput(input AddUserProfileIn) {
	m.Name = input.Name
}
