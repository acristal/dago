// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package main

import "go/ast"

// Entity represents a struct with the @Entity annotation
type Entity struct {
	TypeSpec *ast.TypeSpec

	Annotations AnnotationList
	Fields      []*Field

	TableName string
}

func NewEntity(ts *ast.TypeSpec) *Entity {
	return &Entity{
		TypeSpec:  ts,
		TableName: ts.Name.Name,
	}
}

// Field represents a column of an entity
type Field struct {
	Field *ast.Field

	Annotations AnnotationList
	IsID        bool
	ColumnName  string

	Relation *Relation // Relation object, if any
}

func NewField(f *ast.Field) *Field {
	return &Field{
		Field:      f,
		ColumnName: f.Names[0].Name,
	}
}

// AssociationType list all the association type handled by Dago
type AssociationType uint8

const (
	OneToOne RelationType = iota
	OneToMany
	ManyToOne
	ManyToMany
)

// RelationType specifies the type of the relation (uni/bi directionnal)
type RelationType uint8

const (
	Unidirectionnal RelationType = iota
	Bidirectionnal
)

type Relation struct {
	Association AssociationType
	Type        RelationType
}
