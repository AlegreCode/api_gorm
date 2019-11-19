package main

import (
	"log"

	"github.com/joho/godotenv"

	. "github.com/alegrecode/echo/api_gorm/controllers"

	. "github.com/alegrecode/echo/api_gorm/models"

	. "github.com/alegrecode/echo/api_gorm/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CONEXION DATABASE
	db := GetConexion()
	if err := db.DB().Ping(); err == nil {
		log.Println("Database is connected.")
	}
	defer db.Close()

	// MIGRATIONS
	db.DropTableIfExists(&Profile{}, &Comment{}, &Post{}, &User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Post{}, &Comment{}, &User{}, &Profile{})
	db.Model(&Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
	db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Profile{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	// ROUTES USERS
	e.GET("/api/user", GetAllUsers)

	e.GET("/api/user/:id", GetUser)

	e.POST("/api/user", SaveUser)

	e.PUT("/api/user/:id", UpdateUser)

	e.DELETE("/api/user/:id", DeleteUser)

	// ROUTES PROFILES
	e.GET("/api/profile", GetAllProfiles)

	e.GET("/api/profile/:userId", GetProfile)

	e.POST("/api/profile/:userId", SaveProfile)

	e.PUT("/api/profile/:userId", UpdateProfile)

	e.DELETE("/api/profile/:userId", DeleteProfile)

	// ROUTES POSTS
	e.GET("/api/post", GetAllPost)

	e.GET("/api/post/:id", GetPost)

	e.POST("/api/post/:id", SavePost)

	e.PUT("/api/post/:id", UpdatePost)

	e.DELETE("/api/post/:id", DeletePost)

	// ROUTES COMMENTS
	e.GET("/api/comment", GetAllComment)

	e.GET("/api/comment/:id", GetComment)

	e.POST("/api/comment/:userId/:postId", SaveComment)

	e.PUT("/api/comment/:id", UpdateComment)

	e.DELETE("/api/comment/:id", DeleteComment)

	e.Logger.Fatal(e.Start(":1323"))
}
