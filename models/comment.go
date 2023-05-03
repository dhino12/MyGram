package models

type Comment struct {
	GormModel
	UserID 		uint
	PhotoID		uint
	Message 	string `json:"message" form:"message" valid:"required~Your message is required"`
	User 		*User
}