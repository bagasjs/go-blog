package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    UserID uint
    PostID uint
    Body string
}

type CommentResource struct {
    ID uint `json:"id"`
    UserID uint `json:"userId"`
    PostID uint `json:"postId"`
    Body string `json:"body"`
}

func (ur* CommentResource) ToString() (string, error) {
    r, err := json.Marshal(ur)
    return string(r), err
}

func (u* Comment) Resource() CommentResource {
    return CommentResource {
        ID: u.ID,
        UserID: u.UserID,
        PostID: u.PostID,
        Body: u.Body,
    }
}
