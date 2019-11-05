/*
 * Copyright (C) 2019  Rohith Jayawardene <gambol99@gmail.com>
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

package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"
	"k8s.io/kube-openapi/pkg/common"
)

// Schema is a generate schema
type Schema map[string]map[string]interface{}

// OpenAPIFunc is the schema function
type OpenAPIFunc func(common.ReferenceCallback) map[string]common.OpenAPIDefinition

type schemeReference struct{}

func (s schemeReference) ReferenceCallback(p string) spec.Ref {
	return spec.Ref{
		Ref: jsonreference.MustCreateRef(fmt.Sprintf("#/definitions/%s", convertBasePath(p))),
	}
}

func convertBasePath(fullpath string) string {
	items := strings.Split(path.Base(fullpath), ".")
	name := items[0]
	if len(items) > 0 {
		name = items[1]
	}

	return name
}

// OpenAPISchema is responsible to taking a openapi-gen code and generating the
// schema from it
func OpenAPISchema(sc OpenAPIFunc) (Schema, error) {
	s := make(Schema, 0)
	s["definitions"] = make(map[string]interface{}, 0)

	ref := &schemeReference{}

	for k, v := range sc(ref.ReferenceCallback) {
		s["definitions"][convertBasePath(k)] = v.Schema
	}

	return s, nil
}

// OpenAPISchemaJSON converts the output to a json schema
func OpenAPISchemaJSON(v OpenAPIFunc) ([]byte, error) {
	sv, err := OpenAPISchema(v)
	if err != nil {
		return []byte{}, err
	}
	b := &bytes.Buffer{}

	if err := json.NewEncoder(b).Encode(sv); err != nil {
		return []byte{}, err
	}

	return b.Bytes(), nil
}
