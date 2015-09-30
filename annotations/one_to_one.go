// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package annotations

import (
	"strconv"
	"strings"
)

// OneToOne ...
type OneToOne struct {
	Inverse  string
	MappedBy string
	Optional bool
}

func (a *OneToOne) BuildRelation() *Relation {
	return &Relation{
		Contract:  a,
		Type:      OneToOneRelation,
		Direction: Unidirectional,
	}
}

// IsValidFor ...
func (a *OneToOne) IsValidFor(dest Type) bool {
	return dest == TypeField
}

// Parse ...
func (a *OneToOne) Parse(value string) error {
	params := strings.Split(value, ",")
	for _, param := range params {
		kvp := strings.SplitN(param, "=", 2)
		if len(kvp) != 2 {
			return NewParameterArgumentRequiredError(kvp[0])
		}
		switch strings.TrimSpace(kvp[0]) {
		case "mappedBy":
			a.MappedBy = strings.TrimSpace(kvp[1])
		case "inverse":
			a.Inverse = strings.TrimSpace(kvp[1])
		case "optional":
			var err error
			if a.Optional, err = strconv.ParseBool(strings.TrimSpace(kvp[1])); err != nil {
				return err
			}
		default:
			return NewInvalidParameterError(kvp[0])
		}
	}
	return nil
}

// Validate ...
func (a *OneToOne) Validate() error {
	if a.MappedBy != "" && a.Inverse != "" {
		return NewError("Cannot have mappedBy and inverse set at the same time.")
	} else if a.MappedBy == "" && a.Inverse == "" {
		return NewMissingRequiredParameterError("@OneToOne:mappedBy || @OneToOne:inverse")
	}
	return nil
}
