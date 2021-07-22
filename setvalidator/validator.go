package setvalidator

import (
	"fmt"
	"reflect"
	"time"
)

// Validator
func Validator(args ...interface{}) error {
	for _, val := range args {
		if val == nil {
			return fmt.Errorf("value is nil")
		}
		if t, ok := val.(time.Time); ok {
			if t.String() == "" {
				return fmt.Errorf("time %v is empty", t)
			}
		}

		refVal := reflect.ValueOf(val)

		switch refVal.Kind() {
		case reflect.String:
			if refVal.String() == "" {
				return fmt.Errorf("val %v is empty", refVal)
			}
		}
	}
	return nil
}
