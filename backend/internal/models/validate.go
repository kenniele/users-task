package models

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func PassportValidation(passport any) error {
	if reflect.TypeOf(passport).Kind() != reflect.String {
		return errors.New("password must be a string")
	}
	return func(passport string) error {
		split := strings.Split(passport, " ")
		if len(split) != 2 {
			return errors.New("passport must contain two colons")
		}
		if len(split[0]) != 4 || len(split[1]) != 6 {
			return errors.New("bad format of passport")
		}
		if _, err := strconv.Atoi(split[0]); err != nil {
			return errors.New("passport must contain integers")
		}
		if _, err := strconv.Atoi(split[1]); err != nil {
			return errors.New("passport must contain integers")
		}
		return nil
	}(passport.(string))
}
