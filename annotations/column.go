// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package annotations

import "strings"

// Column ...
type Column struct {
	Name string
}

// IsValidFor ...
func (a *Column) IsValidFor(dest Type) bool {
	return dest == TypeField
}

// Parse ...
func (a *Column) Parse(value string) error {
	params := strings.Split(value, ",")
	for _, param := range params {
		kvp := strings.SplitN(param, "=", 2)
		if len(kvp) != 2 {
			return NewParameterArgumentRequiredError(kvp[0])
		}
		switch strings.TrimSpace(kvp[0]) {
		case "name":
			a.Name = strings.TrimSpace(kvp[1])
		default:
			return NewInvalidParameterError(kvp[0])
		}
	}
	return nil
}
