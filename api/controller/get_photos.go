package controller

import (
	"api/database"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetPhotos(c echo.Context, db *gorm.DB) (err error) {
	request := struct {
		TwitterID int `json:"twitter_id"`
	}{}

	if err = c.Bind(&request); err != nil {
		return
	}

	photos, err := database.Photos(request.TwitterID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"photos": photos,
	})

	return
}
