package models

import "gorm.io/gorm"

type PostTag struct {
	gorm.Model
    PostID uint
    TagID uint
}
