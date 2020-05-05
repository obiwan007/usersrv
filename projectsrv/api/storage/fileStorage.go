package timerservicestorage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	pb "github.com/obiwan007/usersrv/proto"
)

type FileStorage struct{}

var timers []*pb.Project
var maxId int = 0
var lock sync.Mutex

func NewFileStorage() *FileStorage {

	t := &FileStorage{}
	if err := t.Load("file.json", &timers); err != nil {
		log.Println(err)
	}
	maxId = len(timers)
	log.Println("Timers:", maxId)
	return t

}

func (t *FileStorage) Add(timer pb.Project) *pb.Project {
	timer.Id = fmt.Sprint(maxId)
	maxId = maxId + 1

	timers = append(timers, &timer)
	log.Println("Adding to File:", timer)
	if err := t.Save("file.json", timers); err != nil {
		log.Println("Err", err)
	}
	return &timer
}

func (t *FileStorage) Update(timer pb.Project) *pb.Project {
	old, _ := t.Get(timer.Id)

	old.Name = timer.Name
	old.Description = timer.Description
	old.Tags = timer.Tags
	old.Status = timer.Status
	old.Team = timer.Team
	old.Client = timer.Client

	if err := t.Save("file.json", timers); err != nil {
		log.Println("Err", err)
	}
	if err := t.Save("file.json", timers); err != nil {
		log.Println("Err", err)
	}
	return old
}

func (t *FileStorage) Delete(id string) error {
	return nil
}

// GetUser will return a user with a given Id
func (t *FileStorage) Get(id string) (*pb.Project, error) {
	// idx := sort.Search(len(users), func(i int) bool {
	// 	return users[i].Id == id
	// })
	var res *pb.Project = nil
	for _, u := range timers {
		if u.Id == id {
			res = u
			break
		}
	}
	fmt.Println("Found index", res)
	if res != nil {

		return res, nil
	}
	return &pb.Project{}, errors.New("No such id found")
}

// Start will return a user with a given Id

func (t *FileStorage) GetAll() []*pb.Project {
	// for _, u := range users {
	// 	fmt.Println(u)
	// }
	return timers
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
