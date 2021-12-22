package env

import (
	"fmt"
	"math"
)

/*
<potenčni izraz> ::= <osnovni izraz> | <osnovni ozraz> * <potenčni izraz>
*/
func (e *Env) power(expr string) (float64, error) {
	if len(expr) == 0 {
		return 0, fmt.Errorf("empty expression")
	}
	e.Calls++

	// poišči prvi ^
	var oklepaji, mestoZaklepaja = 0, 0
	for i := range expr {
		// preskoči oklepaje
		if expr[i] == '(' {
			oklepaji++
			mestoZaklepaja = i + 1
		} else if expr[i] == ')' {
			oklepaji--
		} else if oklepaji == 0 {
			// prišli smo do ^
			if expr[i] == '^' {
				var e1, err1 = e.simple(expr[0:i])
				var e2, err2 = e.power(expr[i+1:])

				if err1 == nil && err2 == nil {
					return math.Pow(e1, e2), nil
				}
				return 0, formatErr(err1, err2)
			}
		}
	}

	if oklepaji != 0 {
		var mestoNapake = 1
		if oklepaji > 0 {
			mestoNapake = mestoZaklepaja
		}
		return 0, fmt.Errorf("%d: bracket mismatch", mestoNapake)
	}

	return e.simple(expr)
}
