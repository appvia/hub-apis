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

package publish

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// NewExtentionsAPIClient returns an extensions api client
func NewExtentionsAPIClient(cfg *rest.Config) (client.Interface, error) {
	return client.NewForConfig(cfg)
}

// ApplyCustomResourceDefinitions s responsible for applying a collection of CRDs
func ApplyCustomResourceDefinitions(c client.Interface, list []*apiextensions.CustomResourceDefinition) error {
	for _, crd := range list {
		if err := ApplyCustomResourceDefinition(c, crd); err != nil {
			return err
		}
	}

	return nil
}

// ApplyCustomResourceDefinition is responsible for applying the CRD to the cluster
func ApplyCustomResourceDefinition(c client.Interface, crd *apiextensions.CustomResourceDefinition) error {
	fmt.Printf("CRD: %s\n", spew.Sdump(crd))

	// @step: retrieve the current if there
	current, err := c.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crd.Name, metav1.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}

		_, err := c.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)

		return err
	}

	crd.SetGeneration(current.GetGeneration())
	crd.SetResourceVersion(current.GetResourceVersion())

	_, err = c.ApiextensionsV1beta1().CustomResourceDefinitions().Update(crd)

	return err
}
