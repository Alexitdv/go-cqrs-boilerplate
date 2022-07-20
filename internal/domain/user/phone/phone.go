package phone

import (
	"regexp"

	"github.com/pkg/errors"
)

type Phone string

func (p Phone) String() string {
	return string(p)
}

func (p Phone) Hidden() string {
	runes := []rune(p)
	for i := 5; i < len(runes)-2; i++ {
		runes[i] = '*'
	}
	return string(runes)
}

func NewPhone(str string) (Phone, error) {
	matched, _ := regexp.Match(`\+[0-9]{11,15}`, []byte(str))

	if !matched {
		return "", errors.New(
			"incorrect phone number, it should start from + and contains only numbers, min size is 11, max size is 15",
		)
	}

	return Phone(str), nil
}
