// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/obiwan007/usersrv/timersrv/api/storage/ent/timer"
)

// TimerCreate is the builder for creating a Timer entity.
type TimerCreate struct {
	config
	mutation *TimerMutation
	hooks    []Hook
}

// SetDescription sets the description field.
func (tc *TimerCreate) SetDescription(s string) *TimerCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the description field if the given value is not nil.
func (tc *TimerCreate) SetNillableDescription(s *string) *TimerCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetProjectid sets the projectid field.
func (tc *TimerCreate) SetProjectid(s string) *TimerCreate {
	tc.mutation.SetProjectid(s)
	return tc
}

// SetNillableProjectid sets the projectid field if the given value is not nil.
func (tc *TimerCreate) SetNillableProjectid(s *string) *TimerCreate {
	if s != nil {
		tc.SetProjectid(*s)
	}
	return tc
}

// SetTags sets the tags field.
func (tc *TimerCreate) SetTags(s string) *TimerCreate {
	tc.mutation.SetTags(s)
	return tc
}

// SetNillableTags sets the tags field if the given value is not nil.
func (tc *TimerCreate) SetNillableTags(s *string) *TimerCreate {
	if s != nil {
		tc.SetTags(*s)
	}
	return tc
}

// SetTimerStart sets the timerStart field.
func (tc *TimerCreate) SetTimerStart(t time.Time) *TimerCreate {
	tc.mutation.SetTimerStart(t)
	return tc
}

// SetNillableTimerStart sets the timerStart field if the given value is not nil.
func (tc *TimerCreate) SetNillableTimerStart(t *time.Time) *TimerCreate {
	if t != nil {
		tc.SetTimerStart(*t)
	}
	return tc
}

// SetTimerEnd sets the timerEnd field.
func (tc *TimerCreate) SetTimerEnd(t time.Time) *TimerCreate {
	tc.mutation.SetTimerEnd(t)
	return tc
}

// SetNillableTimerEnd sets the timerEnd field if the given value is not nil.
func (tc *TimerCreate) SetNillableTimerEnd(t *time.Time) *TimerCreate {
	if t != nil {
		tc.SetTimerEnd(*t)
	}
	return tc
}

// SetElapsedSeconds sets the elapsedSeconds field.
func (tc *TimerCreate) SetElapsedSeconds(i int) *TimerCreate {
	tc.mutation.SetElapsedSeconds(i)
	return tc
}

// SetNillableElapsedSeconds sets the elapsedSeconds field if the given value is not nil.
func (tc *TimerCreate) SetNillableElapsedSeconds(i *int) *TimerCreate {
	if i != nil {
		tc.SetElapsedSeconds(*i)
	}
	return tc
}

// SetIsBilled sets the IsBilled field.
func (tc *TimerCreate) SetIsBilled(b bool) *TimerCreate {
	tc.mutation.SetIsBilled(b)
	return tc
}

// SetNillableIsBilled sets the IsBilled field if the given value is not nil.
func (tc *TimerCreate) SetNillableIsBilled(b *bool) *TimerCreate {
	if b != nil {
		tc.SetIsBilled(*b)
	}
	return tc
}

// SetIsRunning sets the IsRunning field.
func (tc *TimerCreate) SetIsRunning(b bool) *TimerCreate {
	tc.mutation.SetIsRunning(b)
	return tc
}

// SetNillableIsRunning sets the IsRunning field if the given value is not nil.
func (tc *TimerCreate) SetNillableIsRunning(b *bool) *TimerCreate {
	if b != nil {
		tc.SetIsRunning(*b)
	}
	return tc
}

// SetUserid sets the userid field.
func (tc *TimerCreate) SetUserid(s string) *TimerCreate {
	tc.mutation.SetUserid(s)
	return tc
}

// SetNillableUserid sets the userid field if the given value is not nil.
func (tc *TimerCreate) SetNillableUserid(s *string) *TimerCreate {
	if s != nil {
		tc.SetUserid(*s)
	}
	return tc
}

// SetMandantid sets the mandantid field.
func (tc *TimerCreate) SetMandantid(s string) *TimerCreate {
	tc.mutation.SetMandantid(s)
	return tc
}

// SetNillableMandantid sets the mandantid field if the given value is not nil.
func (tc *TimerCreate) SetNillableMandantid(s *string) *TimerCreate {
	if s != nil {
		tc.SetMandantid(*s)
	}
	return tc
}

// Save creates the Timer in the database.
func (tc *TimerCreate) Save(ctx context.Context) (*Timer, error) {
	if _, ok := tc.mutation.Description(); !ok {
		v := timer.DefaultDescription
		tc.mutation.SetDescription(v)
	}
	if _, ok := tc.mutation.Projectid(); !ok {
		v := timer.DefaultProjectid
		tc.mutation.SetProjectid(v)
	}
	if _, ok := tc.mutation.Tags(); !ok {
		v := timer.DefaultTags
		tc.mutation.SetTags(v)
	}
	if _, ok := tc.mutation.ElapsedSeconds(); !ok {
		v := timer.DefaultElapsedSeconds
		tc.mutation.SetElapsedSeconds(v)
	}
	if _, ok := tc.mutation.IsBilled(); !ok {
		v := timer.DefaultIsBilled
		tc.mutation.SetIsBilled(v)
	}
	if _, ok := tc.mutation.IsRunning(); !ok {
		v := timer.DefaultIsRunning
		tc.mutation.SetIsRunning(v)
	}
	if _, ok := tc.mutation.Userid(); !ok {
		v := timer.DefaultUserid
		tc.mutation.SetUserid(v)
	}
	if _, ok := tc.mutation.Mandantid(); !ok {
		v := timer.DefaultMandantid
		tc.mutation.SetMandantid(v)
	}
	var (
		err  error
		node *Timer
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TimerCreate) SaveX(ctx context.Context) *Timer {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TimerCreate) sqlSave(ctx context.Context) (*Timer, error) {
	var (
		t     = &Timer{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: timer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: timer.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timer.FieldDescription,
		})
		t.Description = value
	}
	if value, ok := tc.mutation.Projectid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timer.FieldProjectid,
		})
		t.Projectid = value
	}
	if value, ok := tc.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timer.FieldTags,
		})
		t.Tags = value
	}
	if value, ok := tc.mutation.TimerStart(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: timer.FieldTimerStart,
		})
		t.TimerStart = value
	}
	if value, ok := tc.mutation.TimerEnd(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: timer.FieldTimerEnd,
		})
		t.TimerEnd = value
	}
	if value, ok := tc.mutation.ElapsedSeconds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: timer.FieldElapsedSeconds,
		})
		t.ElapsedSeconds = value
	}
	if value, ok := tc.mutation.IsBilled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: timer.FieldIsBilled,
		})
		t.IsBilled = value
	}
	if value, ok := tc.mutation.IsRunning(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: timer.FieldIsRunning,
		})
		t.IsRunning = value
	}
	if value, ok := tc.mutation.Userid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timer.FieldUserid,
		})
		t.Userid = value
	}
	if value, ok := tc.mutation.Mandantid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timer.FieldMandantid,
		})
		t.Mandantid = value
	}
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}
