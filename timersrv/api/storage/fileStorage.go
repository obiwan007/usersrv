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
	"time"

	pb "github.com/obiwan007/usersrv/proto"
)

type FileStorage struct{}

var timers []*pb.Timer
var maxId int = 0
var lock sync.Mutex

func NewFileStorage() *FileStorage {

	t := &FileStorage{}
	if err := t.Load("file.json", &timers); err != nil {
		log.Println(err)
	}
	maxId = len(timers)
	fmt.Println("Timers:", maxId)
	return t

}

func (t *FileStorage) Add(timer pb.Timer) pb.Timer {
	timer.Id = fmt.Sprint(maxId)
	maxId = maxId + 1

	timer.IsRunning = false
	timer.IsBilled = false
	timer.TimerStart = ""
	timer.TimerEnd = ""
	timer.ElapsedSeconds = 0

	timers = append(timers, &timer)
	fmt.Println("Adding to File:", timer)
	if err := t.Save("file.json", timers); err != nil {
		fmt.Println("Err", err)
	}
	return timer
}

func (t *FileStorage) Update(timer pb.Timer) *pb.Timer {
	old, _ := t.Get(timer.Id)

	old.Description = timer.Description
	old.Project = timer.Project
	old.Client = timer.Client
	old.IsBilled = timer.IsBilled
	old.Tags = timer.Tags

	if err := t.Save("file.json", timers); err != nil {
		fmt.Println("Err", err)
	}
	if err := t.Save("file.json", timers); err != nil {
		fmt.Println("Err", err)
	}
	return old
}

func (t *FileStorage) Delete(id string) error {
	return nil
}

// GetUser will return a user with a given Id
func (t *FileStorage) Get(id string) (*pb.Timer, error) {
	// idx := sort.Search(len(users), func(i int) bool {
	// 	return users[i].Id == id
	// })
	var res *pb.Timer = nil
	for _, u := range timers {
		if u.Id == id {
			res = u
			break
		}
	}
	fmt.Println("Found index", res)
	if res != nil {
		if res.IsRunning {
			res.ElapsedSeconds = int32(t.Elapsed(res))
			log.Println("Elapsed", res.ElapsedSeconds, t.Elapsed(res))
		}

		return res, nil
	}
	return &pb.Timer{}, errors.New("No such id found")
}

// Start will return a user with a given Id
func (t *FileStorage) Start(id string) (pb.Timer, error) {
	var res *pb.Timer
	for _, u := range timers {
		if u.Id == id {
			res = u
			break
		}
	}
	fmt.Println("Found index", res)
	if res != nil {
		res.IsRunning = true
		now := time.Now()
		res.TimerStart = now.Format(time.RFC3339)
		if err := t.Save("file.json", timers); err != nil {
			fmt.Println("Err", err)
		}
		return *res, nil
	}

	return pb.Timer{}, errors.New("No such id found")
}
func (t *FileStorage) Stop(id string) (pb.Timer, error) {
	var res *pb.Timer = nil
	for _, u := range timers {
		if u.Id == id {
			res = u
			break
		}
	}
	fmt.Println("Found index", res)
	if res != nil {
		res.IsRunning = false
		now := time.Now()
		res.TimerEnd = now.Format(time.RFC3339)
		res.ElapsedSeconds = int32(t.Elapsed(res))
		if err := t.Save("file.json", timers); err != nil {
			fmt.Println("Err", err)
		}
		return *res, nil
	}

	return pb.Timer{}, errors.New("No such id found")
}
func (t *FileStorage) Elapsed(timer *pb.Timer) float64 {
	if timer.IsRunning {
		tNow := time.Now()
		start, _ := time.Parse(time.RFC3339, timer.TimerStart)
		e := tNow.Sub(start).Seconds()
		log.Println("Elapsed:", e)
		log.Println("Now:", tNow.Format(time.RFC3339))
		log.Println("Start:", start.Format(time.RFC3339), timer.TimerStart)
		return e
	}
	start, _ := time.Parse(time.RFC3339, timer.TimerStart)
	end, _ := time.Parse(time.RFC3339, timer.TimerEnd)
	return end.Sub(start).Seconds()
}
func (t *FileStorage) GetAll() []*pb.Timer {
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