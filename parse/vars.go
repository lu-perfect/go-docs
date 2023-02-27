package parse

import "go-docs/types"

func Vars(text string) ([]types.Variable, error) {
	matches := VarRegex.FindAllStringSubmatch(text, -1)
	vars := make([]types.Variable, 0, len(matches))
	for _, match := range matches {
		v, err := parseVariableText(match)
		if err != nil {
			return nil, err
		}
		vars = append(vars, v)
	}
	return vars, nil
}
