package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string
    Password string
    Email string `gorm:"unique"`
}

type UserResource struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

func (ur* UserResource) ToString() (string, error) {
    r, err := json.Marshal(ur)
    return string(r), err
}

func (u* User) Resource() UserResource {
    return UserResource {
        ID: u.ID,
        Name: u.Name,
        Email: u.Email,
        CreatedAt: u.CreatedAt,
        UpdatedAt: u.UpdatedAt,
    }
}
