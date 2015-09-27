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
)

// Entity represents a struct with the @Entity annotation
type Entity struct {
	TypeSpec *ast.TypeSpec

	Annotations AnnotationList
	Fields      []*Field

	TableName string
}

// NewEntity creates an entity with the default TableName being the entity's name
func NewEntity(ts *ast.TypeSpec) *Entity {
	return &Entity{
		TypeSpec:  ts,
		TableName: ts.Name.Name,
	}
}

// String ...
func (e Entity) String() string {
	return fmt.Sprintf("Entity[%s], table[%s]", e.TypeSpec.Name.Name, e.TableName)
}

// Field represents a column of an entity
type Field struct {
	Field *ast.Field

	Annotations AnnotationList
	IsID        bool
	ColumnName  string

	Relation *Relation // Relation object, if any
}

// NewField creates a field with the default ColumnName being the field's name
func NewField(f *ast.Field) *Field {
	return &Field{
		Field:      f,
		ColumnName: f.Names[0].Name,
	}
}

// String ...
func (e Field) String() string {
	return fmt.Sprintf("Field[%s], column[%s]", e.Field.Names[0].Name, e.ColumnName)
}
