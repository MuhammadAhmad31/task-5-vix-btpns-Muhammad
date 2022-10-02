package helpers

import (
	"WebApiGo/app"
	"WebApiGo/models"
	"WebApiGo/repository"
	"github.com/mashingan/smapping"
	"log"
)

type UserHelpers interface {
	Update(user app.UserUpdateApp) models.User
	Profile(userID string) models.User
}

type userHelpers struct {
	userRepository repository.UserRepository
}

func NewUserHelpers(userRepo repository.UserRepository) UserHelpers {
	return &userHelpers{
		userRepository: userRepo,
	}
}

func (helpers *userHelpers) Update(user app.UserUpdateApp) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updateUser := helpers.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (helpers *userHelpers) Profile(userID string) models.User {
	return helpers.userRepository.ProfileUser(userID)
}
