package parse

import (
	"go-docs/types"
	"regexp"
)

var VarRegex = regexp.MustCompile(`\${(\w+)(::(\w+)(::([\w,=]+))?)?}`)

func parseVariableText(matches []string) (types.Variable, error) {
	if len(matches) < 2 {
		return types.Variable{}, nil
	}
	v := types.NewVariable(matches[1])
	if len(matches) > 3 && matches[3] != "" {
		err := v.SetType(matches[3])
		if err != nil {
			return v, err
		}
	}
	if len(matches) > 4 && matches[4] != "" {
		v.SetOptions(matches[4])
	}
	return v, nil
}
