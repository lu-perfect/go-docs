package types

import "strings"

type VarType string

const (
	VarTypeText     VarType = "text"
	VarTypeTextarea VarType = "textarea"
	VarTypeNumber   VarType = "number"
	VarTypeDate     VarType = "date"
	VarTypeOptions  VarType = "options"
)

var varTypesMap = map[string]VarType{
	"text":     VarTypeText,
	"textarea": VarTypeTextarea,
	"number":   VarTypeNumber,
	"date":     VarTypeDate,
	"options":  VarTypeOptions,
}

func parseVarType(s string) (VarType, bool) {
	varType, ok := varTypesMap[strings.ToLower(s)]
	return varType, ok
}
