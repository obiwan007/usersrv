// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"

	"github.com/obiwan007/usersrv/usersrv/api/storage/ent/user"

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
	TypeUser = "User"
)

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	userid        *string
	name          *string
	mandantid     *string
	roles         *string
	clearedFields map[string]struct{}
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op) *UserMutation {
	return &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUserid sets the userid field.
func (m *UserMutation) SetUserid(s string) {
	m.userid = &s
}

// Userid returns the userid value in the mutation.
func (m *UserMutation) Userid() (r string, exists bool) {
	v := m.userid
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserid reset all changes of the userid field.
func (m *UserMutation) ResetUserid() {
	m.userid = nil
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ClearName clears the value of name.
func (m *UserMutation) ClearName() {
	m.name = nil
	m.clearedFields[user.FieldName] = struct{}{}
}

// NameCleared returns if the field name was cleared in this mutation.
func (m *UserMutation) NameCleared() bool {
	_, ok := m.clearedFields[user.FieldName]
	return ok
}

// ResetName reset all changes of the name field.
func (m *UserMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, user.FieldName)
}

// SetMandantid sets the mandantid field.
func (m *UserMutation) SetMandantid(s string) {
	m.mandantid = &s
}

// Mandantid returns the mandantid value in the mutation.
func (m *UserMutation) Mandantid() (r string, exists bool) {
	v := m.mandantid
	if v == nil {
		return
	}
	return *v, true
}

// ClearMandantid clears the value of mandantid.
func (m *UserMutation) ClearMandantid() {
	m.mandantid = nil
	m.clearedFields[user.FieldMandantid] = struct{}{}
}

// MandantidCleared returns if the field mandantid was cleared in this mutation.
func (m *UserMutation) MandantidCleared() bool {
	_, ok := m.clearedFields[user.FieldMandantid]
	return ok
}

// ResetMandantid reset all changes of the mandantid field.
func (m *UserMutation) ResetMandantid() {
	m.mandantid = nil
	delete(m.clearedFields, user.FieldMandantid)
}

// SetRoles sets the roles field.
func (m *UserMutation) SetRoles(s string) {
	m.roles = &s
}

// Roles returns the roles value in the mutation.
func (m *UserMutation) Roles() (r string, exists bool) {
	v := m.roles
	if v == nil {
		return
	}
	return *v, true
}

// ClearRoles clears the value of roles.
func (m *UserMutation) ClearRoles() {
	m.roles = nil
	m.clearedFields[user.FieldRoles] = struct{}{}
}

// RolesCleared returns if the field roles was cleared in this mutation.
func (m *UserMutation) RolesCleared() bool {
	_, ok := m.clearedFields[user.FieldRoles]
	return ok
}

// ResetRoles reset all changes of the roles field.
func (m *UserMutation) ResetRoles() {
	m.roles = nil
	delete(m.clearedFields, user.FieldRoles)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.userid != nil {
		fields = append(fields, user.FieldUserid)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.mandantid != nil {
		fields = append(fields, user.FieldMandantid)
	}
	if m.roles != nil {
		fields = append(fields, user.FieldRoles)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUserid:
		return m.Userid()
	case user.FieldName:
		return m.Name()
	case user.FieldMandantid:
		return m.Mandantid()
	case user.FieldRoles:
		return m.Roles()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUserid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserid(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldMandantid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMandantid(v)
		return nil
	case user.FieldRoles:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRoles(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldName) {
		fields = append(fields, user.FieldName)
	}
	if m.FieldCleared(user.FieldMandantid) {
		fields = append(fields, user.FieldMandantid)
	}
	if m.FieldCleared(user.FieldRoles) {
		fields = append(fields, user.FieldRoles)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldName:
		m.ClearName()
		return nil
	case user.FieldMandantid:
		m.ClearMandantid()
		return nil
	case user.FieldRoles:
		m.ClearRoles()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUserid:
		m.ResetUserid()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldMandantid:
		m.ResetMandantid()
		return nil
	case user.FieldRoles:
		m.ResetRoles()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User edge %s", name)
}
