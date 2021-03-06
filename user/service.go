package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

//mapping struct input ke struct User
//simpan struct User malalui repository

/*	email dari user/input di mapping ke struct CheckEmailInput*/
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email //menampung email dari user

	user, err := s.repository.FindByEmail(email) // cari email melalui repository berdasarkan email diinput
	if err != nil {
		return false, err
	}

	// jika user tidak ditemukan berarti benar email masih available
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	// dapartkan user berdasarkan ID
	// user update attribute avatar file name
	// simpan perubahan file name

	// temukan user
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	// simpan perubahan
	user.AvatarFileName = fileLocation

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}
	return updateUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	//panggil dari repository find by id
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found with that ID")
	}

	return user, nil
}
