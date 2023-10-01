package controllers

import (
	"net/http"
	"strings"

	"github.com/bagasjs/go-blog/helper"
	"github.com/bagasjs/go-blog/models"
	"github.com/labstack/echo/v4"
)

func listUsers(c echo.Context) error {
    cc := c.(helper.CustomContext)
    var users []models.User
    result := cc.Database.Find(&users)

    if result.Error != nil {
        return cc.JSON(http.StatusNotFound, helper.ResponseErr("Failed to fetch all users"))
    }

    var userResources []models.UserResource
    for _, user := range users {
        userResources = append(userResources, user.Resource())
    }

    return cc.JSON(http.StatusOK, helper.ResponseOK("All users fetched", userResources))
}

func findUser(c echo.Context) error {
    cc := c.(helper.CustomContext)
    user := models.User{}

    result := cc.Database.Where(map[string]interface{}{ "ID" : cc.Param("id")  }).First(&user)
    if result.Error != nil {
        return c.JSON(http.StatusBadRequest, helper.ResponseErr(result.Error.Error()))
    }
    return c.JSON(http.StatusOK, helper.ResponseOK("User fetched", user.Resource()))
}

func createUsers(c echo.Context) error {
    cc := c.(helper.CustomContext)

    name := c.FormValue("name")
    password := c.FormValue("password")
    passwordConfirmation := c.FormValue("password_confirmation")
    email := c.FormValue("email")

    if strings.Compare(password, passwordConfirmation) != 0 {
        return c.JSON(http.StatusBadRequest, helper.ResponseErr("Password is not confirmed"))
    }

    password = helper.HashMake(password)
    user := models.User{ Name: name, Password: password, Email: email }
    result := cc.Database.Create(&user)
    if result.Error != nil {
        return c.JSON(http.StatusBadGateway, helper.ResponseErr("Failed to create user in database"))
    }

    cc.Database.Save(&user)

    return c.JSON(http.StatusOK, helper.ResponseOK("User creation success", user.Resource()))
}

func updateUsers(c echo.Context) error {
    cc := c.(helper.CustomContext)
    user := models.User{}
    result := cc.Database.Where(map[string]interface{}{ "ID" : cc.Param("id")  }).First(&user)
    if result.Error != nil {
        return c.JSON(http.StatusBadRequest, helper.ResponseErr(result.Error.Error()))
    }

    user.Name = helper.ChooseString(c.FormValue("name"), user.Name)
    user.Email = helper.ChooseString(c.FormValue("email"), user.Email)

    password := c.FormValue("password")
    if len(password) != 0 {
        if strings.Compare(password, c.FormValue("password_confirmation")) != 0 {
            return c.JSON(http.StatusBadRequest, helper.ResponseErr("Password is not confirmed"))
        }
        user.Password = password
    }

    cc.Database.Save(&user)

    return c.JSON(http.StatusOK, helper.ResponseOK("User updated", user.Resource()))
}

func destroyUsers(c echo.Context) error {
    cc := c.(helper.CustomContext)
    cc.Database.Delete(&models.User{}, cc.Param("id"))
    return c.JSON(http.StatusOK, helper.ResponseOK("User is destroyed", nil))
}

func UserController(g *echo.Group) {
    g.GET("", listUsers)
    g.GET("/:id", findUser)
    g.POST("", createUsers)
    g.PUT("/:id", updateUsers)
    g.DELETE("/:id", destroyUsers)
}
