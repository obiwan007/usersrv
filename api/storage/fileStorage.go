package userservicestorage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type FileStorage struct{}

var users []User
var maxId int = 0
var lock sync.Mutex

func NewFileStorage() *FileStorage {

	t := &FileStorage{}
	if err := t.Load("file.json", &users); err != nil {
		log.Fatalln(err)
	}
	maxId = len(users)
	fmt.Println("Max Length", maxId)
	return t

}

func (t *FileStorage) AddUser(user User) User {
	fmt.Println("Adding to File:", user)
	user.Id = maxId
	maxId = maxId + 1
	users = append(users, user)
	t.ListUser()
	if err := t.Save("file.json", users); err != nil {
		fmt.Println("Err", err)
	}
	return user
}

func (t *FileStorage) DeleteUser(user User) error {
	return nil
}

func (t *FileStorage) GetUser(id int) (User, error) {
	// idx := sort.Search(len(users), func(i int) bool {
	// 	return users[i].Id == id
	// })
	var res *User = nil
	for _, u := range users {
		if u.Id == id {
			res = &u
			break
		}
	}
	fmt.Println("Found index", res)
	if res != nil {
		return *res, nil
	}
	return User{}, errors.New("No such id found")

}

func (t *FileStorage) ListUser() []User {
	// for _, u := range users {
	// 	fmt.Println(u)
	// }
	return users
}

// Save saves a representation of v to the file at path.
func (t *FileStorage) Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func (t *FileStorage) Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
