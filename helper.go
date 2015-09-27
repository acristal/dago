// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package main

import (
	"reflect"

	"github.com/acrisal/dago/annotations"
)

type AnnotationList []annotations.Annotation

// AnyOf returns true if the AnnotationList contains at least one annotation of the type passed as argument
func (l AnnotationList) AnyOf(a annotations.Annotation) bool {
	destType := reflect.TypeOf(a)

	for _, elem := range l {
		if reflect.TypeOf(elem).AssignableTo(destType) {
			return true
		}
	}
	return false
}
