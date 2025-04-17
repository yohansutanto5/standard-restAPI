package model

type UserProfile struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:VARCHAR(20);not null;"`

}

// DTO input and func to populate it
type UserProfileInput struct {
	Name string `json:"Name" binding:"required"`

}
func (m *UserProfile) PopulateFromDTOInput(input UserProfileInput) {
	m.Name = input.Name

}

// DTO out and func to populate it
type UserProfileOutput struct {
	ID int `json:"ID"`
	Name string `json:"Name"`

}

func (m *UserProfile) PopulateDTOOutput() (output UserProfileOutput){
	output.ID = m.ID
	output.Name = m.Name

  return
}