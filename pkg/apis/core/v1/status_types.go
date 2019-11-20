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

package v1

// Status is the status of a thing
type Status string

const (
	// SuccessStatus is a successful resource
	SuccessStatus Status = "Success"
	// FailureStatus indicates the resource has failed for one or more reasons
	FailureStatus Status = "Failure"
)

// Condition is a reason why something failed
// +k8s:openapi-gen=true
type Condition struct {
	// Message is a human readable message
	Message string `json:"message"`
	// Detail is a actual error which might contain technical reference
	Detail string `json:"detail"`
	// Code is machine readable code of the error type
	Code int `json:"code"`
}
