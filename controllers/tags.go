package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func listTags(c echo.Context) error {
    return c.String(http.StatusOK, "List Tags")
}

func findTag(c echo.Context) error {
    return c.String(http.StatusOK, "List Tags")
}

func createTags(c echo.Context) error {
    return c.String(http.StatusOK, "List Tags")
}

func updateTags(c echo.Context) error {
    return c.String(http.StatusOK, "List Tags")
}

func destroyTags(c echo.Context) error {
    return c.String(http.StatusOK, "List Tags")
}

func TagController(g *echo.Group) {
    g.GET("", listTags)
    g.GET(":id", findTag)
    g.POST("", createTags)
    g.PUT("", updateTags)
    g.DELETE(":id", destroyTags)
}
