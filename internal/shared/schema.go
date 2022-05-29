package shared

import (
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
	compiled, err := jsonschema.Compile(so.src)
	if err != nil {
		return err
	}

	so.compiled = compiled
	return nil
}

func (so *SchemaOptions) Validate(data interface{}) error {
	return so.compiled.Validate(data)
}
