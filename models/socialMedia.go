package models

type SocialMedia struct {
	GormModel
	Name 			string `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaUrl 	string `json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	UserID 			uint
	User			*User
}