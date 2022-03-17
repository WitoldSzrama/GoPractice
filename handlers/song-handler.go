package handlers

import (
	"errors"
	"practiceTwo/database"
	"practiceTwo/entities"
	"strconv"

	_"github.com/swaggo/swag/example/celler/httputil"
	"github.com/gin-gonic/gin"
)

// ShowSongs godoc
// @Summary      Show all songs
// @Tags         songs
// @Accept       json
// @Produce      json
// @Success      200  {array}  entities.Song
// @Router       / [get]
func GetSongs() []entities.Song {
	songs := []entities.Song{}
	database.DB.Find(&songs)

	return songs
}
// ShowSongs godoc
// @Summary      Show song by id
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  entities.Song
// @Failure      400  {object}  httputil.HTTPError
// @Router       /{id} [get]
func GetSong(c *gin.Context) (entities.Song, error) {
	song := entities.Song{}
	id, err := checkId(c)
	if err != nil {
		return song, err
	}
	DBRes := database.DB.First(&song, id)

	if DBRes.Error != nil {
		return song, DBRes.Error 
	}

	return song, nil
}
// ShowSongs godoc
// @Summary      Update song by id
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      201  {object}  entities.Song
// @Failure      400  {object}  httputil.HTTPError
// @Router       /{id} [put]
func UpdateSong(c *gin.Context) (entities.Song, error) {
	update := entities.Song{}
	err := c.BindJSON(&update)
	if err != nil {
		return update, err
	}

	song, err := GetSong(c)
	if err != nil {
		return song, err
	}

	DBRes := database.DB.Model(&song).Updates(update)
	if DBRes.Error != nil {
		return song, DBRes.Error 
	}
	return song, nil
}

// ShowSongs godoc
// @Summary      Delete song by Id
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      204  
// @Failure      400  {object}  httputil.HTTPError
// @Router       /{id} [delete]
func DeleteSong(c *gin.Context) error {
	song := entities.Song{}
	id, err := checkId(c)
	if err != nil {
		return err
	}

	DBRes := database.DB.Delete(&song, id)
	if DBRes.Error != nil {
		return DBRes.Error 
	}

	return nil
}

// ShowSongs godoc
// @Summary      Add new song
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      201  {object}  entities.Song
// @Failure      400  {object}  httputil.HTTPError
// @Router       /{id} [post]
func AddSong(c *gin.Context) (entities.Song, error) {
	song := entities.Song{}
	err := c.BindJSON(&song)
	if err != nil {
		return song, err
	}
	DBRes := database.DB.Create(&song)
	if DBRes.Error != nil {
		return song, DBRes.Error
	}

	return song, nil
}

func checkId(c *gin.Context) (int, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return id, errors.New("wrong id provided")
	}
	return id, nil
}