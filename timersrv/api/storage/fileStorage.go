package timerservicestorage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/timersrv/api/storage/ent"
	"github.com/obiwan007/usersrv/timersrv/api/storage/ent/timer"

	claims "github.com/obiwan007/usersrv/pkg/claims"
)

type FileStorage struct{ Db *ent.Client }

var timers []*pb.Timer
var maxId int = 0
var lock sync.Mutex

func NewFileStorage(dbconnection string) *FileStorage {
	log.Println("DB Connection", dbconnection)
	client, err := ent.Open("postgres", dbconnection)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	log.Println("DB connected")

	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	t := &FileStorage{Db: client}
	if err := t.Load("file.json", &timers); err != nil {
		log.Println(err)
	}
	maxId = len(timers)
	fmt.Println("Timers:", maxId)
	return t

}

func (t *FileStorage) Add(ctx context.Context, timer *pb.Timer, c *claims.MyCustomClaims) (*pb.Timer, error) {
	timer.IsRunning = false
	timer.IsBilled = false
	timer.TimerStart = ""
	timer.TimerEnd = ""
	timer.ElapsedSeconds = 0

	newtimer, err := t.Db.Timer. // UserClient.
					Create().                          // User create builder.
					SetDescription(timer.Description). // Set field value.
					SetProjectid(timer.Project).       // Set field value.
					SetMandantid(c.Mandant).
					SetUserid(c.Subject).
					Save(ctx) // Create and return.
	if err != nil {
		return nil, err
	}
	response := &pb.Timer{Id: strconv.Itoa(newtimer.ID),
		Description: newtimer.Description,
		Project:     newtimer.Projectid,
	}

	//timers = append(timers, &timer)
	fmt.Println("Adding to File:", timer)
	// if err := t.Save("file.json", timers); err != nil {
	// 	fmt.Println("Err", err)
	// }
	return response, nil
}

func (t *FileStorage) Update(ctx context.Context, timerEntity *pb.Timer, c *claims.MyCustomClaims) (*pb.Timer, error) {
	id, _ := strconv.Atoi(timerEntity.Id)

	existingTimer, err := t.Db.Timer.
		Query().
		Where(timer.And(timer.ID(id), timer.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	newtimer, err := existingTimer.
		Update().
		SetNillableDescription(checkNil(timerEntity.Description)).
		SetNillableProjectid(checkNil(timerEntity.Project)).
		SetNillableTags(checkNil(timerEntity.Tags)).
		SetNillableIsBilled(&timerEntity.IsBilled).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	response := toPb(newtimer)
	return response, nil
}

func (t *FileStorage) Delete(ctx context.Context, timerid string, c *claims.MyCustomClaims) (*pb.Timer, error) {
	id, _ := strconv.Atoi(timerid)
	existingTimer, err := t.Db.Timer.
		Query().
		Where(timer.And(timer.ID(id), timer.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	t.Db.Timer.DeleteOne(existingTimer).Exec(ctx)

	return nil, err
}

// GetUser will return a user with a given Id
func (t *FileStorage) Get(ctx context.Context, timerid string, c *claims.MyCustomClaims) (*pb.Timer, error) {
	// idx := sort.Search(len(users), func(i int) bool {
	// 	return users[i].Id == id
	// })

	id, _ := strconv.Atoi(timerid)
	newtimer, err := t.Db.Timer.
		Query().
		Where(timer.And(timer.ID(id), timer.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	response := toPb(newtimer)
	return response, nil
}

// Start will return a user with a given Id
func (t *FileStorage) Start(ctx context.Context, timerid string, c *claims.MyCustomClaims) (*pb.Timer, error) {
	id, _ := strconv.Atoi(timerid)

	existingTimer, err := t.Db.Timer.
		Query().
		Where(timer.And(timer.ID(id), timer.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	newtimer, err := existingTimer.
		Update().
		SetTimerStart(time.Now()).
		SetIsRunning(true).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	response := toPb(newtimer)
	return response, nil
}
func (t *FileStorage) Stop(ctx context.Context, timerid string, c *claims.MyCustomClaims) (*pb.Timer, error) {
	id, _ := strconv.Atoi(timerid)
	tEnd := time.Now()
	existingTimer, err := t.Db.Timer.
		Query().
		Where(timer.And(timer.ID(id), timer.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("Elapsed", tEnd.Sub(existingTimer.TimerStart).Seconds())
	log.Println("Elapsed", int(tEnd.Sub(existingTimer.TimerStart).Seconds()))
	newtimer, err := existingTimer.
		Update().
		SetElapsedSeconds(int(tEnd.Sub(existingTimer.TimerStart).Seconds())).
		SetTimerEnd(tEnd).
		SetIsRunning(false).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	response := toPb(newtimer)
	return response, nil
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
func (t *FileStorage) GetAll(ctx context.Context, c *claims.MyCustomClaims) ([]*pb.Timer, error) {
	// for _, u := range users {
	// 	fmt.Println(u)
	// }
	existingTimer, err := t.Db.Timer.
		Query().Where(timer.Userid(c.Subject)).All(ctx)

	if err != nil {
		return nil, err
	}
	var list []*pb.Timer
	for _, res := range existingTimer {
		t := toPb(res)
		list = append(list, t)
	}
	return list, err
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

func toPb(newtimer *ent.Timer) *pb.Timer {
	response := &pb.Timer{Id: strconv.Itoa(newtimer.ID),
		Description:    newtimer.Description,
		Project:        newtimer.Projectid,
		Tags:           newtimer.Tags,
		TimerStart:     newtimer.TimerStart.Format(time.RFC3339),
		TimerEnd:       newtimer.TimerEnd.Format(time.RFC3339),
		IsRunning:      newtimer.IsRunning,
		IsBilled:       newtimer.IsBilled,
		ElapsedSeconds: int32(newtimer.ElapsedSeconds),
	}
	return response
}
func checkNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
