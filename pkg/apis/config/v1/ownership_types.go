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

package v1

// Ownership indicates the ownership of a resource
type Ownership struct {
	// Group is the apigroup the resource lives under
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	APIGroup string `json:"apiGroup"`
	// Kind is the name of the resource under the group
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`
	// Version is the apigroup version of the kind
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Version string `json:"version"`
	// Namespace is the location of the object
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Namespace string `json:"namespace"`
	// Name is name of the resource
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}
