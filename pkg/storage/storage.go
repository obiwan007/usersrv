package userservicestorage

type Storage interface {
	AddUser(user User) User
	DeleteUser(user User) error
	ListUser() []User
	GetUser(id int) (User, error)
}
