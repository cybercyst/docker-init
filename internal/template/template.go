package template

import (
	"context"
	"docker-new/internal/types"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/bclicn/color"
	"github.com/flosch/pongo2/v6"
	"github.com/qri-io/jsonschema"
	"sigs.k8s.io/yaml"
)

func validateInput(schemaPath string, userInput map[string]interface{}) error {
	ctx := context.Background()

	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return err
	}
	schemaJson, err := yaml.YAMLToJSON(schemaBytes)
	if err != nil {
		return err
	}

	rs := &jsonschema.Schema{}
	err = json.Unmarshal(schemaJson, rs)
	if err != nil {
		return err
	}

	userInputBytes, err := json.Marshal(userInput)
	if err != nil {
		return err
	}
	validationErrors, err := rs.ValidateBytes(ctx, userInputBytes)
	if err != nil {
		return err
	}

	if len(validationErrors) > 0 {
		fmt.Println("The following validation errors were discovered while attempting to generate this template:")
		for _, validationError := range validationErrors {
			fmt.Println(validationError.Error())
		}
		return fmt.Errorf("the provided user input did not pass this template's schema")
	}

	return nil
}

func Generate(target types.Target) error {
	templateDir, err := getTemplateDir(target.TargetType)
	if err != nil {
		return err
	}

	schemaPath := filepath.Join(templateDir, "schema.yaml")
	err = validateInput(schemaPath, target.Input)
	if err != nil {
		return err
	}

	templateFilesDir := filepath.Join(templateDir, "template")
	err = fs.WalkDir(os.DirFS(templateFilesDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// noop errors
			return nil
		}

		if d.IsDir() {
			return nil
		}

		templateFilePath := filepath.Join(templateFilesDir, path)
		template := pongo2.Must(pongo2.FromFile(templateFilePath))

		out, err := template.Execute(target.Input)
		if err != nil {
			return err
		}

		outputPath := filepath.Join(target.Path, path)
		err = os.MkdirAll(filepath.Dir(outputPath), 0755)
		if err != nil {
			return err
		}

		err = os.WriteFile(outputPath, []byte(out), 0644)
		if err != nil {
			return err
		}
		fmt.Println(color.Green("CREATE"), path)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func getTemplateDir(targetType types.TargetType) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	rootTemplateDir := filepath.Join(homeDir, ".docker-new", "templates")

	switch targetType {
	case types.Go:
		return filepath.Join(rootTemplateDir, "gomod"), nil
	case types.Angular:
		return filepath.Join(rootTemplateDir, "angular"), nil
	}

	err = fmt.Errorf("no generator found for target type %s", targetType.ToString())
	return "", err
}
