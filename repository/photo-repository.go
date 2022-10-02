package repository

import (
	"WebApiGo/models"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	InsertPhoto(p models.Photos) models.Photos
	UpdatePhoto(p models.Photos) models.Photos
	DeletePhoto(p models.Photos)
	AllPhoto() []models.Photos
	FindPhotoByID(photoID uint64) models.Photos
}

type photoConnection struct {
	connection *gorm.DB
}

func NewPhotoRepository(dbConn *gorm.DB) PhotoRepository {
	return &photoConnection{
		connection: dbConn,
	}
}

func (db *photoConnection) InsertPhoto(p models.Photos) models.Photos {
	db.connection.Save(&p)
	db.connection.Preload("User").Find(&p)
	return p
}

func (db *photoConnection) UpdatePhoto(p models.Photos) models.Photos {
	db.connection.Save(&p)
	db.connection.Preload("User").Find(&p)
	return p
}

func (db *photoConnection) DeletePhoto(p models.Photos) {
	db.connection.Delete(&p)
}

func (db *photoConnection) FindPhotoByID(photoID uint64) models.Photos {
	var photo models.Photos
	db.connection.Preload("User").Find(&photo, photoID)
	return photo
}

func (db *photoConnection) AllPhoto() []models.Photos {
	var photos []models.Photos
	db.connection.Preload("User").Find(&photos)
	return photos
}
