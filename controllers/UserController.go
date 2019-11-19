package controllers

import (
	"net/http"

	. "github.com/alegrecode/echo/api_gorm/db"
	. "github.com/alegrecode/echo/api_gorm/models"

	"github.com/labstack/echo"
)

func SaveUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	DB.Create(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The user has been saved successfuly.",
		"data": user,
	})
}

func GetAllUsers(c echo.Context) error {
	var users Users
	DB.Preload("Posts").Preload("Comments").Find(&users.Users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

func GetUser(c echo.Context) error {
	var user User
	id := c.Param("id")
	DB.Where("id = ?", id).Preload("Posts").Preload("Comments").Find(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}

func UpdateUser(c echo.Context) error {
	var user User
	id := c.Param("id")
	data := map[string]interface{}{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	DB.Model(&user).Where("id = ?", id).Update(User{Name:data["name"].(string), Age:int(data["age"].(float64))})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The user has been updated successfuly.",
		"data": user,
	})
}

func DeleteUser(c echo.Context) error {
	var user User
	id := c.Param("id")
	DB.Where("id = ?", id).Delete(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The user has been deleted successfuly.",
	})
}