package cmd

import (
	"fmt"
	"log"
	"os"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"strings"
)

type EnvRules struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default,omitempty"`
}

type Config struct {
	Variables []EnvRules `yaml:"variables"`
}

var outputFile string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate a template schema.yaml file",
	Long: `Generate a schema template file with example environment variables.

	The schema defines rules for your environment variables:
	• Which variables are required
	• Expected types (string, int, bool, url)
	• Default values for optional variables
	• Descriptions for documentation

	Edit the generated file to match your application's needs.`,

		Example: `  # Create schema.yaml in current directory
	envcheck create

	# Create with custom filename
	envcheck create --output my-schema.yaml

	# Create and edit immediately
	envcheck create && vim schema.yaml`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return CreateSchemaTemplate(outputFile)
	},
}

func CreateCmd() *cobra.Command {
	return createCmd
}

func init() {
	CreateCmd().Flags().StringVarP(
		&outputFile,
		"output",
		"o",
		"schema.yaml",
		"output file path",
	)
}

func CreateSchemaTemplate(filename string) error {
	cfg := Config{
		Variables: []EnvRules{
			{
				Name:        "APP_DATABASE_URL",
				Required:    true,
				Type:        "string",
				Description: "PostgreSQL connection string",
			},
			{
				Name:        "APP_PORT",
				Required:    false,
				Type:        "int",
				Default:     "8080",
				Description: "Server port",
			},
			{
				Name:        "APP_DEBUG",
				Required:    false,
				Type:        "bool",
				Default:     "false",
				Description: "Enable debug mode",
			},
		},
	}
	_ = cfg

	for _, vars := range cfg.Variables {
		if !strings.HasPrefix(vars.Name, "APP") {
			fmt.Fprintf(os.Stderr, "%s: %v\n", "APP Should be a prefix", fmt.Errorf("unknown variable: %s", vars.Name))
			os.Exit(1)
		}
	}

	// marshal the struct into a yaml format
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		panic(err)
	}
	fileName := "schema.yaml"
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Successfully generated YAML file: %s\n", fileName)
	fmt.Println("--- Generated YAML ---")
	fmt.Println(string(data), "data")
	return nil
}
