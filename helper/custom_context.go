package helper

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomContext struct {
    echo.Context
    Database *gorm.DB
}


