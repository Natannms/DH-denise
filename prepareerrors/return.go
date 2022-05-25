package prepareerrors

import "errors"

func One() any {
	a := 0

	if a == 1 {
		err := errors.New("a is 1")
		return err
	}

	return nil
}
