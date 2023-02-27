package main

import (
	"fmt"
	"go-docs/parse"
	"log"
)

func putData(text string, data map[string]string) (string, error) {
	var resErr error
	result := parse.VarRegex.ReplaceAllStringFunc(text, func(match string) string {
		v, err := parse.Var(match)
		if err != nil {
			return match
		}

		value, ok := data[v.Name]
		if !ok {
			return match
		}

		res, err := v.Validate(match, value)
		if err != nil {
			resErr = err
		}

		return res
	})

	return result, resErr
}

func main() {
	text := "Mr. ${name::text::uppercase} has ${value::number::min=0} $."
	vars, err := parse.Vars(text)
	if err != nil {
		err := fmt.Errorf("cannot parse variables, err: %w", err)
		log.Fatal(err)
	}
	fmt.Println("Variables found:", vars)

	data := map[string]string{
		"name":  "Sam",
		"value": "200000",
	}
	replaced, err := putData(text, data)
	if err != nil {
		err := fmt.Errorf("cannot replace variables, err: %w", err)
		log.Fatal(err)
	}
	fmt.Println("Replaced text:", replaced)
}
