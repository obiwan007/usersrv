package userservicestorage

import (
	"context"
	"log"
	"strconv"
	"sync"

	_ "github.com/lib/pq"
	claims "github.com/obiwan007/usersrv/pkg/claims"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/usersrv/api/storage/ent"
	"github.com/obiwan007/usersrv/usersrv/api/storage/ent/user"
)

type FileStorage struct{ Db *ent.Client }

var users []User
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

func (t *FileStorage) AddUser(ctx context.Context, entity *pb.User, c *claims.MyCustomClaims) (*pb.User, error) {

	existing, err := t.Db.User.
		Query().
		Where(user.Userid(c.Subject)).
		Only(ctx)
	if err == nil {
		new, err := existing.
			Update().
			SetNillableName(checkNil(entity.Name)).
			SetMandantid(c.Mandant).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		response := toPb(new)
		log.Println("Updated existing user")
		return response, nil
	}

	newEntity, err := t.Db.User. // UserClient.
					Create().             // User create builder.
					SetName(entity.Name). // Set field value.
					SetMandantid(c.Mandant).
					SetUserid(c.Subject).
					Save(ctx) // Create and return.
	if err != nil {
		return nil, err
	}
	response := &pb.User{Id: strconv.Itoa(newEntity.ID),
		Name:  newEntity.Name,
		Email: newEntity.Userid,
	}

	return response, err
}

func (t *FileStorage) DeleteUser(ctx context.Context, entityID string, c *claims.MyCustomClaims) error {
	id, _ := strconv.Atoi(entityID)
	existingTimer, err := t.Db.User.
		Query().
		Where(user.And(user.ID(id), user.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return err
	}
	t.Db.User.DeleteOne(existingTimer).Exec(ctx)

	return err
}

// GetUser will return a user with a given Id
func (t *FileStorage) GetUser(ctx context.Context, ID string, c *claims.MyCustomClaims) (*pb.User, error) {
	id, _ := strconv.Atoi(ID)
	newtimer, err := t.Db.User.
		Query().
		Where(user.And(user.ID(id), user.Userid(c.Subject))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	response := toPb(newtimer)
	return response, nil
}

// GetUserFromEmail will return a user with a given email
func (t *FileStorage) GetUserFromEmail(ctx context.Context, mail string, c *claims.MyCustomClaims) (*pb.User, error) {
	newtimer, err := t.Db.User.
		Query().
		Where(user.Userid(c.Subject)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	response := toPb(newtimer)
	return response, nil
}

func (t *FileStorage) ListUser(ctx context.Context, c *claims.MyCustomClaims) ([]*pb.User, error) {
	existingTimer, err := t.Db.User.
		Query().Where(user.Userid(c.Subject)).All(ctx)

	if err != nil {
		return nil, err
	}
	var list []*pb.User
	for _, res := range existingTimer {
		t := toPb(res)
		list = append(list, t)
	}
	return list, err
}

func toPb(newtimer *ent.User) *pb.User {
	response := &pb.User{Id: strconv.Itoa(newtimer.ID),
		Email: newtimer.Userid,
		Name:  newtimer.Name,
	}
	return response
}

func checkNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
