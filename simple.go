package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var numRegex = regexp.MustCompile("[0-9]+|[0-9]+.[0-9]+$")

/*
<osnovni izraz> ::= <spremenljivka> | <številka> | (<aditivni izraz>)
                  | sqrt <osnovni izraz>
				  | sin <osnovni izraz> | cos <osnovni izraz>
				  | tan <osnovni izraz> | cot <osnovni izraz>
<spremenljivka> ::= [a-zA-Z]+
<številka>      ::= -? [0-9]+
*/
func (e *Env) simple(expr string) (float64, error) {
	expr = strings.TrimSpace(expr)
	e.Calls++
	//fmt.Println(expr)

	if len(expr) == 0 {
		return 0, fmt.Errorf("empty expression")
	} else if strings.HasPrefix(expr, "(") {
		// (<aditivni izraz>)
		return e.addit(expr[1 : len(expr)-1])
	} else if strings.HasPrefix(expr, "|") {
		// |<aditivni izraz>|
		x, err := e.addit(expr[1 : len(expr)-1])
		return math.Abs(x), err
	} else if strings.HasPrefix(expr, "sqrt") {
		// sqrt(<aditivni izraz>)
		x, err := e.simple(expr[4:])
		return math.Sqrt(x), err
	} else if strings.HasPrefix(expr, "log") {
		x, err := e.simple(expr[3:])
		return math.Log10(x), err
	} else if strings.HasPrefix(expr, "sin") {
		// sin(<aditivni izraz>)
		x, err := e.simple(expr[3:])
		return math.Sin(x), err
	} else if strings.HasPrefix(expr, "cos") {
		// cos(<aditivni izraz>)
		x, err := e.simple(expr[3:])
		return math.Cos(x), err
	} else if strings.HasPrefix(expr, "tan") {
		// tan(<aditivni izraz>)
		x, err := e.simple(expr[3:])
		return math.Tan(x), err
	} else if strings.HasPrefix(expr, "cot") {
		// cot(<aditivni izraz>)
		x, err := e.simple(expr[3:])
		return 1 / math.Tan(x), err
	} else {
		if numRegex.MatchString((expr[len(expr)-1:])) {
			for i := len(expr) - 2; i >= 0; i-- {
				if !numRegex.MatchString(expr[i:]) {
					break
				}
			}
		}

		// poskusi pretvoriti v številko
		var i, err = strconv.ParseFloat(expr, 64)
		if err == nil {
			// <številka>
			return i, nil
		}

		// poskusi dobiti vrednost spremenljivke
		i, ok := e.vars[expr]
		if ok {
			// <spremenljivka>
			return i, nil
		}

		// NAPAKA
		return 0, fmt.Errorf("value of '%s' not known", expr)
	}
}
