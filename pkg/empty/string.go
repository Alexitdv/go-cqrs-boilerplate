package empty

import (
	"database/sql/driver"

	"github.com/pkg/errors"
)

type String struct {
	String string
	Set    bool // Set is true if String is set
}

// Scan implements the Scanner interface.
func (ns *String) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Set = "", false
		return nil
	}

	ns.Set = true // TODO: Pay attention
	if sv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := sv.(string); ok {
			// set the value of the pointer yne to YesNoEnum(v)
			ns.String = v
			return nil
		}
	}

	return errors.New("failed to scan Empty.String")
}

// Value implements the driver Valuer interface.
func (ns String) Value() (driver.Value, error) {
	return ns.String, nil
}
