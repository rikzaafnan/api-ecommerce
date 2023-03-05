package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	RegisterUser(input RegisterUserInput) (User, error)
	RegisterUserByPassSuperAdmin(input RegisterUserInput) (User, error)
	Login(input Logininput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	// SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(Id int) (User, error)
	Verification(email string) (User, error)
}

type serviceUserImpl struct {
	repositoryUser RepositoryUser
}

func NewServiceUser(repositoryUser RepositoryUser) *serviceUserImpl {
	return &serviceUserImpl{repositoryUser}
}

func (s *serviceUserImpl) RegisterUser(input RegisterUserInput) (User, error) {

	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(PasswordHash)
	user.Role = "user"
	user.IsVerification = 0

	// check Email
	userEmail, err := s.repositoryUser.FindByEmail(user.Email)
	if err != nil {
		return User{}, err
	}

	if userEmail.ID != 0 {
		return User{}, errors.New("email telah digunakan")
	}

	newUser, err := s.repositoryUser.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *serviceUserImpl) Login(input Logininput) (User, error) {
	var user User

	email := input.Email
	password := input.Password

	user, err := s.repositoryUser.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	if user.IsVerification == 0 {
		return user, errors.New("user belum verification")
	}

	return user, nil

}

func (s *serviceUserImpl) IsEmailAvailable(input CheckEmailInput) (bool, error) {

	email := input.Email

	user, err := s.repositoryUser.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

func (s *serviceUserImpl) RegisterUserByPassSuperAdmin(input RegisterUserInput) (User, error) {

	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(PasswordHash)
	user.Role = "admin"
	user.IsVerification = 1

	// check Email
	userEmail, err := s.repositoryUser.FindByEmail(user.Email)
	if err != nil {
		return User{}, err
	}

	if userEmail.ID != 0 {
		return User{}, errors.New("email telah digunakan")
	}

	newUser, err := s.repositoryUser.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *serviceUserImpl) GetUserByID(Id int) (User, error) {

	var user User

	user, err := s.repositoryUser.FindByID(Id)
	if err != nil {
		return user, err
	}

	// check if user id == 0
	if user.ID == 0 {
		return user, errors.New("no user found on the ID")
	}

	return user, nil
}

func (s *serviceUserImpl) Verification(email string) (User, error) {

	var user User

	user, err := s.repositoryUser.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on the ID")
	}

	if user.IsVerification == 1 {
		return user, errors.New("user has been verification")
	}

	err = s.repositoryUser.PatchIsVerification(email)
	if err != nil {
		return user, err
	}

	return user, nil

}

/*
// func (s *serviceUserImpl) SaveAvatar(ID int, fileLocation string) (User, error) {
// 	// mencari user
// 	user, err := s.repositoryUser.FindByID(ID)
// 	if err != nil {
// 		return user, err
// 	}

// 	// check if user id == 0
// 	if user.ID == 0 {
// 		return user, errors.New("no user found on the ID")
// 	}

// 	// end mencari user

// 	// simpan perubhaan ke db
// 	user.AvatarFileName = fileLocation

// 	updatedUser, err := s.repositoryUser.Update(user)
// 	if err != nil {

// 		return user, err

// 	}

// 	return updatedUser, nil

// }

*/
