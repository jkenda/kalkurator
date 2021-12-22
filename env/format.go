package env

import (
	"errors"
)

func formatErr(err1 error, err2 error) error {
	var errStr = ""
	if err1 != nil {
		errStr += err1.Error()
	}
	if err2 != nil {
		if err1 != nil {
			errStr += ", "
		}
		errStr += err2.Error()
	}
	return errors.New(errStr)
}
