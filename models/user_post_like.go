package models

import "gorm.io/gorm"

type UserLikePost struct {
    gorm.Model
    UserID uint
    PostID uint
}
