// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"

	"github.com/obiwan007/usersrv/clientsrv/api/storage/ent/timerclient"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTimerClient = "TimerClient"
)

// TimerClientMutation represents an operation that mutate the TimerClients
// nodes in the graph.
type TimerClientMutation struct {
	config
	op            Op
	typ           string
	id            *int
	description   *string
	name          *string
	address       *string
	userid        *string
	mandantid     *string
	clearedFields map[string]struct{}
}

var _ ent.Mutation = (*TimerClientMutation)(nil)

// newTimerClientMutation creates new mutation for $n.Name.
func newTimerClientMutation(c config, op Op) *TimerClientMutation {
	return &TimerClientMutation{
		config:        c,
		op:            op,
		typ:           TypeTimerClient,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TimerClientMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TimerClientMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *TimerClientMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetDescription sets the description field.
func (m *TimerClientMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the description value in the mutation.
func (m *TimerClientMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// ResetDescription reset all changes of the description field.
func (m *TimerClientMutation) ResetDescription() {
	m.description = nil
}

// SetName sets the name field.
func (m *TimerClientMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *TimerClientMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *TimerClientMutation) ResetName() {
	m.name = nil
}

// SetAddress sets the address field.
func (m *TimerClientMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the address value in the mutation.
func (m *TimerClientMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// ClearAddress clears the value of address.
func (m *TimerClientMutation) ClearAddress() {
	m.address = nil
	m.clearedFields[timerclient.FieldAddress] = struct{}{}
}

// AddressCleared returns if the field address was cleared in this mutation.
func (m *TimerClientMutation) AddressCleared() bool {
	_, ok := m.clearedFields[timerclient.FieldAddress]
	return ok
}

// ResetAddress reset all changes of the address field.
func (m *TimerClientMutation) ResetAddress() {
	m.address = nil
	delete(m.clearedFields, timerclient.FieldAddress)
}

// SetUserid sets the userid field.
func (m *TimerClientMutation) SetUserid(s string) {
	m.userid = &s
}

// Userid returns the userid value in the mutation.
func (m *TimerClientMutation) Userid() (r string, exists bool) {
	v := m.userid
	if v == nil {
		return
	}
	return *v, true
}

// ClearUserid clears the value of userid.
func (m *TimerClientMutation) ClearUserid() {
	m.userid = nil
	m.clearedFields[timerclient.FieldUserid] = struct{}{}
}

// UseridCleared returns if the field userid was cleared in this mutation.
func (m *TimerClientMutation) UseridCleared() bool {
	_, ok := m.clearedFields[timerclient.FieldUserid]
	return ok
}

// ResetUserid reset all changes of the userid field.
func (m *TimerClientMutation) ResetUserid() {
	m.userid = nil
	delete(m.clearedFields, timerclient.FieldUserid)
}

// SetMandantid sets the mandantid field.
func (m *TimerClientMutation) SetMandantid(s string) {
	m.mandantid = &s
}

// Mandantid returns the mandantid value in the mutation.
func (m *TimerClientMutation) Mandantid() (r string, exists bool) {
	v := m.mandantid
	if v == nil {
		return
	}
	return *v, true
}

// ClearMandantid clears the value of mandantid.
func (m *TimerClientMutation) ClearMandantid() {
	m.mandantid = nil
	m.clearedFields[timerclient.FieldMandantid] = struct{}{}
}

// MandantidCleared returns if the field mandantid was cleared in this mutation.
func (m *TimerClientMutation) MandantidCleared() bool {
	_, ok := m.clearedFields[timerclient.FieldMandantid]
	return ok
}

// ResetMandantid reset all changes of the mandantid field.
func (m *TimerClientMutation) ResetMandantid() {
	m.mandantid = nil
	delete(m.clearedFields, timerclient.FieldMandantid)
}

// Op returns the operation name.
func (m *TimerClientMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (TimerClient).
func (m *TimerClientMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *TimerClientMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.description != nil {
		fields = append(fields, timerclient.FieldDescription)
	}
	if m.name != nil {
		fields = append(fields, timerclient.FieldName)
	}
	if m.address != nil {
		fields = append(fields, timerclient.FieldAddress)
	}
	if m.userid != nil {
		fields = append(fields, timerclient.FieldUserid)
	}
	if m.mandantid != nil {
		fields = append(fields, timerclient.FieldMandantid)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *TimerClientMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case timerclient.FieldDescription:
		return m.Description()
	case timerclient.FieldName:
		return m.Name()
	case timerclient.FieldAddress:
		return m.Address()
	case timerclient.FieldUserid:
		return m.Userid()
	case timerclient.FieldMandantid:
		return m.Mandantid()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TimerClientMutation) SetField(name string, value ent.Value) error {
	switch name {
	case timerclient.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case timerclient.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case timerclient.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case timerclient.FieldUserid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserid(v)
		return nil
	case timerclient.FieldMandantid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMandantid(v)
		return nil
	}
	return fmt.Errorf("unknown TimerClient field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *TimerClientMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *TimerClientMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TimerClientMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown TimerClient numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *TimerClientMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(timerclient.FieldAddress) {
		fields = append(fields, timerclient.FieldAddress)
	}
	if m.FieldCleared(timerclient.FieldUserid) {
		fields = append(fields, timerclient.FieldUserid)
	}
	if m.FieldCleared(timerclient.FieldMandantid) {
		fields = append(fields, timerclient.FieldMandantid)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *TimerClientMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *TimerClientMutation) ClearField(name string) error {
	switch name {
	case timerclient.FieldAddress:
		m.ClearAddress()
		return nil
	case timerclient.FieldUserid:
		m.ClearUserid()
		return nil
	case timerclient.FieldMandantid:
		m.ClearMandantid()
		return nil
	}
	return fmt.Errorf("unknown TimerClient nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *TimerClientMutation) ResetField(name string) error {
	switch name {
	case timerclient.FieldDescription:
		m.ResetDescription()
		return nil
	case timerclient.FieldName:
		m.ResetName()
		return nil
	case timerclient.FieldAddress:
		m.ResetAddress()
		return nil
	case timerclient.FieldUserid:
		m.ResetUserid()
		return nil
	case timerclient.FieldMandantid:
		m.ResetMandantid()
		return nil
	}
	return fmt.Errorf("unknown TimerClient field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *TimerClientMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *TimerClientMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *TimerClientMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *TimerClientMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *TimerClientMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *TimerClientMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *TimerClientMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown TimerClient unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *TimerClientMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown TimerClient edge %s", name)
}
