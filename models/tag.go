package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Tag struct {
    gorm.Model
    Name string
    Slug string `gorm:"unique"`
}

type TagResource struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Slug string `json:"slug"`
}

func (ur* TagResource) ToString() (string, error) {
    r, err := json.Marshal(ur)
    return string(r), err
}

func (u* Tag) Resource() TagResource {
    return TagResource {
        ID: u.ID,
        Name: u.Name,
        Slug: u.Slug,
    }
}
