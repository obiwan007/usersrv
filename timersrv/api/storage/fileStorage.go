package timerservicestorage

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/timersrv/api/storage/ent"
	"github.com/obiwan007/usersrv/timersrv/api/storage/ent/timer"

	claims "github.com/obiwan007/usersrv/pkg/claims"
)

type FileStorage struct{ Db *ent.Client }

func NewFileStorage(dbconnection string) *FileStorage {
	log.Println("DB Connection", dbconnection)
	client, err := ent.Open("postgres", dbconnection)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	log.Println("DB connected")

	// run the auto migration tool.
	if err := client.Debug().Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	t := &FileStorage{Db: client}
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
func (t *FileStorage) GetAll(ctx context.Context, dayRange int32, c *claims.MyCustomClaims) ([]*pb.Timer, error) {
	// for _, u := range users {
	// 	fmt.Println(u)
	// }
	log.Println("Dayrange", dayRange)
	if dayRange == -1 {
		dayRange = 1000
	}

	dayRangeDate := time.Now().AddDate(0, 0, int(-dayRange))
	log.Println("Dayrange", dayRange, dayRangeDate)
	existingTimer, err := t.Db.Timer.
		Query().Where(timer.And(timer.Userid(c.Subject), timer.Or(timer.IsRunning(true), timer.TimerStartGTE(dayRangeDate)))).All(ctx)

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
