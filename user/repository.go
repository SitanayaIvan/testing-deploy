package user

import "gorm.io/gorm"

type Repository interface {
	GetAllUser() ([]User, error)
	GetUserById(id int) (User, error)
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) GetAllUser() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	return users, err
}

func (r repository) GetUserById(id int) (User, error) {
	var user User

	err := r.db.First(&user, id).Error
	return user, err
}

func (r repository) CreateUser(user User) error {
	err := r.db.Create(user).Error
	return err
}

func (r repository) UpdateUser(user User) error {
	err := r.db.Save(user).Error
	return err
}

func (r repository) DeleteUser(id int) error {
	err := r.db.Delete(User{}, id).Error
	return err
}
