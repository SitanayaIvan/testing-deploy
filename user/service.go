package user

type Service interface {
	GetAllUser() ([]User, error)
	GetUserById(id int) (User, error)
	CreateUser(user User) error
	UpdateUser(user User, id int) error
	DeleteUser(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s service) GetAllUser() ([]User, error) {
	users, err := s.repository.GetAllUser()
	return users, err
}

func (s service) GetUserById(id int) (User, error) {
	user, err := s.repository.GetUserById(id)
	return user, err
}

func (s service) CreateUser(user User) error {
	err := s.repository.CreateUser(user)
	return err
}

func (s service) UpdateUser(user User, id int) error {
	userUpdate, err := s.repository.GetUserById(id)
	if err != nil {
		return err
	}

	userUpdate.FirstName = user.FirstName
	userUpdate.LastName = user.LastName
	userUpdate.Gender = user.Gender
	userUpdate.Address = user.Address
	userUpdate.Dob = user.Dob
	userUpdate.Image = user.Image

	err = s.repository.UpdateUser(userUpdate)
	return err
}

func (s service) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)
	return err
}
