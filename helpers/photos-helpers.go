package helpers

import (
	"WebApiGo/app"
	"WebApiGo/models"
	"WebApiGo/repository"
	"fmt"
	"github.com/mashingan/smapping"
	"log"
)

type PhotoHelpers interface {
	Insert(b app.PhotosCreateApp) models.Photos
	Update(b app.PhotosUpdateApp) models.Photos
	Delete(b models.Photos)
	All() []models.Photos
	FindByID(bookID uint64) models.Photos
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type photoHelpers struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoHelpers(photoRepo repository.PhotoRepository) PhotoHelpers {
	return &photoHelpers{
		photoRepository: photoRepo,
	}
}

func (helpers *photoHelpers) Insert(b app.PhotosCreateApp) models.Photos {
	photo := models.Photos{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := helpers.photoRepository.InsertPhoto(photo)
	return res
}

func (helpers *photoHelpers) Update(b app.PhotosUpdateApp) models.Photos {
	photo := models.Photos{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := helpers.photoRepository.UpdatePhoto(photo)
	return res
}

func (helpers *photoHelpers) Delete(b models.Photos) {
	helpers.photoRepository.DeletePhoto(b)
}

func (helpers *photoHelpers) All() []models.Photos {
	return helpers.photoRepository.AllPhoto()
}

func (helpers *photoHelpers) FindByID(bookID uint64) models.Photos {
	return helpers.photoRepository.FindPhotoByID(bookID)
}

func (helpers *photoHelpers) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := helpers.photoRepository.FindPhotoByID(bookID)
	id := fmt.Sprintf("%v", b.UserId)
	return userID == id
}
