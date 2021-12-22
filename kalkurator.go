package main

import (
	"bufio"
	"env"
	"fmt"
	"io"
	"os"
)

/*
BNF
<izraz>                 ::= <ime spremenljivke> = <aditivni izraz>
<aditivni izraz>        ::= <multiplikativni izraz> | <aditivni izraz>        + <multiplikativni izraz>
<multiplikativni izraz> ::= <potenčni izraz>        | <multiplikativni izraz> * <potenčni izraz>
<potenčni izraz>        ::= <osnovni izraz>         | <osnovni izraz>         ^ <potenčni izraz>
<osnovni izraz>         ::= <spremenljivka> | <številka> | (<aditivni izraz>) | '|' <osnovni izraz> '|'
                          | sqrt <osnovni izraz>
				          | sin <osnovni izraz> | cos <osnovni izraz>
				          | tan <osnovni izraz> | cot <osnovni izraz>
<spremenljivka> ::= [a-zA-Z]+
<številka>      ::= -? [0-9]+
<spremenljivka>         ::= [a-zA-Z]+
<številka>              ::= -? [0-9]+

*/

func main() {

	var args = os.Args[1:]
	var env = env.New()

	if len(args) == 0 {
		var expr string
		var reader = bufio.NewReader(os.Stdin)
		for expr != "exit" {
			// preberi naslednji izraz
			fmt.Print("> ")
			line, _, err := reader.ReadLine()
			if err != nil {
				fmt.Printf("error: %s", err)
			}
			expr = string(line)
			if len(expr) == 0 || expr == "exit" {
				break
			}

			// assign, calculate variable
			val, err := env.Assign(expr)
			if err == nil {
				fmt.Printf("= %v\n\n", val)
			} else {
				fmt.Println(err)
			}
		}

	} else if len(args) == 1 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Printf("cannot open file: %s", err)
		}

		var expressions = make([]string, 0, 10)
		var reader = bufio.NewReader(file)

		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			if len(line) > 0 {
				expressions = append(expressions, string(line))
			}
		}

		err = file.Close()
		if err != nil {
			fmt.Println("cannot close file")
		}

		var result float64
		for i, s := range expressions {
			result, err = env.Assign(s)
			if err != nil {
				fmt.Printf("syntax error in line %d: %s\n", i+1, err)
			}
		}
		fmt.Println(result)
		fmt.Printf("calls: %d\n", env.Calls)

	} else {
		fmt.Print("invalid arguments")
	}

}
