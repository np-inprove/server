// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/course"
	"github.com/np-inprove/server/internal/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// First name of the user
	FirstName string `json:"first_name,omitempty"`
	// Last name of the user
	LastName string `json:"last_name,omitempty"`
	// Email of the user
	Email string `json:"email,omitempty"`
	// Password hash of the user
	PasswordHash string `json:"-"`
	// Points of the user.
	// Must always be positive
	Points int `json:"points,omitempty"`
	// Points awarded by the user after points_awarded_reset_time.
	// Must always be positive
	PointsAwardedCount int `json:"points_awarded_count,omitempty"`
	// Time when points_awarded_count was last reset to 0
	PointsAwardedResetTime time.Time `json:"points_awarded_reset_time,omitempty"`
	// Superuser of the iNProve platform
	Superuser bool `json:"superuser,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges           UserEdges `json:"edges"`
	course_students *int
	selectValues    sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Institution holds the value of the institution edge.
	Institution []*Institution `json:"institution,omitempty"`
	// Course holds the value of the course edge.
	Course *Course `json:"course,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// InstitutionOrErr returns the Institution value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) InstitutionOrErr() ([]*Institution, error) {
	if e.loadedTypes[0] {
		return e.Institution, nil
	}
	return nil, &NotLoadedError{edge: "institution"}
}

// CourseOrErr returns the Course value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CourseOrErr() (*Course, error) {
	if e.loadedTypes[1] {
		if e.Course == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: course.Label}
		}
		return e.Course, nil
	}
	return nil, &NotLoadedError{edge: "course"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldSuperuser:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldPoints, user.FieldPointsAwardedCount:
			values[i] = new(sql.NullInt64)
		case user.FieldFirstName, user.FieldLastName, user.FieldEmail, user.FieldPasswordHash:
			values[i] = new(sql.NullString)
		case user.FieldPointsAwardedResetTime:
			values[i] = new(sql.NullTime)
		case user.ForeignKeys[0]: // course_students
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				u.PasswordHash = value.String
			}
		case user.FieldPoints:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field points", values[i])
			} else if value.Valid {
				u.Points = int(value.Int64)
			}
		case user.FieldPointsAwardedCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field points_awarded_count", values[i])
			} else if value.Valid {
				u.PointsAwardedCount = int(value.Int64)
			}
		case user.FieldPointsAwardedResetTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field points_awarded_reset_time", values[i])
			} else if value.Valid {
				u.PointsAwardedResetTime = value.Time
			}
		case user.FieldSuperuser:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field superuser", values[i])
			} else if value.Valid {
				u.Superuser = value.Bool
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field course_students", value)
			} else if value.Valid {
				u.course_students = new(int)
				*u.course_students = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryInstitution queries the "institution" edge of the User entity.
func (u *User) QueryInstitution() *InstitutionQuery {
	return NewUserClient(u.config).QueryInstitution(u)
}

// QueryCourse queries the "course" edge of the User entity.
func (u *User) QueryCourse() *CourseQuery {
	return NewUserClient(u.config).QueryCourse(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password_hash=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("points=")
	builder.WriteString(fmt.Sprintf("%v", u.Points))
	builder.WriteString(", ")
	builder.WriteString("points_awarded_count=")
	builder.WriteString(fmt.Sprintf("%v", u.PointsAwardedCount))
	builder.WriteString(", ")
	builder.WriteString("points_awarded_reset_time=")
	builder.WriteString(u.PointsAwardedResetTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("superuser=")
	builder.WriteString(fmt.Sprintf("%v", u.Superuser))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
