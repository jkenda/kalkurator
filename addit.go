package main

import (
	"fmt"
)

/*
 <aditivni izraz> ::= <multiplikativni izraz> | <aditivni izraz> + <multiplikativni izraz>
*/
func (e *Env) addit(expr string) (float64, error) {
	if len(expr) == 0 {
		return 0, fmt.Errorf("empty expression")
	}
	e.Calls++

	var oklepaji, mestoZaklepaja, znotrajAbs = 0, 0, false
	// poišči zadnji + ali -
	for i := len(expr) - 1; i >= 0; i-- {
		// preskoči oklepaje
		if expr[i] == ')' {
			oklepaji++
			mestoZaklepaja = i + 1
		} else if expr[i] == '(' {
			oklepaji--
		} else if expr[i] == '|' {
			znotrajAbs = !znotrajAbs
		} else if oklepaji == 0 && !znotrajAbs {
			// prišli smo do +
			if expr[i] == '+' {
				var e1, err1 = e.addit(expr[0:i])
				var e2, err2 = e.multi(expr[i+1:])

				if err1 == nil && err2 == nil {
					return e1 + e2, nil
				} else if err1 != nil && err1.Error() == "empty expression" {
					return e2, nil
				} else {
					return 0, formatErr(err1, err2)
				}
				// prišli smo do -
			} else if expr[i] == '-' {
				var e1, err1 = e.addit(expr[:i])
				var e2, err2 = e.multi(expr[i+1:])

				if err1 == nil && err2 == nil {
					return e1 - e2, nil
				} else if err1 != nil && err1.Error() == "empty expression" {
					return -e2, nil
				} else {
					return 0, formatErr(err1, err2)
				}
			}
		}
	}

	if oklepaji != 0 {
		var mestoNapake = 1
		if oklepaji > 0 {
			mestoNapake = mestoZaklepaja
		}
		return 0, fmt.Errorf("%d bracket mismatch", mestoNapake)
	}

	return e.multi(expr)
}
