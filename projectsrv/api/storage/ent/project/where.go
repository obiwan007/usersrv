// Code generated by entc, DO NOT EDIT.

package project

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/obiwan007/usersrv/projectsrv/api/storage/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Clientid applies equality check predicate on the "clientid" field. It's identical to ClientidEQ.
func Clientid(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientid), v))
	})
}

// Userid applies equality check predicate on the "userid" field. It's identical to UseridEQ.
func Userid(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserid), v))
	})
}

// Mandantid applies equality check predicate on the "mandantid" field. It's identical to MandantidEQ.
func Mandantid(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMandantid), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Tags applies equality check predicate on the "tags" field. It's identical to TagsEQ.
func Tags(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTags), v))
	})
}

// Team applies equality check predicate on the "team" field. It's identical to TeamEQ.
func Team(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeam), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ClientidEQ applies the EQ predicate on the "clientid" field.
func ClientidEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientid), v))
	})
}

// ClientidNEQ applies the NEQ predicate on the "clientid" field.
func ClientidNEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientid), v))
	})
}

// ClientidIn applies the In predicate on the "clientid" field.
func ClientidIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClientid), v...))
	})
}

// ClientidNotIn applies the NotIn predicate on the "clientid" field.
func ClientidNotIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClientid), v...))
	})
}

// ClientidGT applies the GT predicate on the "clientid" field.
func ClientidGT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientid), v))
	})
}

// ClientidGTE applies the GTE predicate on the "clientid" field.
func ClientidGTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientid), v))
	})
}

// ClientidLT applies the LT predicate on the "clientid" field.
func ClientidLT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientid), v))
	})
}

// ClientidLTE applies the LTE predicate on the "clientid" field.
func ClientidLTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientid), v))
	})
}

// ClientidContains applies the Contains predicate on the "clientid" field.
func ClientidContains(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientid), v))
	})
}

// ClientidHasPrefix applies the HasPrefix predicate on the "clientid" field.
func ClientidHasPrefix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientid), v))
	})
}

// ClientidHasSuffix applies the HasSuffix predicate on the "clientid" field.
func ClientidHasSuffix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientid), v))
	})
}

// ClientidIsNil applies the IsNil predicate on the "clientid" field.
func ClientidIsNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClientid)))
	})
}

// ClientidNotNil applies the NotNil predicate on the "clientid" field.
func ClientidNotNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClientid)))
	})
}

// ClientidEqualFold applies the EqualFold predicate on the "clientid" field.
func ClientidEqualFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientid), v))
	})
}

// ClientidContainsFold applies the ContainsFold predicate on the "clientid" field.
func ClientidContainsFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientid), v))
	})
}

// UseridEQ applies the EQ predicate on the "userid" field.
func UseridEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserid), v))
	})
}

// UseridNEQ applies the NEQ predicate on the "userid" field.
func UseridNEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserid), v))
	})
}

// UseridIn applies the In predicate on the "userid" field.
func UseridIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserid), v...))
	})
}

// UseridNotIn applies the NotIn predicate on the "userid" field.
func UseridNotIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserid), v...))
	})
}

// UseridGT applies the GT predicate on the "userid" field.
func UseridGT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserid), v))
	})
}

// UseridGTE applies the GTE predicate on the "userid" field.
func UseridGTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserid), v))
	})
}

// UseridLT applies the LT predicate on the "userid" field.
func UseridLT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserid), v))
	})
}

// UseridLTE applies the LTE predicate on the "userid" field.
func UseridLTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserid), v))
	})
}

// UseridContains applies the Contains predicate on the "userid" field.
func UseridContains(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserid), v))
	})
}

// UseridHasPrefix applies the HasPrefix predicate on the "userid" field.
func UseridHasPrefix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserid), v))
	})
}

// UseridHasSuffix applies the HasSuffix predicate on the "userid" field.
func UseridHasSuffix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserid), v))
	})
}

// UseridIsNil applies the IsNil predicate on the "userid" field.
func UseridIsNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserid)))
	})
}

// UseridNotNil applies the NotNil predicate on the "userid" field.
func UseridNotNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserid)))
	})
}

// UseridEqualFold applies the EqualFold predicate on the "userid" field.
func UseridEqualFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserid), v))
	})
}

// UseridContainsFold applies the ContainsFold predicate on the "userid" field.
func UseridContainsFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserid), v))
	})
}

// MandantidEQ applies the EQ predicate on the "mandantid" field.
func MandantidEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMandantid), v))
	})
}

// MandantidNEQ applies the NEQ predicate on the "mandantid" field.
func MandantidNEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMandantid), v))
	})
}

// MandantidIn applies the In predicate on the "mandantid" field.
func MandantidIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMandantid), v...))
	})
}

// MandantidNotIn applies the NotIn predicate on the "mandantid" field.
func MandantidNotIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMandantid), v...))
	})
}

