package datatypes

import (
	"database/sql/driver"
	"fmt"
)

type BitBool bool

func (bb BitBool) Value() (driver.Value, error) {
	return bool(bb), nil
}

func (bb *BitBool) Scan(src interface{}) error {
	if src == nil {
		*bb = false
		return nil
	}

	bs, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("value is not byte slice")
	}
	*bb = bs[0] == 1
	return nil
}
