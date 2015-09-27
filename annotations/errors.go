// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package annotations

import (
	"fmt"
	"strings"
)

func NewError(message string) error {
	return fmt.Errorf("AnnotationError: %v", message)
}

func NewInvalidParameterError(parameterName string) error {
	return fmt.Errorf("AnnotationError: Invalid parameter [%v]", strings.TrimSpace(parameterName))
}

func NewParameterArgumentRequiredError(parameterName string) error {
	return fmt.Errorf("AnnotationError: Parameter [%v] requires an argument", strings.TrimSpace(parameterName))
}
