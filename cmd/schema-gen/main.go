package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"
	"time"

	yaml "github.com/ghodss/yaml"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

var (
	crdPath     string
	crdFiles    string
	packageName string
	schemePath  string
)

func init() {
	flag.StringVar(&crdPath, "crd-path", "", "path to the folder containing the controller-gen generated crds")
	flag.StringVar(&packageName, "package", "schema", "the name of the package the code should reside")
	flag.StringVar(&crdFiles, "crd-files", "*_crd.yaml", "the glob used to find the crd files in the deployment directory")
	flag.StringVar(&schemePath, "schema-file", "pkg/apis/schema/schema.go", "the path to the folder where the schema should be placed")
}

var packageTemplate = template.Must(template.New("").Parse(`
/*
 * Copyright (C) 2019 Rohith Jayawardene <gambol99@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// Code generated by go generate; DO NOT EDIT.
// This file was generated by github.com/appvia/hub-apis/cmd/schema-gen
// {{ .Timestamp }}

package {{ .Package }}

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	yaml "github.com/ghodss/yaml"
)

var (
	// Schema is the openapi schema
	Schema = ` + "`" + `
{{ .Schema -}}
	` + "`" + `
)

func ConvertToJSON() apiextensions.JSON {
	encoded, err := yaml.YAMLToJSON([]byte(Schema))
	if err != nil {
		panic("invalid crd schema")
	}

	return apiextensions.JSON{Raw: encoded}
}
`))

func main() {
	flag.Parse()

	err := func() error {
		filter := fmt.Sprintf("%s/%s", crdPath, crdFiles)
		// @step: scan the directory for crd files
		crds, err := filepath.Glob(filter)
		if err != nil {
			return err
		}

		var definitions []*apiextensions.CustomResourceDefinition

		for _, x := range crds {
			// @step: read in the contents of the file
			content, err := ioutil.ReadFile(x)
			if err != nil {
				return err
			}

			// @step: convert the yaml to json
			converted, err := yaml.YAMLToJSON(content)
			if err != nil {
				return err
			}

			crd := &apiextensions.CustomResourceDefinition{}

			// @step: decode the yaml
			if err := json.NewDecoder(bytes.NewReader(converted)).Decode(crd); err != nil {
				return err
			}

			definitions = append(definitions, crd)
		}

		// @step: convert the into a definition
		swagger := make(map[string]map[string]*apiextensions.JSONSchemaProps, 0)
		swagger["definitions"] = make(map[string]*apiextensions.JSONSchemaProps, 0)

		for _, x := range definitions {
			swagger["definitions"][x.Spec.Names.Kind] = x.Spec.Validation.OpenAPIV3Schema
		}

		encoded, err := yaml.Marshal(swagger)
		if err != nil {
			return err
		}

		// @step: ensure the directory structure
		base := path.Dir(schemePath)
		if base != "" {
			if err := os.MkdirAll(base, os.FileMode(0764)); err != nil {
				return err
			}
		}

		// @step: create the the schema file
		file, err := os.Create(schemePath)
		if err != nil {
			return err
		}

		packageTemplate.Execute(file, struct {
			Timestamp time.Time
			Package   string
			Schema    string
		}{
			Timestamp: time.Now(),
			Package:   packageName,
			Schema:    string(encoded),
		})

		return nil
	}()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] %s", err)
		os.Exit(1)
	}
}
