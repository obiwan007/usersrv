// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/obiwan007/usersrv/timersrv/api/storage/ent/timer"

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
	TypeTimer = "Timer"
)

// TimerMutation represents an operation that mutate the Timers
// nodes in the graph.
type TimerMutation struct {
	config
	op                Op
	typ               string
	id                *int
	description       *string
	projectid         *string
	tags              *string
	timerStart        *time.Time
	timerEnd          *time.Time
	elapsedSeconds    *int
	addelapsedSeconds *int
	_IsBilled         *bool
	_IsRunning        *bool
	userid            *string
	mandantid         *string
	clearedFields     map[string]struct{}
}

var _ ent.Mutation = (*TimerMutation)(nil)

// newTimerMutation creates new mutation for $n.Name.
func newTimerMutation(c config, op Op) *TimerMutation {
	return &TimerMutation{
		config:        c,
		op:            op,
		typ:           TypeTimer,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TimerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TimerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *TimerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetDescription sets the description field.
func (m *TimerMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the description value in the mutation.
func (m *TimerMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// ResetDescription reset all changes of the description field.
func (m *TimerMutation) ResetDescription() {
	m.description = nil
}

// SetProjectid sets the projectid field.
func (m *TimerMutation) SetProjectid(s string) {
	m.projectid = &s
}

// Projectid returns the projectid value in the mutation.
func (m *TimerMutation) Projectid() (r string, exists bool) {
	v := m.projectid
	if v == nil {
		return
	}
	return *v, true
}

// ClearProjectid clears the value of projectid.
func (m *TimerMutation) ClearProjectid() {
	m.projectid = nil
	m.clearedFields[timer.FieldProjectid] = struct{}{}
}

// ProjectidCleared returns if the field projectid was cleared in this mutation.
func (m *TimerMutation) ProjectidCleared() bool {
	_, ok := m.clearedFields[timer.FieldProjectid]
	return ok
}

// ResetProjectid reset all changes of the projectid field.
func (m *TimerMutation) ResetProjectid() {
	m.projectid = nil
	delete(m.clearedFields, timer.FieldProjectid)
}

// SetTags sets the tags field.
func (m *TimerMutation) SetTags(s string) {
	m.tags = &s
}

// Tags returns the tags value in the mutation.
func (m *TimerMutation) Tags() (r string, exists bool) {
	v := m.tags
	if v == nil {
		return
	}
	return *v, true
}

// ClearTags clears the value of tags.
func (m *TimerMutation) ClearTags() {
	m.tags = nil
	m.clearedFields[timer.FieldTags] = struct{}{}
}

// TagsCleared returns if the field tags was cleared in this mutation.
func (m *TimerMutation) TagsCleared() bool {
	_, ok := m.clearedFields[timer.FieldTags]
	return ok
}

// ResetTags reset all changes of the tags field.
func (m *TimerMutation) ResetTags() {
	m.tags = nil
	delete(m.clearedFields, timer.FieldTags)
}

// SetTimerStart sets the timerStart field.
func (m *TimerMutation) SetTimerStart(t time.Time) {
	m.timerStart = &t
}

// TimerStart returns the timerStart value in the mutation.
func (m *TimerMutation) TimerStart() (r time.Time, exists bool) {
	v := m.timerStart
	if v == nil {
		return
	}
	return *v, true
}

// ClearTimerStart clears the value of timerStart.
func (m *TimerMutation) ClearTimerStart() {
	m.timerStart = nil
	m.clearedFields[timer.FieldTimerStart] = struct{}{}
}

// TimerStartCleared returns if the field timerStart was cleared in this mutation.
func (m *TimerMutation) TimerStartCleared() bool {
	_, ok := m.clearedFields[timer.FieldTimerStart]
	return ok
}

// ResetTimerStart reset all changes of the timerStart field.
func (m *TimerMutation) ResetTimerStart() {
	m.timerStart = nil
	delete(m.clearedFields, timer.FieldTimerStart)
}

// SetTimerEnd sets the timerEnd field.
func (m *TimerMutation) SetTimerEnd(t time.Time) {
	m.timerEnd = &t
}

// TimerEnd returns the timerEnd value in the mutation.
func (m *TimerMutation) TimerEnd() (r time.Time, exists bool) {
	v := m.timerEnd
	if v == nil {
		return
	}
	return *v, true
}

// ClearTimerEnd clears the value of timerEnd.
func (m *TimerMutation) ClearTimerEnd() {
	m.timerEnd = nil
	m.clearedFields[timer.FieldTimerEnd] = struct{}{}
}

// TimerEndCleared returns if the field timerEnd was cleared in this mutation.
func (m *TimerMutation) TimerEndCleared() bool {
	_, ok := m.clearedFields[timer.FieldTimerEnd]
	return ok
}

// ResetTimerEnd reset all changes of the timerEnd field.
func (m *TimerMutation) ResetTimerEnd() {
	m.timerEnd = nil
	delete(m.clearedFields, timer.FieldTimerEnd)
}

// SetElapsedSeconds sets the elapsedSeconds field.
func (m *TimerMutation) SetElapsedSeconds(i int) {
	m.elapsedSeconds = &i
	m.addelapsedSeconds = nil
}

// ElapsedSeconds returns the elapsedSeconds value in the mutation.
func (m *TimerMutation) ElapsedSeconds() (r int, exists bool) {
	v := m.elapsedSeconds
	if v == nil {
		return
	}
	return *v, true
}

// AddElapsedSeconds adds i to elapsedSeconds.
func (m *TimerMutation) AddElapsedSeconds(i int) {
	if m.addelapsedSeconds != nil {
		*m.addelapsedSeconds += i
	} else {
		m.addelapsedSeconds = &i
	}
}

// AddedElapsedSeconds returns the value that was added to the elapsedSeconds field in this mutation.
func (m *TimerMutation) AddedElapsedSeconds() (r int, exists bool) {
	v := m.addelapsedSeconds
	if v == nil {
		return
	}
	return *v, true
}

// ClearElapsedSeconds clears the value of elapsedSeconds.
func (m *TimerMutation) ClearElapsedSeconds() {
	m.elapsedSeconds = nil
	m.addelapsedSeconds = nil
	m.clearedFields[timer.FieldElapsedSeconds] = struct{}{}
}

// ElapsedSecondsCleared returns if the field elapsedSeconds was cleared in this mutation.
func (m *TimerMutation) ElapsedSecondsCleared() bool {
	_, ok := m.clearedFields[timer.FieldElapsedSeconds]
	return ok
}

// ResetElapsedSeconds reset all changes of the elapsedSeconds field.
func (m *TimerMutation) ResetElapsedSeconds() {
	m.elapsedSeconds = nil
	m.addelapsedSeconds = nil
	delete(m.clearedFields, timer.FieldElapsedSeconds)
}

// SetIsBilled sets the IsBilled field.
func (m *TimerMutation) SetIsBilled(b bool) {
	m._IsBilled = &b
}

// IsBilled returns the IsBilled value in the mutation.
func (m *TimerMutation) IsBilled() (r bool, exists bool) {
	v := m._IsBilled
	if v == nil {
		return
	}
	return *v, true
}

// ClearIsBilled clears the value of IsBilled.
func (m *TimerMutation) ClearIsBilled() {
	m._IsBilled = nil
	m.clearedFields[timer.FieldIsBilled] = struct{}{}
}

// IsBilledCleared returns if the field IsBilled was cleared in this mutation.
func (m *TimerMutation) IsBilledCleared() bool {
	_, ok := m.clearedFields[timer.FieldIsBilled]
	return ok
}

// ResetIsBilled reset all changes of the IsBilled field.
func (m *TimerMutation) ResetIsBilled() {
	m._IsBilled = nil
	delete(m.clearedFields, timer.FieldIsBilled)
}

// SetIsRunning sets the IsRunning field.
func (m *TimerMutation) SetIsRunning(b bool) {
	m._IsRunning = &b
}

// IsRunning returns the IsRunning value in the mutation.
func (m *TimerMutation) IsRunning() (r bool, exists bool) {
	v := m._IsRunning
	if v == nil {
		return
	}
	return *v, true
}

// ClearIsRunning clears the value of IsRunning.
func (m *TimerMutation) ClearIsRunning() {
	m._IsRunning = nil
	m.clearedFields[timer.FieldIsRunning] = struct{}{}
}

// IsRunningCleared returns if the field IsRunning was cleared in this mutation.
func (m *TimerMutation) IsRunningCleared() bool {
	_, ok := m.clearedFields[timer.FieldIsRunning]
	return ok
}

// ResetIsRunning reset all changes of the IsRunning field.
func (m *TimerMutation) ResetIsRunning() {
	m._IsRunning = nil
	delete(m.clearedFields, timer.FieldIsRunning)
}

// SetUserid sets the userid field.
func (m *TimerMutation) SetUserid(s string) {
	m.userid = &s
}

// Userid returns the userid value in the mutation.
func (m *TimerMutation) Userid() (r string, exists bool) {
	v := m.userid
	if v == nil {
		return
	}
	return *v, true
}

// ClearUserid clears the value of userid.
func (m *TimerMutation) ClearUserid() {
	m.userid = nil
	m.clearedFields[timer.FieldUserid] = struct{}{}
}

// UseridCleared returns if the field userid was cleared in this mutation.
func (m *TimerMutation) UseridCleared() bool {
	_, ok := m.clearedFields[timer.FieldUserid]
	return ok
}

// ResetUserid reset all changes of the userid field.
func (m *TimerMutation) ResetUserid() {
	m.userid = nil
	delete(m.clearedFields, timer.FieldUserid)
}

// SetMandantid sets the mandantid field.
func (m *TimerMutation) SetMandantid(s string) {
	m.mandantid = &s
}

// Mandantid returns the mandantid value in the mutation.
func (m *TimerMutation) Mandantid() (r string, exists bool) {
	v := m.mandantid
	if v == nil {
		return
	}
	return *v, true
}

// ClearMandantid clears the value of mandantid.
func (m *TimerMutation) ClearMandantid() {
	m.mandantid = nil
	m.clearedFields[timer.FieldMandantid] = struct{}{}
}

// MandantidCleared returns if the field mandantid was cleared in this mutation.
func (m *TimerMutation) MandantidCleared() bool {
	_, ok := m.clearedFields[timer.FieldMandantid]
	return ok
}

// ResetMandantid reset all changes of the mandantid field.
func (m *TimerMutation) ResetMandantid() {
	m.mandantid = nil
	delete(m.clearedFields, timer.FieldMandantid)
}

// Op returns the operation name.
func (m *TimerMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Timer).
func (m *TimerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *TimerMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.description != nil {
		fields = append(fields, timer.FieldDescription)
	}
	if m.projectid != nil {
		fields = append(fields, timer.FieldProjectid)
	}
	if m.tags != nil {
		fields = append(fields, timer.FieldTags)
	}
	if m.timerStart != nil {
		fields = append(fields, timer.FieldTimerStart)
	}
	if m.timerEnd != nil {
		fields = append(fields, timer.FieldTimerEnd)
	}
	if m.elapsedSeconds != nil {
		fields = append(fields, timer.FieldElapsedSeconds)
	}
	if m._IsBilled != nil {
		fields = append(fields, timer.FieldIsBilled)
	}
	if m._IsRunning != nil {
		fields = append(fields, timer.FieldIsRunning)
	}
	if m.userid != nil {
		fields = append(fields, timer.FieldUserid)
	}
	if m.mandantid != nil {
		fields = append(fields, timer.FieldMandantid)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *TimerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case timer.FieldDescription:
		return m.Description()
	case timer.FieldProjectid:
		return m.Projectid()
	case timer.FieldTags:
		return m.Tags()
	case timer.FieldTimerStart:
		return m.TimerStart()
	case timer.FieldTimerEnd:
		return m.TimerEnd()
	case timer.FieldElapsedSeconds:
		return m.ElapsedSeconds()
	case timer.FieldIsBilled:
		return m.IsBilled()
	case timer.FieldIsRunning:
		return m.IsRunning()
	case timer.FieldUserid:
		return m.Userid()
	case timer.FieldMandantid:
		return m.Mandantid()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TimerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case timer.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case timer.FieldProjectid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProjectid(v)
		return nil
	case timer.FieldTags:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTags(v)
		return nil
	case timer.FieldTimerStart:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTimerStart(v)
		return nil
	case timer.FieldTimerEnd:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTimerEnd(v)
		return nil
	case timer.FieldElapsedSeconds:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetElapsedSeconds(v)
		return nil
	case timer.FieldIsBilled:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsBilled(v)
		return nil
	case timer.FieldIsRunning:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsRunning(v)
		return nil
	case timer.FieldUserid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserid(v)
		return nil
	case timer.FieldMandantid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMandantid(v)
		return nil
	}
	return fmt.Errorf("unknown Timer field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *TimerMutation) AddedFields() []string {
	var fields []string
	if m.addelapsedSeconds != nil {
		fields = append(fields, timer.FieldElapsedSeconds)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *TimerMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case timer.FieldElapsedSeconds:
		return m.AddedElapsedSeconds()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TimerMutation) AddField(name string, value ent.Value) error {
	switch name {
	case timer.FieldElapsedSeconds:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddElapsedSeconds(v)
		return nil
	}
	return fmt.Errorf("unknown Timer numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *TimerMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(timer.FieldProjectid) {
		fields = append(fields, timer.FieldProjectid)
	}
	if m.FieldCleared(timer.FieldTags) {
		fields = append(fields, timer.FieldTags)
	}
	if m.FieldCleared(timer.FieldTimerStart) {
		fields = append(fields, timer.FieldTimerStart)
	}
	if m.FieldCleared(timer.FieldTimerEnd) {
		fields = append(fields, timer.FieldTimerEnd)
	}
	if m.FieldCleared(timer.FieldElapsedSeconds) {
		fields = append(fields, timer.FieldElapsedSeconds)
	}
	if m.FieldCleared(timer.FieldIsBilled) {
		fields = append(fields, timer.FieldIsBilled)
	}
	if m.FieldCleared(timer.FieldIsRunning) {
		fields = append(fields, timer.FieldIsRunning)
	}
	if m.FieldCleared(timer.FieldUserid) {
		fields = append(fields, timer.FieldUserid)
	}
	if m.FieldCleared(timer.FieldMandantid) {
		fields = append(fields, timer.FieldMandantid)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *TimerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *TimerMutation) ClearField(name string) error {
	switch name {
	case timer.FieldProjectid:
		m.ClearProjectid()
		return nil
	case timer.FieldTags:
		m.ClearTags()
		return nil
	case timer.FieldTimerStart:
		m.ClearTimerStart()
		return nil
	case timer.FieldTimerEnd:
		m.ClearTimerEnd()
		return nil
	case timer.FieldElapsedSeconds:
		m.ClearElapsedSeconds()
		return nil
	case timer.FieldIsBilled:
		m.ClearIsBilled()
		return nil
	case timer.FieldIsRunning:
		m.ClearIsRunning()
		return nil
	case timer.FieldUserid:
		m.ClearUserid()
		return nil
	case timer.FieldMandantid:
		m.ClearMandantid()
		return nil
	}
	return fmt.Errorf("unknown Timer nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *TimerMutation) ResetField(name string) error {
	switch name {
	case timer.FieldDescription:
		m.ResetDescription()
		return nil
	case timer.FieldProjectid:
		m.ResetProjectid()
		return nil
	case timer.FieldTags:
		m.ResetTags()
		return nil
	case timer.FieldTimerStart:
		m.ResetTimerStart()
		return nil
	case timer.FieldTimerEnd:
		m.ResetTimerEnd()
		return nil
	case timer.FieldElapsedSeconds:
		m.ResetElapsedSeconds()
		return nil
	case timer.FieldIsBilled:
		m.ResetIsBilled()
		return nil
	case timer.FieldIsRunning:
		m.ResetIsRunning()
		return nil
	case timer.FieldUserid:
		m.ResetUserid()
		return nil
	case timer.FieldMandantid:
		m.ResetMandantid()
		return nil
	}
	return fmt.Errorf("unknown Timer field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *TimerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *TimerMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *TimerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *TimerMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *TimerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *TimerMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *TimerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Timer unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *TimerMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Timer edge %s", name)
}
