package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var inputFile string
var strict bool
var verbose bool

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a template schema.yaml file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ValidateEnvSchema(inputFile)
	},
	Long: `Validate current environment variables against a schema file.

	The validator checks:
	• If all required variables are set
	• If the variable types match the schema (int, bool, string, url)
	• Applies default values when specified
	• Optionally enforces strict mode (no unexpected variables)

	Exit codes:
	0 - All validations passed
	1 - Validation failed (missing/invalid variables)`,

	Example: `  # Basic validation
	envcheck validate

	# Use custom schema file
	envcheck validate --schema prod-schema.yaml

	# Strict mode: fail if unexpected APP_* variables exist
	envcheck validate --strict

	# Show detailed information about each variable
	envcheck validate --verbose

	# Combine flags for production verification
	envcheck validate --schema prod.yaml --strict --verbose`,
}

//create a function that makes the defined command resuable

func ValidateCmd() *cobra.Command {
	return validateCmd
}

func init() {
	//create a flag for the schema file
	ValidateCmd().Flags().StringVarP(
		&inputFile,
		"schema",
		"s",
		"schema.yaml",
		"schema file path",
	)
	//create a flag for the strict mode
	ValidateCmd().Flags().BoolVarP(
		&strict,
		"strict",
		"t",
		false,
		"strict mode",
	)
	//create a flag for the verbose mode

	ValidateCmd().Flags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"verbose mode",
	)
}

func ValidateEnvSchema(fileName string) error {
	dataFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading schema.yaml file")
		return nil
	}

	//to unMarshal the data into a struct

	var config Config

	err = yaml.Unmarshal(dataFile, &config)

	if err != nil {
		log.Fatalf("error reading unmarshing file")
		return nil
	}

	allowedVars := make(map[string]bool)

	hasErrors := false
	for _, envVarRule := range config.Variables {
		//loop through the slice of variables env rules
		if strict {
			for _, envVarRule := range config.Variables {
				allowedVars[envVarRule.Name] = true
			}
			for _, env := range os.Environ() {
				parts := strings.SplitN(env, "=", 2)
				if len(parts) == 2 {
					key := parts[0]
					if strings.HasPrefix(key, "APP") {
						if !allowedVars[key] {
							fmt.Fprintf(os.Stderr, "❌ %s: %v\n", key, fmt.Errorf("unknown variable: %s", key))
							hasErrors = true
						}
					}
				}
			}

		}
		err := validateVar(envVarRule)

		if err != nil {
			//write our custom error formatter to the std err output if the error type
			if validationErr, ok := err.(*ValidationError); ok {
				fmt.Fprintf(os.Stderr, "\n%s\n", validationErr.Error()) /// call the formatter error
			} else {
				fmt.Fprintf(os.Stderr, "❌ %s: %v\n", envVarRule.Name, err)
				hasErrors = true
			}
		}
		if hasErrors {
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "💡 Tip: Run 'envcheck validate --verbose' to see more details")
			fmt.Fprintln(os.Stderr, "💡 Tip: Check your .env file or export variables manually")
			return fmt.Errorf("validation failed")
		}
	}
	fmt.Fprintln(os.Stdout, "\n✓ All environment variables validated successfully")
	return nil
	//unMarshal
}

func validateVar(envVar EnvRules) error {
	envValue := os.Getenv(envVar.Name)
	requiredStr := ""
	if envVar.Default != "" {
		requiredStr = "from default"
	}
	if envVar.Default != "" && envValue == "" {
		envValue = envVar.Default
	}
	if envVar.Required && envValue == "" {
		hint := "Set this variable in your environment or .env file"
		example := getExampleValue(envVar)
		return NewMissingVarError(envVar.Name, hint, example)
	}
	// Validate type
	hint := "Set your type in your enviroment to the correct one"
	switch envVar.Type {
	case "int":
		if _, err := strconv.Atoi(envValue); err != nil {
			return NewTypeMismatchError(envVar.Name, "int", envValue, hint, getExampleValue(envVar))
		}
	case "bool":
		if _, err := strconv.ParseBool(envValue); err != nil {
			return NewTypeMismatchError(envVar.Name, "bool", envValue, hint, getExampleValue(envVar))
		}
	case "string":
		// Always valid
	default:
		return NewUnknownTypeError(envVar.Name, envValue, hint, getExampleValue(envVar))
	}

	if verbose {
		fmt.Fprint(os.Stdout, "\n✓", envVar.Name)
		fmt.Fprintf(os.Stdout, "\n Value: %s %v", envValue, (requiredStr))
		fmt.Fprintln(os.Stdout, "\n Type:", envVar.Type)
	}

	return nil
}

// Helper function to generate example values based on variable name and type
func getExampleValue(envVar EnvRules) string {
	// Smart examples based on variable name
	name := strings.ToLower(envVar.Name)

	switch {
	case strings.Contains(name, "database") || strings.Contains(name, "db"):
		return "\"postgres://user:pass@localhost:5432/dbname\""
	case strings.Contains(name, "port"):
		return "8080"
	case strings.Contains(name, "url") || strings.Contains(name, "endpoint"):
		return "\"https://api.example.com\""
	case strings.Contains(name, "key") || strings.Contains(name, "secret"):
		return "\"your-secret-key-here\""
	case strings.Contains(name, "debug") || strings.Contains(name, "enable"):
		return "true"
	default:
		// Fallback based on type
		switch envVar.Type {
		case "int":
			return "8080"
		case "bool":
			return "true"
		case "url":
			return "\"https://example.com\""
		default:
			return "\"your-value-here\""
		}
	}
}
