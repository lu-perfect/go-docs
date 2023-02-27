package parse

import "go-docs/types"

func Var(text string) (types.Variable, error) {
	matches := VarRegex.FindStringSubmatch(text)
	return parseVariableText(matches)
}
