package cmd

import (
	"os"
	"testing"
	"fmt"
)

type test struct {
	name string
	data EnvRules
	value interface{}
	wantError bool
}

func TestValidateVar(t *testing.T) {
	tests := []test{
		{
			name: "missing prefix",
			data: EnvRules{
				Name: "DATABASE_URL",
				Description: "PostgreSQL connection string",
				Type: "string",
				Default: "postgresql://user:password@localhost:5432/mydatabase",
			},
			value: "postgresql",
			wantError: false,
		},
		{
			name: "unknown type",
			data: EnvRules{
				Name: "PORT",
				Description: "Application Port",
				Type: "int",
				Default: "3080",
			},
			value: "postgresql",
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if(tt.name) != "" {
				os.Setenv(tt.data.Name, fmt.Sprintf("%v", tt.value))
				defer os.Unsetenv(tt.data.Name)
			}
		err := validateVar(tt.data)
		// Check result
		if (err != nil) != tt.wantError {
			t.Errorf("validateVar() error = %v, wantError %v", err, tt.wantError)
		}
		})
	}

}