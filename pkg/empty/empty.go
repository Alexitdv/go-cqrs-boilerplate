package empty

import (
	"reflect"

	"github.com/pkg/errors"
)

type Valuer interface {
	EmptyValue() (value interface{}, empty bool)
}

func GetValues(obj interface{}) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		return result, errors.New("input is not a struct")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsZero() {
			iFace := field.Interface()
			if valuer, ok := iFace.(Valuer); ok {
				r, empty := valuer.EmptyValue()
				if !empty {
					result[t.Field(i).Name] = r
				}
				continue
			}

			result[t.Field(i).Name] = field.Interface()
		}
	}

	return result, nil
}
