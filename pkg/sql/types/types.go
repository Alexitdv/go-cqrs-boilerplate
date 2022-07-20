package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type JsonStringArray []string

func NewJsonStringArray(values []string) *JsonStringArray {
	if values == nil {
		typ := JsonStringArray([]string{})
		return &typ
	}

	check := make(map[string]struct{})
	res := make([]string, 0)
	for _, val := range values {
		check[val] = struct{}{}
	}
	for val, _ := range check {
		res = append(res, val)
	}

	typ := JsonStringArray(res)
	return &typ
}

func (s *JsonStringArray) Append(value string) {
	for _, v := range *s {
		if v == value {
			return
		}
	}
	*s = append(*s, value)
}

func (s *JsonStringArray) GetValue() []string {
	return *s
}

func (s *JsonStringArray) Json() string {
	b, err := json.Marshal(s)
	if err != nil {
		return "null"
	}
	return string(b)
}

func (s *JsonStringArray) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		err := json.Unmarshal(v, &s)
		if err != nil {
			return err
		}
		return nil
	case string:
		err := json.Unmarshal([]byte(v), &s)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (s *JsonStringArray) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
