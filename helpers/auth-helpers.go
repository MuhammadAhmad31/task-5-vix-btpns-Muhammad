package helpers

import (
	"WebApiGo/app"
	"WebApiGo/models"
	"WebApiGo/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthHelpers interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user app.RegisterApp) models.User
	FindByEmail(enail string) models.User
	IsDuplicateEmail(email string) bool
}

type authHelpers struct {
	userRepository repository.UserRepository
}

func NewAuthHelpers(userRep repository.UserRepository) AuthHelpers {
	return &authHelpers{
		userRepository: userRep,
	}
}

func (helpers *authHelpers) VerifyCredential(email string, password string) interface{} {
	res := helpers.userRepository.VerifyCredential(email, password)
	if v, ok := res.(models.User); ok {
		comparePassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparePassword {
			return res
		}
		return false
	}
	return false
}

func (helpers *authHelpers) CreateUser(user app.RegisterApp) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := helpers.userRepository.InsertUser(userToCreate)
	return res
}

func (helpers *authHelpers) FindByEmail(email string) models.User {
	return helpers.userRepository.FindByEmail(email)
}

func (helpers *authHelpers) IsDuplicateEmail(email string) bool {
	res := helpers.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
