package template

import (
	"context"
	"docker-init/internal/types"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/qri-io/jsonschema"
	"sigs.k8s.io/yaml"
)

func validateInput(schemaPath string, userInput interface{}) error {
	ctx := context.Background()

	schemaBytes, err := os.ReadFile(schemaPath)
	schemaJson, err := yaml.YAMLToJSON(schemaBytes)

	rs := &jsonschema.Schema{}
	err = json.Unmarshal(schemaJson, rs)
	if err != nil {
		return err
	}

	var templateInput = []byte(`{ "binary": "docker-init" }`)
	validationErrors, err := rs.ValidateBytes(ctx, templateInput)
	if err != nil {
		return err
	}

	if len(validationErrors) > 0 {
		fmt.Println("The following validation errors were discovered while attempting to generate this template:")
		for _, validationError := range validationErrors {
			fmt.Println(validationError.Error())
		}
		return fmt.Errorf("the provided user input did not pass this template's schema\n")
	}

	return nil
}

func Generate(target types.Target) error {

	templateDir, err := getTemplateDir(target.TargetType)
	if err != nil {
		return err
	}

	schemaPath := filepath.Join(templateDir, "schema.yaml")
	err := validateInput(schemaPath, input)
	if err != nil {
		return err
	}

	templateFilesDir := filepath.Join(templateDir, "template")
	err = fs.WalkDir(os.DirFS(templateFilesDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// noop errors
			return nil
		}

		fmt.Println(path)
		return nil
	})

	return nil
}

func getTemplateDir(targetType types.TargetType) (string, error) {
	switch targetType {
	case types.Go:
		return "internal/template/gomod", nil
	}

	err := fmt.Errorf("no generator found for target type %s\n", targetType.ToString())
	return "", err
}
