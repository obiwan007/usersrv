package timerservicestorage

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	_ "github.com/lib/pq"
	"github.com/obiwan007/usersrv/projectsrv/api/storage/ent"
	"github.com/obiwan007/usersrv/projectsrv/api/storage/ent/project"
	pb "github.com/obiwan007/usersrv/proto"

	claims "github.com/obiwan007/usersrv/pkg/claims"
)

type FileStorage struct{ Db *ent.Client }

var timers []*pb.Project
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

	return t
}

func (t *FileStorage) Add(ctx context.Context, prj pb.Project, c *claims.MyCustomClaims) (*pb.Project, error) {

	newprj, err := t.Db.Project. // UserClient.
					Create().                        // User create builder.
					SetName(prj.Name).               // Set field value.
					SetDescription(prj.Description). // Set field value.
					SetMandantid(c.Mandant).
					SetUserid(c.Subject).
					SetStatus(prj.Status).
					SetTeam(prj.Team).
					SetTags(prj.Tags).
					SetStatus(prj.Tags).
					Save(ctx) // Create and return.
	if err != nil {
		return nil, err
	}
	response := &pb.Project{Id: strconv.Itoa(newprj.ID),
		Description: newprj.Description,
		Name:        newprj.Name,
	}

	//timers = append(timers, &timer)
	fmt.Println("Adding to File:", prj)
	// if err := t.Save("file.json", timers); err != nil {
	// 	fmt.Println("Err", err)
	// }

	return response, err
}

func (t *FileStorage) Update(ctx context.Context, prj *pb.Project, c *claims.MyCustomClaims) (*pb.Project, error) {
	id, _ := strconv.Atoi(prj.Id)

	existingTimer, err := t.Db.Project.
		Query().
		Where(project.And(project.ID(id), project.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	new, err := existingTimer.
		Update().
		SetNillableDescription(checkNil(prj.Description)).
		SetNillableClientid(checkNil(prj.Client)).
		SetNillableTags(checkNil(prj.Tags)).
		SetNillableStatus(&prj.Status).
		SetNillableName(&prj.Name).
		SetNillableTeam(&prj.Team).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	response := toPb(new)
	return response, err
}

func (t *FileStorage) Delete(ctx context.Context, pid string, c *claims.MyCustomClaims) error {
	id, _ := strconv.Atoi(pid)
	existingTimer, err := t.Db.Project.
		Query().
		Where(project.And(project.ID(id), project.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return err
	}
	t.Db.Project.DeleteOne(existingTimer).Exec(ctx)

	return nil
}

// GetUser will return a user with a given Id
func (t *FileStorage) Get(ctx context.Context, pid string, c *claims.MyCustomClaims) (*pb.Project, error) {
	id, _ := strconv.Atoi(pid)
	newtimer, err := t.Db.Project.
		Query().
		Where(project.And(project.ID(id), project.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	response := toPb(newtimer)
	return response, nil
}

// Start will return a user with a given Id

func (t *FileStorage) GetAll(ctx context.Context, c *claims.MyCustomClaims) ([]*pb.Project, error) {
	existingTimer, err := t.Db.Project.
		Query().Where(project.Userid(c.Subject)).All(ctx)

	if err != nil {
		return nil, err
	}
	var list []*pb.Project
	for _, res := range existingTimer {
		t := toPb(res)
		list = append(list, t)
	}
	return list, err
}

func toPb(newtimer *ent.Project) *pb.Project {
	response := &pb.Project{Id: strconv.Itoa(newtimer.ID),
		Description: newtimer.Description,
		Client:      newtimer.Clientid,
		Tags:        newtimer.Tags,
		Name:        newtimer.Name,
		Status:      newtimer.Status,
		Team:        newtimer.Team,
	}
	return response
}
func checkNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
