// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/obiwan007/usersrv/clientsrv/api/storage/ent/timerclient"
)

// TimerClient is the model entity for the TimerClient schema.
type TimerClient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Userid holds the value of the "userid" field.
	Userid string `json:"userid,omitempty"`
	// Mandantid holds the value of the "mandantid" field.
	Mandantid string `json:"mandantid,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TimerClient) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // description
		&sql.NullString{}, // name
		&sql.NullString{}, // address
		&sql.NullString{}, // userid
		&sql.NullString{}, // mandantid
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TimerClient fields.
func (tc *TimerClient) assignValues(values ...interface{}) error {
	if m, n := len(values), len(timerclient.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	tc.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[0])
	} else if value.Valid {
		tc.Description = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		tc.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field address", values[2])
	} else if value.Valid {
		tc.Address = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field userid", values[3])
	} else if value.Valid {
		tc.Userid = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field mandantid", values[4])
	} else if value.Valid {
		tc.Mandantid = value.String
	}
	return nil
}

// Update returns a builder for updating this TimerClient.
// Note that, you need to call TimerClient.Unwrap() before calling this method, if this TimerClient
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TimerClient) Update() *TimerClientUpdateOne {
	return (&TimerClientClient{config: tc.config}).UpdateOne(tc)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (tc *TimerClient) Unwrap() *TimerClient {
	tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TimerClient is not a transactional entity")
	}
	tc.config.driver = tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TimerClient) String() string {
	var builder strings.Builder
	builder.WriteString("TimerClient(")
	builder.WriteString(fmt.Sprintf("id=%v", tc.ID))
	builder.WriteString(", description=")
	builder.WriteString(tc.Description)
	builder.WriteString(", name=")
	builder.WriteString(tc.Name)
	builder.WriteString(", address=")
	builder.WriteString(tc.Address)
	builder.WriteString(", userid=")
	builder.WriteString(tc.Userid)
	builder.WriteString(", mandantid=")
	builder.WriteString(tc.Mandantid)
	builder.WriteByte(')')
	return builder.String()
}

// TimerClients is a parsable slice of TimerClient.
type TimerClients []*TimerClient

func (tc TimerClients) config(cfg config) {
	for _i := range tc {
		tc[_i].config = cfg
	}
}