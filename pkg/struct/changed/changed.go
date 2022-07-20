package changed

import (
	"errors"
	"reflect"
	"strings"
)

type Changed map[string]interface{}

func (c Changed) SetChanged(key string, val interface{}) {
	c[key] = val
}

func (c Changed) Changed() map[string]interface{} {
	return c
}

func (c Changed) ChangedTagMapped(tagKey string, obj interface{}) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if obj == nil {
		return result, errors.New("user is nil")
	}
	t := reflect.TypeOf(obj)
	for name, val := range c {
		field, ok := t.FieldByName(name)
		if ok {
			tag := field.Tag.Get(tagKey)
			params := strings.Split(tag, ",")
			if len(params) == 0 {
				continue
			}
			if len(params[0]) == 0 {
				continue
			}
			result[params[0]] = val
		}
	}
	return result, nil
}
