// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package main

import (
	"fmt"

	"github.com/acrisal/dago/annotations"
)

func NewError(message string) error {
	return fmt.Errorf("DagoError: %v", message)
}

func NewTooManyRelationsError(fieldName string) error {
	return fmt.Errorf("DagoError: Field %v cannot have multiple relations", fieldName)
}

func NewInvalidRelationTypeError(relationType annotations.RelationType, fieldName, typeName string) error {
	return fmt.Errorf("DagoError: %v requires field %v to be a %v", annotations.RelationName[relationType], fieldName, typeName)
}
