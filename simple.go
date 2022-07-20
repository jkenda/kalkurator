package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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
		if unicode.IsDigit(rune(expr[0])) {
			// poskusi pretvoriti v številko
			var i, err = strconv.ParseFloat(expr, 64)
			if err == nil {
				// <številka>
				return i, nil
			}
			// če ne uspe, jo poskusi razdeliti
			for i, c := range expr {
				if !unicode.IsDigit(c) && c != '.' {
					var e1, err1 = e.simple(expr[:i])
					var e2, err2 = e.simple(expr[i:])
					if err1 == nil && err2 == nil {
						return e1 * e2, nil
					}
					return 0, formatErr(err1, err2)
				}
			}
		} else {
			// poskusi dobiti vrednost spremenljivke
			i, ok := e.vars[expr]
			if ok {
				// <spremenljivka>
				return i, nil
			}

			// če ne uspe, jo poskusi razdeliti
			for i, c := range expr {
				if c == ' ' || c == '(' {
					var e1, err1 = e.simple(expr[:i])
					var e2, err2 = e.simple(expr[i:])
					if err1 == nil && err2 == nil {
						return e1 * e2, nil
					}
					return 0, formatErr(err1, err2)
				}
			}
		}

		// NAPAKA
		return 0, fmt.Errorf("value of '%s' not known", expr)
	}
}
