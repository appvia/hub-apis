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

	configv1 "github.com/appvia/hub-apis/pkg/apis/config/v1"

	aerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// PublishAll is responsible for publishing the classes and plans into the cluster
func PublishAll(ctx context.Context, c client.Client, class *configv1.Class, plans []*configv1.Plan) error {
	if err := PublishClass(ctx, c, class); err != nil {
		return err
	}

	return PublishPlans(ctx, c, plans)
}

// PublishClass is responsible for publishing the class into the cluster
func PublishClass(ctx context.Context, c client.Client, class *configv1.Class) error {
	err := c.Get(ctx, class, &configv1.Class{})
	if err != nil {
		if aerrors.IsNotFound(err) {
			if err := c.Create(ctx, class); err != nil {
				return err
			}
		}
	}

	return c.Update(ctx, class)
}

// PublishPlans is responisble for pushing the plans into the cluster
func PublishPlans(ctx context.Context, c client.Client, plans []*configv1.Plan) error {
	for _, x := range plans {
		reference := types.NamespacedName{Name: x.Name}

		if err := c.Get(ctx, reference, x); err != nil {
			if aerrors.IsNotFound(err) {
				if err := c.Create(ctx, x); err != nil {
					return err
				}
			}
		} else {
			if err := c.Update(ctx, x); err != nil {
				return err
			}
		}
	}

	return nil
}
