package timerservicestorage

import (
	"context"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/obiwan007/usersrv/clientsrv/api/storage/ent"
	"github.com/obiwan007/usersrv/clientsrv/api/storage/ent/timerclient"

	claims "github.com/obiwan007/usersrv/pkg/claims"
	pb "github.com/obiwan007/usersrv/proto"
)

// FileStorage type
type FileStorage struct{ Db *ent.Client }

// NewFileStorage delvers an instance of the DB
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

// Add new Client to DB
func (t *FileStorage) Add(ctx context.Context, entity *pb.Client, c *claims.MyCustomClaims) (*pb.Client, error) {

	newEntity, err := t.Db.TimerClient. // UserClient.
						Create().                           // User create builder.
						SetName(entity.Name).               // Set field value.
						SetDescription(entity.Description). // Set field value.
						SetAddress(entity.Address).         // Set field value.
						SetMandantid(c.Mandant).
						SetUserid(c.Subject).
						Save(ctx) // Create and return.
	if err != nil {
		return nil, err
	}
	response := &pb.Client{Id: strconv.Itoa(newEntity.ID),
		Description: newEntity.Description,
		Name:        newEntity.Name,
		Address:     newEntity.Address,
	}

	return response, err
}

// Update Client via new Client object
func (t *FileStorage) Update(ctx context.Context, entity *pb.Client, c *claims.MyCustomClaims) (*pb.Client, error) {
	id, _ := strconv.Atoi(entity.Id)

	existingTimer, err := t.Db.TimerClient.
		Query().
		Where(timerclient.And(timerclient.ID(id), timerclient.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	new, err := existingTimer.
		Update().
		SetNillableDescription(checkNil(entity.Description)).
		SetNillableName(checkNil(entity.Name)).
		SetNillableAddress(checkNil(entity.Address)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	response := toPb(new)
	return response, err
}

// Delete Client from DB
func (t *FileStorage) Delete(ctx context.Context, entityID string, c *claims.MyCustomClaims) (*pb.Client, error) {
	id, _ := strconv.Atoi(entityID)
	existingTimer, err := t.Db.TimerClient.
		Query().
		Where(timerclient.And(timerclient.ID(id), timerclient.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	t.Db.TimerClient.DeleteOne(existingTimer).Exec(ctx)

	return nil, nil
}

// Get Client will return a user with a given Id
func (t *FileStorage) Get(ctx context.Context, ID string, c *claims.MyCustomClaims) (*pb.Client, error) {
	id, _ := strconv.Atoi(ID)
	newtimer, err := t.Db.TimerClient.
		Query().
		Where(timerclient.And(timerclient.ID(id), timerclient.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	response := toPb(newtimer)
	return response, nil
}

// GetAll returns list of your clients
func (t *FileStorage) GetAll(ctx context.Context, c *claims.MyCustomClaims) ([]*pb.Client, error) {
	existingTimer, err := t.Db.TimerClient.
		Query().Where(timerclient.Userid(c.Subject)).All(ctx)

	if err != nil {
		return nil, err
	}
	var list []*pb.Client
	for _, res := range existingTimer {
		t := toPb(res)
		list = append(list, t)
	}
	return list, err
}

func toPb(newtimer *ent.TimerClient) *pb.Client {
	response := &pb.Client{Id: strconv.Itoa(newtimer.ID),
		Description: newtimer.Description,
		Address:     newtimer.Address,
		Name:        newtimer.Name,
	}
	return response
}

func checkNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
