package cmd

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	VarName     string
	ErrorType   string // "missing", "type_mismatch", "invalid_format"
	ExpectedType string
	ActualValue string
	Hint        string
	Example     string
}

func (e *ValidationError) Error() string {
	var b strings.Builder
	
	b.WriteString(fmt.Sprintf("Invalid variable: %s\n", e.VarName))
	
	switch e.ErrorType {
	case "missing":
		b.WriteString("   Status: Required variable is missing\n")
		b.WriteString("   \n")
		b.WriteString("   This variable must be set for the application to run.\n")
		if e.Hint != "" {
			b.WriteString(fmt.Sprintf("   %s\n", e.Hint))
		}
		
	case "type_mismatch":
		b.WriteString(fmt.Sprintf("   Expected type: %s\n", e.ExpectedType))
		b.WriteString(fmt.Sprintf("   Got: \"%s\"\n", e.ActualValue))
		b.WriteString("   \n")
		if e.Hint != "" {
			b.WriteString(fmt.Sprintf("   %s\n", e.Hint))
		}
		
	case "invalid_format":
		b.WriteString(fmt.Sprintf("   Value: \"%s\"\n", e.ActualValue))
		b.WriteString("   \n")
		b.WriteString(fmt.Sprintf("   %s\n", e.Hint))

	case "unknown_type":
		b.WriteString(fmt.Sprintf("  unknown type: %s\n", e.ErrorType))
		b.WriteString("   \n")
		b.WriteString(fmt.Sprintf("   %s\n", e.Hint))
	}

	
	
	if e.Example != "" {
		b.WriteString("   \n")
		b.WriteString(fmt.Sprintf("   Example: export %s=%s\n", e.VarName, e.Example))
	}
	
	return b.String()
}

// Helper constructors
func NewMissingVarError(varName, hint, example string) *ValidationError {
	return &ValidationError{
		VarName:   varName,
		ErrorType: "missing",
		Hint:      hint,
		Example:   example,
	}
}

func NewTypeMismatchError(varName, expectedType, actualValue, hint, example string) *ValidationError {
	return &ValidationError{
		VarName:      varName,
		ErrorType:    "type_mismatch",
		ExpectedType: expectedType,
		ActualValue:  actualValue,
		Hint:         hint,
		Example:      example,
	}
}

func NewInvalidFormatError(varName, actualValue, hint, example string) *ValidationError {
	return &ValidationError{
		VarName:     varName,
		ErrorType:   "invalid_format",
		ActualValue: actualValue,
		Hint:        hint,
		Example:     example,
	}
}

func NewUnknownTypeError(varName, actualValue, hint, example string) *ValidationError {
	return &ValidationError{
		VarName:     varName,
		ErrorType:   "unknown type",
		ActualValue: actualValue,
		Hint:        hint,
		Example:     example,
	}
}