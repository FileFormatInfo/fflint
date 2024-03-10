package shared

import (
	"fmt"
	"os"
	"reflect"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"
)

type SchemaOptions struct {
	src      string
	compiled *jsonschema.Schema
}

func (so *SchemaOptions) AddFlags(theCmd *cobra.Command) {
	theCmd.Flags().StringVar(&so.src, "schema", "", "JSON schema path (or URL)")
}

func (so *SchemaOptions) Prepare() error {
	if Debug {
		fmt.Fprintf(os.Stderr, "DEBUG: trying to compile schema: %s\n", so.src)
	}
	if so.src == "" {
		return nil
	}
	compiled, err := jsonschema.Compile(so.src)
	if err != nil {
		if Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: schema compilation error: %v\n", err)
		}
		return err
	}

	if Debug {
		fmt.Fprintf(os.Stderr, "DEBUG: successfully compiled schema from %s\n", so.src)
	}

	so.compiled = compiled
	return nil
}

func (so *SchemaOptions) Validate(f *FileContext, data any) error {
	if so.compiled == nil {
		return nil
	}
	validationErr := so.compiled.Validate(data)
	// check if err is a validationerror
	if Debug && validationErr != nil {
		switch validationErr.(type) {
		case *jsonschema.ValidationError:
			fmt.Fprintf(os.Stderr, "DEBUG: schema validation error: %#v, data %v", validationErr.(*jsonschema.ValidationError).BasicOutput(), data)
		default:
			fmt.Fprintf(os.Stderr, "DEBUG: schema non-validation error: %#v, (%s)", validationErr, reflect.TypeOf(validationErr))
		}
	}
	f.RecordResult("jsonSchemaValidatation", validationErr == nil, map[string]interface{}{
		"error": fmt.Sprintf("%#v", validationErr),
	})

	return validationErr
}
