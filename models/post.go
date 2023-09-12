package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title string
    Slug string `gorm:"unique"`
    Body string

    UserID uint
}

type PostResource struct {
    ID uint `json:"id"`
    Title string `json:"name"`
    Slug string `json:"slug"`
    Body string `json:"body"`

    UserID uint `json:"userId"`

    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

func (ur* PostResource) ToString() (string, error) {
    r, err := json.Marshal(ur)
    return string(r), err
}

func (u* Post) Resource() PostResource {
    return PostResource {
        ID: u.ID,
        Title: u.Title,
        Slug: u.Slug,
        Body: u.Body,

        UserID: u.UserID,
        CreatedAt: u.CreatedAt,
        UpdatedAt: u.UpdatedAt,
    }
}
