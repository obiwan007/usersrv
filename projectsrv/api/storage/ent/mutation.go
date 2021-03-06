// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"

	"github.com/obiwan007/usersrv/projectsrv/api/storage/ent/project"

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
	TypeProject = "Project"
)

// ProjectMutation represents an operation that mutate the Projects
// nodes in the graph.
type ProjectMutation struct {
	config
	op            Op
	typ           string
	id            *int
	description   *string
	name          *string
	clientid      *string
	userid        *string
	mandantid     *string
	status        *string
	tags          *string
	team          *string
	clearedFields map[string]struct{}
}

var _ ent.Mutation = (*ProjectMutation)(nil)

// newProjectMutation creates new mutation for $n.Name.
func newProjectMutation(c config, op Op) *ProjectMutation {
	return &ProjectMutation{
		config:        c,
		op:            op,
		typ:           TypeProject,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProjectMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProjectMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *ProjectMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetDescription sets the description field.
func (m *ProjectMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the description value in the mutation.
func (m *ProjectMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// ResetDescription reset all changes of the description field.
func (m *ProjectMutation) ResetDescription() {
	m.description = nil
}

// SetName sets the name field.
func (m *ProjectMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *ProjectMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *ProjectMutation) ResetName() {
	m.name = nil
}

// SetClientid sets the clientid field.
func (m *ProjectMutation) SetClientid(s string) {
	m.clientid = &s
}

// Clientid returns the clientid value in the mutation.
func (m *ProjectMutation) Clientid() (r string, exists bool) {
	v := m.clientid
	if v == nil {
		return
	}
	return *v, true
}

// ClearClientid clears the value of clientid.
func (m *ProjectMutation) ClearClientid() {
	m.clientid = nil
	m.clearedFields[project.FieldClientid] = struct{}{}
}

// ClientidCleared returns if the field clientid was cleared in this mutation.
func (m *ProjectMutation) ClientidCleared() bool {
	_, ok := m.clearedFields[project.FieldClientid]
	return ok
}

// ResetClientid reset all changes of the clientid field.
func (m *ProjectMutation) ResetClientid() {
	m.clientid = nil
	delete(m.clearedFields, project.FieldClientid)
}

// SetUserid sets the userid field.
func (m *ProjectMutation) SetUserid(s string) {
	m.userid = &s
}

// Userid returns the userid value in the mutation.
func (m *ProjectMutation) Userid() (r string, exists bool) {
	v := m.userid
	if v == nil {
		return
	}
	return *v, true
}

// ClearUserid clears the value of userid.
func (m *ProjectMutation) ClearUserid() {
	m.userid = nil
	m.clearedFields[project.FieldUserid] = struct{}{}
}

// UseridCleared returns if the field userid was cleared in this mutation.
func (m *ProjectMutation) UseridCleared() bool {
	_, ok := m.clearedFields[project.FieldUserid]
	return ok
}

// ResetUserid reset all changes of the userid field.
func (m *ProjectMutation) ResetUserid() {
	m.userid = nil
	delete(m.clearedFields, project.FieldUserid)
}

// SetMandantid sets the mandantid field.
func (m *ProjectMutation) SetMandantid(s string) {
	m.mandantid = &s
}

// Mandantid returns the mandantid value in the mutation.
func (m *ProjectMutation) Mandantid() (r string, exists bool) {
	v := m.mandantid
	if v == nil {
		return
	}
	return *v, true
}

// ClearMandantid clears the value of mandantid.
func (m *ProjectMutation) ClearMandantid() {
	m.mandantid = nil
	m.clearedFields[project.FieldMandantid] = struct{}{}
}

// MandantidCleared returns if the field mandantid was cleared in this mutation.
func (m *ProjectMutation) MandantidCleared() bool {
	_, ok := m.clearedFields[project.FieldMandantid]
	return ok
}

// ResetMandantid reset all changes of the mandantid field.
func (m *ProjectMutation) ResetMandantid() {
	m.mandantid = nil
	delete(m.clearedFields, project.FieldMandantid)
}

// SetStatus sets the status field.
func (m *ProjectMutation) SetStatus(s string) {
	m.status = &s
}

// Status returns the status value in the mutation.
func (m *ProjectMutation) Status() (r string, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// ClearStatus clears the value of status.
func (m *ProjectMutation) ClearStatus() {
	m.status = nil
	m.clearedFields[project.FieldStatus] = struct{}{}
}

// StatusCleared returns if the field status was cleared in this mutation.
func (m *ProjectMutation) StatusCleared() bool {
	_, ok := m.clearedFields[project.FieldStatus]
	return ok
}

// ResetStatus reset all changes of the status field.
func (m *ProjectMutation) ResetStatus() {
	m.status = nil
	delete(m.clearedFields, project.FieldStatus)
}

// SetTags sets the tags field.
func (m *ProjectMutation) SetTags(s string) {
	m.tags = &s
}

// Tags returns the tags value in the mutation.
func (m *ProjectMutation) Tags() (r string, exists bool) {
	v := m.tags
	if v == nil {
		return
	}
	return *v, true
}

// ClearTags clears the value of tags.
func (m *ProjectMutation) ClearTags() {
	m.tags = nil
	m.clearedFields[project.FieldTags] = struct{}{}
}

// TagsCleared returns if the field tags was cleared in this mutation.
func (m *ProjectMutation) TagsCleared() bool {
	_, ok := m.clearedFields[project.FieldTags]
	return ok
}

// ResetTags reset all changes of the tags field.
func (m *ProjectMutation) ResetTags() {
	m.tags = nil
	delete(m.clearedFields, project.FieldTags)
}

// SetTeam sets the team field.
func (m *ProjectMutation) SetTeam(s string) {
	m.team = &s
}

// Team returns the team value in the mutation.
func (m *ProjectMutation) Team() (r string, exists bool) {
	v := m.team
	if v == nil {
		return
	}
	return *v, true
}

// ClearTeam clears the value of team.
func (m *ProjectMutation) ClearTeam() {
	m.team = nil
	m.clearedFields[project.FieldTeam] = struct{}{}
}

// TeamCleared returns if the field team was cleared in this mutation.
func (m *ProjectMutation) TeamCleared() bool {
	_, ok := m.clearedFields[project.FieldTeam]
	return ok
}

// ResetTeam reset all changes of the team field.
func (m *ProjectMutation) ResetTeam() {
	m.team = nil
	delete(m.clearedFields, project.FieldTeam)
}

// Op returns the operation name.
func (m *ProjectMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Project).
func (m *ProjectMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *ProjectMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.description != nil {
		fields = append(fields, project.FieldDescription)
	}
	if m.name != nil {
		fields = append(fields, project.FieldName)
	}
	if m.clientid != nil {
		fields = append(fields, project.FieldClientid)
	}
	if m.userid != nil {
		fields = append(fields, project.FieldUserid)
	}
	if m.mandantid != nil {
		fields = append(fields, project.FieldMandantid)
	}
	if m.status != nil {
		fields = append(fields, project.FieldStatus)
	}
	if m.tags != nil {
		fields = append(fields, project.FieldTags)
	}
	if m.team != nil {
		fields = append(fields, project.FieldTeam)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *ProjectMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case project.FieldDescription:
		return m.Description()
	case project.FieldName:
		return m.Name()
	case project.FieldClientid:
		return m.Clientid()
	case project.FieldUserid:
		return m.Userid()
	case project.FieldMandantid:
		return m.Mandantid()
	case project.FieldStatus:
		return m.Status()
	case project.FieldTags:
		return m.Tags()
	case project.FieldTeam:
		return m.Team()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *ProjectMutation) SetField(name string, value ent.Value) error {
	switch name {
	case project.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case project.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case project.FieldClientid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetClientid(v)
		return nil
	case project.FieldUserid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserid(v)
		return nil
	case project.FieldMandantid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMandantid(v)
		return nil
	case project.FieldStatus:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case project.FieldTags:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTags(v)
		return nil
	case project.FieldTeam:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTeam(v)
		return nil
	}
	return fmt.Errorf("unknown Project field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *ProjectMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *ProjectMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *ProjectMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Project numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *ProjectMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(project.FieldClientid) {
		fields = append(fields, project.FieldClientid)
	}
	if m.FieldCleared(project.FieldUserid) {
		fields = append(fields, project.FieldUserid)
	}
	if m.FieldCleared(project.FieldMandantid) {
		fields = append(fields, project.FieldMandantid)
	}
	if m.FieldCleared(project.FieldStatus) {
		fields = append(fields, project.FieldStatus)
	}
	if m.FieldCleared(project.FieldTags) {
		fields = append(fields, project.FieldTags)
	}
	if m.FieldCleared(project.FieldTeam) {
		fields = append(fields, project.FieldTeam)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *ProjectMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProjectMutation) ClearField(name string) error {
	switch name {
	case project.FieldClientid:
		m.ClearClientid()
		return nil
	case project.FieldUserid:
		m.ClearUserid()
		return nil
	case project.FieldMandantid:
		m.ClearMandantid()
		return nil
	case project.FieldStatus:
		m.ClearStatus()
		return nil
	case project.FieldTags:
		m.ClearTags()
		return nil
	case project.FieldTeam:
		m.ClearTeam()
		return nil
	}
	return fmt.Errorf("unknown Project nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *ProjectMutation) ResetField(name string) error {
	switch name {
	case project.FieldDescription:
		m.ResetDescription()
		return nil
	case project.FieldName:
		m.ResetName()
		return nil
	case project.FieldClientid:
		m.ResetClientid()
		return nil
	case project.FieldUserid:
		m.ResetUserid()
		return nil
	case project.FieldMandantid:
		m.ResetMandantid()
		return nil
	case project.FieldStatus:
		m.ResetStatus()
		return nil
	case project.FieldTags:
		m.ResetTags()
		return nil
	case project.FieldTeam:
		m.ResetTeam()
		return nil
	}
	return fmt.Errorf("unknown Project field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *ProjectMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *ProjectMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *ProjectMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *ProjectMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *ProjectMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *ProjectMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *ProjectMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Project unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *ProjectMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Project edge %s", name)
}
