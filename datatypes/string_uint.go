package types

import (
	"encoding/json"
	"strconv"
)

// StringInt create a type alias for type int
type StringUint uint

// UnmarshalJSON create a custom unmarshal for the StringInt
/// this helps us check the type of our value before unmarshalling it

func (st *StringUint) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	//if it is a string that can be converted into an int we convert it
	///otherwise we return an error
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringUint(uint(v))
	case float64:
		*st = StringUint(uint(v))
	case string:
		///here convert the string into
		///an integer
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err

		}
		*st = StringUint(i)

	}
	return nil
}
