package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func listComments(c echo.Context) error {
    return c.String(http.StatusOK, "List Comments")
}

func findComment(c echo.Context) error {
    return c.String(http.StatusOK, "List Comments")
}

func createComments(c echo.Context) error {
    return c.String(http.StatusOK, "List Comments")
}

func updateComments(c echo.Context) error {
    return c.String(http.StatusOK, "List Comments")
}

func destroyComments(c echo.Context) error {
    return c.String(http.StatusOK, "List Comments")
}

func CommentController(g *echo.Group) {
    g.GET("", listComments)
    g.GET(":id", findComment)
    g.POST("", createComments)
    g.PUT("", updateComments)
    g.DELETE(":id", destroyComments)
}
