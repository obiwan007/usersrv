package userservicestorage

import "fmt"

type DBStorage struct{}

func NewDBStorage() *DBStorage {
	return &DBStorage{}
}

func (t *DBStorage) AddUser(user User) {
	fmt.Println("Adding to DB:", user)
}

func (t *DBStorage) DeleteUser(user User) {
}

func (t *DBStorage) GetUser(id string) {
}

func (t *DBStorage) ListUser() {
	for _, u := range users {
		fmt.Println(u)
	}
}
