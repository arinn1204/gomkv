package utils

import "errors"

func ConcatErr(firstErr error, secondErr error) error {
	if firstErr == nil {
		return secondErr
	}

	if secondErr == nil {
		return firstErr
	}

	return errors.New(firstErr.Error() + "::" + secondErr.Error())
}
