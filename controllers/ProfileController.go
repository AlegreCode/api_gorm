package controllers

import (
	"net/http"
	. "github.com/alegrecode/echo/api_gorm/db"
	. "github.com/alegrecode/echo/api_gorm/models"
	"github.com/labstack/echo"
)

func SaveProfile (c echo.Context) error {
	var user User
	data := map[string]interface{}{}
	userId := c.Param("userId")
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	DB.Find(&user, userId)
	profile := Profile{
		Username: data["username"].(string),
		Email: data["email"].(string),
		User: &user,
	}
	DB.Create(&profile)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The profile has been created.",
		"data": profile,
	})
}

func GetAllProfiles(c echo.Context) error {
	var profiles Profiles
	DB.Preload("User.Comments").Preload("User.Posts").Find(&profiles.Profiles)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": profiles,
	})
}

func GetProfile(c echo.Context) error {
	var profile Profile
	userId := c.Param("userId")
	DB.Where("user_id = ?", userId).Preload("User.Comments").Preload("User.Posts").Find(&profile)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": profile,
	})
}

func UpdateProfile(c echo.Context) error {
	var profile Profile
	data := map[string]interface{}{}
	id := c.Param("userId")
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	DB.Model(&profile).Where("id = ?", id).Preload("User").Update(Profile{
		Username: data["username"].(string),
		Email: data["email"].(string),
	})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The profile has been udpated.",
		"data": profile,
	})
}

func DeleteProfile(c echo.Context) error {
	var profile Profile
	id := c.Param("userId")
	DB.Where("id = ?", id).Delete(&profile)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "The profile has been deleted.",
	})
}