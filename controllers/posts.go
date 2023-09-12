package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func listPosts(c echo.Context) error {
    return c.String(http.StatusOK, "List Posts")
}

func findPost(c echo.Context) error {
    return c.String(http.StatusOK, "List Posts")
}

func createPosts(c echo.Context) error {
    return c.String(http.StatusOK, "List Posts")
}

func updatePosts(c echo.Context) error {
    return c.String(http.StatusOK, "List Posts")
}

func destroyPosts(c echo.Context) error {
    return c.String(http.StatusOK, "List Posts")
}

func PostController(g *echo.Group) {
    g.GET("", listPosts)
    g.GET(":id", findPost)
    g.POST("", createPosts)
    g.PUT("", updatePosts)
    g.DELETE(":id", destroyPosts)
}
