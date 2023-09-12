package main

import (
	"net/http"
	"os"
	"path"

	"github.com/bagasjs/go-blog/controllers"
	"github.com/bagasjs/go-blog/helper"
	"github.com/bagasjs/go-blog/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

    cmd := &cobra.Command {
        Use: "go-blog",
        Short: "Go-Blog is a simple blogging and article website",
        Long: "Go-Blog is a simple blogging and article website. Go-Blog is also configurable",
        Run: func(cmd *cobra.Command, args []string) {
        },
    }

    serve := &cobra.Command{
        Use: "serve",
        Short: "Server",
        Run: func(cmd *cobra.Command, args []string) {
            Serve()
        },
    }

    setup := &cobra.Command{
        Use: "setup",
        Short: "Setup the environment",
        Long: "Create a resource directory, database migrations and etc.",
        Run: func(cmd *cobra.Command, args []string) {
            var err error
            var cwd string
            if cwd, err = os.Getwd(); err != nil {
                panic("Failed to setup your application")
            }
            if err = os.Mkdir(path.Join(cwd, "res"), os.ModePerm); err != nil {
                panic("Failed to setup your application")
            }

            db, err := gorm.Open(sqlite.Open(path.Join(cwd, "res", "db.sqlite3")), &gorm.Config{})
            if err != nil {
                panic("Failed to open database connection")
            }

            db.AutoMigrate(&models.User{})
            db.AutoMigrate(&models.Post{})
            db.AutoMigrate(&models.Tag{})
            db.AutoMigrate(&models.Comment{})
            db.AutoMigrate(&models.UserLikePost{})
            db.AutoMigrate(&models.PostTag{})
        },
    }
    cmd.AddCommand(serve)
    cmd.AddCommand(setup)
    cmd.Execute()
}

func Serve() {
    e := echo.New()
    db, err := gorm.Open(sqlite.Open("./res/db.sqlite3"), &gorm.Config{})
    if err != nil {
        panic("Failed to open database connection")
    }
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            cc := helper.CustomContext{ 
                Context: c,
                Database: db,
            }
            return next(cc)
        }
    })

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World")
    })

    var g *echo.Group

    g = e.Group("/api/users")
    controllers.UserController(g)
    g = e.Group("/api/posts")
    controllers.PostController(g)
    g = e.Group("/api/comments")
    controllers.CommentController(g)
    g = e.Group("/api/tags")
    controllers.TagController(g)

    e.Logger.Fatal(e.Start(":3000"))
}
