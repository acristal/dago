// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package annotations

// Transient ...
type Transient struct {
}

// IsValidFor ...
func (a *Transient) IsValidFor(dest Type) bool {
	return dest == TypeField
}

// Parse ...
func (a *Transient) Parse(value string) error {
	return nil
}
