package models

type LoginResponse struct {
	Token string `json:"token"`
}

type UserProfileResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}
