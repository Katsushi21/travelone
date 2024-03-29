// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Katsushi21/travelone/ent/account"
	"github.com/Katsushi21/travelone/ent/mute"
	"github.com/google/uuid"
)

// Mute is the model entity for the Mute schema.
type Mute struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// MuteID holds the value of the "mute_id" field.
	MuteID uuid.UUID `json:"mute_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MuteQuery when eager-loading is set.
	Edges MuteEdges `json:"edges"`
}

// MuteEdges holds the relations/edges for other nodes in the graph.
type MuteEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// Mute holds the value of the mute edge.
	Mute *Account `json:"mute,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MuteEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// MuteOrErr returns the Mute value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MuteEdges) MuteOrErr() (*Account, error) {
	if e.loadedTypes[1] {
		if e.Mute == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Mute, nil
	}
	return nil, &NotLoadedError{edge: "mute"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mute) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mute.FieldCreatedAt, mute.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case mute.FieldID, mute.FieldAccountID, mute.FieldMuteID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Mute", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mute fields.
func (m *Mute) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mute.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case mute.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case mute.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case mute.FieldAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				m.AccountID = *value
			}
		case mute.FieldMuteID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field mute_id", values[i])
			} else if value != nil {
				m.MuteID = *value
			}
		}
	}
	return nil
}

// QueryAccount queries the "account" edge of the Mute entity.
func (m *Mute) QueryAccount() *AccountQuery {
	return (&MuteClient{config: m.config}).QueryAccount(m)
}

// QueryMute queries the "mute" edge of the Mute entity.
func (m *Mute) QueryMute() *AccountQuery {
	return (&MuteClient{config: m.config}).QueryMute(m)
}

// Update returns a builder for updating this Mute.
// Note that you need to call Mute.Unwrap() before calling this method if this Mute
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mute) Update() *MuteUpdateOne {
	return (&MuteClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Mute entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mute) Unwrap() *Mute {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mute is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mute) String() string {
	var builder strings.Builder
	builder.WriteString("Mute(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", m.AccountID))
	builder.WriteString(", ")
	builder.WriteString("mute_id=")
	builder.WriteString(fmt.Sprintf("%v", m.MuteID))
	builder.WriteByte(')')
	return builder.String()
}

// Mutes is a parsable slice of Mute.
type Mutes []*Mute

func (m Mutes) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
