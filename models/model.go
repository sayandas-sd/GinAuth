package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            *string   `gorm:"primaryKey;autoIncrement" json:"_id"`
	FirstName     *string   `json:"first_name" validate:"required,min=2,max=50"`
	LastName      *string   `json:"last_name" validate:"required,min=2,max=50"`
	Email         *string   `json:"email" validate:"email,required"`
	Password      *string   `json:"password" validate:"required,min=4"`
	Phone         *int      `json:"phone" validate:"required"`
	Token         *string   `json:"token"`
	User_Type     *string   `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_Token *string   `json:"refersh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       *string   `json:"user_id"`
}
