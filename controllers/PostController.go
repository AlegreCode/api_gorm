package controllers

import (
	"net/http"

	. "github.com/alegrecode/echo/api_gorm/db"
	. "github.com/alegrecode/echo/api_gorm/models"

	"github.com/labstack/echo"
)

func SavePost(c echo.Context) error {
	var user User
	data := map[string]interface{}{}
	userId := c.Param("id")
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	DB.Find(&user, userId)
	post := Post{
		Title: data["title"].(string),
		Body: data["body"].(string),
		User: &user,
	}
	DB.Create(&post)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The post has been saved successfully.",
		"data": post,
	})
}

func GetAllPost(c echo.Context) error {
	var posts Posts
	DB.Preload("User").Preload("Comments").Find(&posts.Posts)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": posts,
	})
}

func GetPost(c echo.Context) error {
	var post Post
	id := c.Param("id")
	DB.Where("id = ?", id).Preload("User").Preload("Comments").Find(&post)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": post,
	})
}

func UpdatePost(c echo.Context) error {
	var post Post
	data:= map[string]interface{}{}
	id := c.Param("id")
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	DB.Model(&post).Where("id = ?", id).Update(Post{
		Title: data["title"].(string),
		Body: data["body"].(string),
	})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The post has been updated.",
		"data": post,
	})
}

func DeletePost(c echo.Context) error {
	var post Post
	id := c.Param("id")
	DB.Where("id = ?", id).Delete(&post)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The post has been deleted.",
	})
}
