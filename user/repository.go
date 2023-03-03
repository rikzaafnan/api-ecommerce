package user

import "gorm.io/gorm"

type RepositoryUser interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repositoryUserImpl struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUserImpl {
	return &repositoryUserImpl{db}
}

func (r *repositoryUserImpl) Save(user User) (User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repositoryUserImpl) FindByEmail(email string) (User, error) {

	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositoryUserImpl) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {

		return user, err
	}

	return user, nil
}

func (r *repositoryUserImpl) Update(user User) (User, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}
