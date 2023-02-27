package types

import (
	"fmt"
	"strconv"
	"strings"
)

type Variable struct {
	Name    string
	Type    VarType
	Options map[string]string
}

func NewVariable(name string) Variable {
	return Variable{
		Name:    name,
		Type:    VarTypeText,
		Options: make(map[string]string),
	}
}

func (v *Variable) SetType(s string) error {
	t, ok := parseVarType(s)
	if !ok {
		err := fmt.Errorf("unsupported variable type: %s", t)
		return err
	}

	v.Type = t
	return nil
}

func (v *Variable) SetOptions(s string) {
	options := strings.Split(s, ",")
	for _, option := range options {
		option = strings.Replace(option, "::", "", 1)
		parts := strings.Split(option, "=")
		key := parts[0]

		if len(parts) == 1 {
			v.Options[key] = "true"
		}
		if len(parts) == 2 {
			value := parts[1]
			v.Options[key] = value
		}
	}
}

func (v *Variable) Validate(match string, value string) (string, error) {
	if v.Type == VarTypeNumber {
		num, err := strconv.Atoi(value)
		if err != nil {
			return match, err
		}
		minStr, ok := v.Options["min"]
		if ok {
			min, err := strconv.Atoi(minStr)
			if err != nil || num < min {
				err := fmt.Errorf("variable %s less than min: %d", v.Name, min)
				return match, err
			}
		}
		maxStr, ok := v.Options["max"]
		if ok {
			max, err := strconv.Atoi(maxStr)
			if err != nil || num > max {
				err := fmt.Errorf("variable %s more than max: %d", v.Name, max)
				return match, err
			}
		}
		return value, nil
	}
	if v.Type == VarTypeText {
		if _, ok := v.Options["uppercase"]; ok {
			return strings.ToUpper(value), nil
		}
		if _, ok := v.Options["lowercase"]; ok {
			return strings.ToLower(value), nil
		}
		return value, nil
	}

	return value, nil
}
