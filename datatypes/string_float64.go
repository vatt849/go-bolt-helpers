package types

import (
	"encoding/json"
	"strconv"
)

// StringFloat64 create a type alias for type float64
type StringFloat64 float64

// UnmarshalJSON create a custom unmarshal for the StringInt
/// this helps us check the type of our value before unmarshalling it

func (st *StringFloat64) UnmarshalJSON(b []byte) error {
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
		*st = StringFloat64(float64(v))
	case float64:
		*st = StringFloat64(v)
	case string:
		///here convert the string into
		///an float64
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			///the string might not be of float64 type
			///so return an error
			return err

		}
		*st = StringFloat64(i)

	}
	return nil
}
