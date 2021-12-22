package env

import (
	"fmt"
	"regexp"
	"strings"
)

var varRegex = regexp.MustCompile("[a-z][a-zA-Z0-9]*$")

// Env - struktura za okolje
type Env struct {
	vars  map[string]float64
	Calls int
}

// New - konstruktor
func New() Env {
	var e Env
	e.vars = map[string]float64{
		"e":   2.7182818284590452,
		"phi": 1.61803398874989485,
		"pi":  3.1415926535897932,
	}
	e.Calls = 0
	return e
}

// Eval - vrne vrednost izraza
func (e *Env) Eval(expr string) (float64, error) {
	return e.addit(expr)
}

// Assign - priredi vrednost spremenljivke
func (e *Env) Assign(expr string) (float64, error) {

	var name, isAssignment = "", false
	for i, c := range expr {
		if c == '=' {
			name, expr, isAssignment = expr[0:i], expr[i+1:], true
			break
		}
	}

	value, err := e.Eval(expr)

	if err != nil {
		return 0, err
	}

	if isAssignment {
		name = strings.TrimSpace(name)
		if !varRegex.MatchString(name) {
			return 0, fmt.Errorf("'%s' is not a valid variable name", name)
		}
		e.vars[name] = value
	}
	return value, nil
}
