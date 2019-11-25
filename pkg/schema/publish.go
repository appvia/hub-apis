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
	"context"

	clustersv1 "github.com/appvia/hub-apis/pkg/apis/clusters/v1"
	configv1 "github.com/appvia/hub-apis/pkg/apis/config/v1"
	orgv1 "github.com/appvia/hub-apis/pkg/apis/org/v1"
	rbacv1 "github.com/appvia/hub-apis/pkg/apis/rbac/v1"
	storev1 "github.com/appvia/hub-apis/pkg/apis/store/v1"

	aerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewClient returns a client for the hub schema
func NewClient(cfg *rest.Config) (client.Client, error) {
	s := scheme.Scheme

	clustersv1.AddToScheme(s)
	configv1.AddToScheme(s)
	orgv1.AddToScheme(s)
	rbacv1.AddToScheme(s)
	storev1.AddToScheme(s)

	return client.New(cfg, client.Options{Scheme: s})
}

// PublishAll is responsible for publishing the classes and plans into the cluster
func PublishAll(ctx context.Context, cc client.Client, class configv1.Class, plans []configv1.Plan) error {
	if err := PublishClass(ctx, cc, class); err != nil {
		return err
	}

	return PublishPlans(ctx, cc, plans)
}

// PublishClass is responsible for publishing the class into the cluster
func PublishClass(ctx context.Context, cc client.Client, class configv1.Class) error {
	// @step: check if a class it already published
	current := &configv1.Class{}

	return func() error {
		reference := types.NamespacedName{
			Namespace: "hub",
			Name:      class.Name,
		}

		if err := cc.Get(ctx, reference, current); err != nil {
			if !aerrors.IsNotFound(err) {
				return err
			}

			return cc.Create(ctx, &class)
		}
		class.SetGeneration(current.GetGeneration())
		class.SetResourceVersion(current.GetResourceVersion())

		return cc.Update(ctx, &class)
	}()
}

// PublishPlans is responisble for pushing the plans into the cluster
func PublishPlans(ctx context.Context, cc client.Client, plans []configv1.Plan) error {
	for _, x := range plans {
		if err := func() error {
			current := &configv1.Plan{}

			reference := types.NamespacedName{
				Namespace: "hub",
				Name:      x.Name,
			}

			if err := cc.Get(ctx, reference, current); err != nil {
				if !aerrors.IsNotFound(err) {
					return err
				}

				return cc.Create(ctx, &x)
			}
			x.SetGeneration(current.GetGeneration())
			x.SetResourceVersion(current.GetResourceVersion())

			return cc.Update(ctx, &x)
		}(); err != nil {
			return err
		}
	}

	return nil
}
