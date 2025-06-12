package validation

import "errors"

func ComparePassword(password, confirm string) error {
	if password != confirm {
		return errors.New("password not match")
	}
	return nil
}
