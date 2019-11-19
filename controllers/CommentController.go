package controllers

import (
	"net/http"

	. "github.com/alegrecode/echo/api_gorm/db"
	. "github.com/alegrecode/echo/api_gorm/models"

	"github.com/labstack/echo"
)

func SaveComment(c echo.Context) error {
	var user User
	var post Post
	data := map[string]interface{}{}
	userId := c.Param("userId")
	postId := c.Param("postId")
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	DB.Find(&user, userId)
	DB.Find(&post, postId)
	comment := Comment{
		Body: data["body"].(string),
		User: &user,
		Post: &post,
	}
	DB.Create(&comment)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The comment has been saved successfully.",
		"data": comment,
	})
}

func GetAllComment(c echo.Context) error {
	var comments Comments
	DB.Preload("User").Preload("Post").Find(&comments.Comments)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": comments,
	})
}

func GetComment(c echo.Context) error {
	var comment Comment
	id := c.Param("id")
	DB.Where("id = ?", id).Preload("User").Preload("Post").Find(&comment)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": comment,
	})
}

func UpdateComment(c echo.Context) error {
	var comment Comment
	data:= map[string]interface{}{}
	id := c.Param("id")
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	DB.Model(&comment).Where("id = ?", id).Update(Comment{
		Body: data["body"].(string),
	})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The comment has been updated.",
		"data": comment,
	})
}

func DeleteComment(c echo.Context) error {
	var comment Comment
	id := c.Param("id")
	DB.Where("id = ?", id).Delete(&comment)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The comment has been deleted.",
	})
}
