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
	"go/ast"
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

func TypeNameOfExpr(e ast.Expr) (string, error) {
	var ident *ast.Ident

	switch t := e.(type) {
	case *ast.StarExpr:
		var ok bool
		ident, ok = t.X.(*ast.Ident)
		if !ok {
			return "", fmt.Errorf("Unknown field type: %T", e)
		}
	case *ast.Ident:
		ident = t
	case *ast.ArrayType:
		return TypeNameOfExpr(t.Elt)
	default:
		return "", fmt.Errorf("Unknown field type: %T", e)
	}
	return ident.Name, nil
}

func ExprIsPointer(e ast.Expr) bool {
	_, ok := e.(*ast.StarExpr)
	return ok
}

func ExprIsPtrArray(e ast.Expr) bool {
	array, ok := e.(*ast.ArrayType)
	if !ok {
		return false
	}
	return ExprIsPointer(array.Elt)
}

func (g *Generator) FindEntityByName(name string) *Entity {
	for _, entity := range g.Entities {
		if entity.Name == name {
			return entity
		}
	}
	return nil
}

func (e *Entity) FindFieldByName(name string) *Field {
	for _, field := range e.Fields {
		if field.Name == name {
			return field
		}
	}
	return nil
}
