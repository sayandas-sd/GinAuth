package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"_id"`
	FirstName     string    `json:"first_name" validate:"required,min=2,max=50"`
	LastName      string    `json:"last_name" validate:"required,min=2,max=50"`
	Email         string    `json:"email" validate:"email,required" gorm:"unique"`
	Password      string    `json:"password" validate:"required,min=4"`
	Phone         int       `json:"phone" validate:"required"`
	Token         string    `json:"token"`
	User_Type     string    `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_Token string    `json:"refersh_token"`
	Created_at    time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at    time.Time `json:"updated_at" gorm:"autoCreateTime"`
	User_id       string    `json:"user_id" gorm:"unique"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	user.User_id = uuid.NewString()
	return
}