// MandantidGT applies the GT predicate on the "mandantid" field.
func MandantidGT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMandantid), v))
	})
}

// MandantidGTE applies the GTE predicate on the "mandantid" field.
func MandantidGTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMandantid), v))
	})
}

// MandantidLT applies the LT predicate on the "mandantid" field.
func MandantidLT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMandantid), v))
	})
}

// MandantidLTE applies the LTE predicate on the "mandantid" field.
func MandantidLTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMandantid), v))
	})
}

// MandantidContains applies the Contains predicate on the "mandantid" field.
func MandantidContains(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMandantid), v))
	})
}

// MandantidHasPrefix applies the HasPrefix predicate on the "mandantid" field.
func MandantidHasPrefix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMandantid), v))
	})
}

// MandantidHasSuffix applies the HasSuffix predicate on the "mandantid" field.
func MandantidHasSuffix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMandantid), v))
	})
}

// MandantidIsNil applies the IsNil predicate on the "mandantid" field.
func MandantidIsNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMandantid)))
	})
}

// MandantidNotNil applies the NotNil predicate on the "mandantid" field.
func MandantidNotNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMandantid)))
	})
}

// MandantidEqualFold applies the EqualFold predicate on the "mandantid" field.
func MandantidEqualFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMandantid), v))
	})
}

// MandantidContainsFold applies the ContainsFold predicate on the "mandantid" field.
func MandantidContainsFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMandantid), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// TagsEQ applies the EQ predicate on the "tags" field.
func TagsEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTags), v))
	})
}

// TagsNEQ applies the NEQ predicate on the "tags" field.
func TagsNEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTags), v))
	})
}

// TagsIn applies the In predicate on the "tags" field.
func TagsIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTags), v...))
	})
}

// TagsNotIn applies the NotIn predicate on the "tags" field.
func TagsNotIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTags), v...))
	})
}

// TagsGT applies the GT predicate on the "tags" field.
func TagsGT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTags), v))
	})
}

// TagsGTE applies the GTE predicate on the "tags" field.
func TagsGTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTags), v))
	})
}

// TagsLT applies the LT predicate on the "tags" field.
func TagsLT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTags), v))
	})
}

// TagsLTE applies the LTE predicate on the "tags" field.
func TagsLTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTags), v))
	})
}

// TagsContains applies the Contains predicate on the "tags" field.
func TagsContains(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTags), v))
	})
}

// TagsHasPrefix applies the HasPrefix predicate on the "tags" field.
func TagsHasPrefix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTags), v))
	})
}

// TagsHasSuffix applies the HasSuffix predicate on the "tags" field.
func TagsHasSuffix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTags), v))
	})
}

// TagsIsNil applies the IsNil predicate on the "tags" field.
func TagsIsNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTags)))
	})
}

// TagsNotNil applies the NotNil predicate on the "tags" field.
func TagsNotNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTags)))
	})
}

// TagsEqualFold applies the EqualFold predicate on the "tags" field.
func TagsEqualFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTags), v))
	})
}

// TagsContainsFold applies the ContainsFold predicate on the "tags" field.
func TagsContainsFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTags), v))
	})
}

// TeamEQ applies the EQ predicate on the "team" field.
func TeamEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeam), v))
	})
}

// TeamNEQ applies the NEQ predicate on the "team" field.
func TeamNEQ(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTeam), v))
	})
}

// TeamIn applies the In predicate on the "team" field.
func TeamIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTeam), v...))
	})
}

// TeamNotIn applies the NotIn predicate on the "team" field.
func TeamNotIn(vs ...string) predicate.Project {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Project(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTeam), v...))
	})
}

// TeamGT applies the GT predicate on the "team" field.
func TeamGT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTeam), v))
	})
}

// TeamGTE applies the GTE predicate on the "team" field.
func TeamGTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTeam), v))
	})
}

// TeamLT applies the LT predicate on the "team" field.
func TeamLT(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTeam), v))
	})
}

// TeamLTE applies the LTE predicate on the "team" field.
func TeamLTE(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTeam), v))
	})
}

// TeamContains applies the Contains predicate on the "team" field.
func TeamContains(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTeam), v))
	})
}

// TeamHasPrefix applies the HasPrefix predicate on the "team" field.
func TeamHasPrefix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTeam), v))
	})
}

// TeamHasSuffix applies the HasSuffix predicate on the "team" field.
func TeamHasSuffix(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTeam), v))
	})
}

// TeamIsNil applies the IsNil predicate on the "team" field.
func TeamIsNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTeam)))
	})
}

// TeamNotNil applies the NotNil predicate on the "team" field.
func TeamNotNil() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTeam)))
	})
}

// TeamEqualFold applies the EqualFold predicate on the "team" field.
func TeamEqualFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTeam), v))
	})
}

// TeamContainsFold applies the ContainsFold predicate on the "team" field.
func TeamContainsFold(v string) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTeam), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		p(s.Not())
	})
}
