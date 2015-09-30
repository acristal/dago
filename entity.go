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

	"github.com/acrisal/dago/annotations"
	"github.com/golang/glog"
)

// Entity represents a struct with the @Entity annotation
type Entity struct {
	Parent   *Generator
	TypeSpec *ast.TypeSpec
	Name     string

	Annotations AnnotationList
	Fields      []*Field

	TableName string
}

// Field represents a column of an entity
type Field struct {
	Parent *Entity
	Field  *ast.Field
	Name   string

	Annotations AnnotationList
	IsID        bool
	ColumnName  string

	Relation *Relation // Relation object, if any
}

// NewEntity creates an entity with the default TableName being the entity's name
func (g *Generator) NewEntity(ts *ast.TypeSpec) *Entity {
	return &Entity{
		Parent:    g,
		TypeSpec:  ts,
		Name:      ts.Name.Name,
		TableName: ts.Name.Name,
	}
}

// NewField creates a field with the default ColumnName being the field's name
func (e *Entity) NewField(f *ast.Field) *Field {
	return &Field{
		Parent:     e,
		Field:      f,
		Name:       f.Names[0].Name,
		ColumnName: f.Names[0].Name,
	}
}

func (e *Entity) EvalAnnotations() error {
	for _, annotation := range e.Annotations {
		switch t := annotation.(type) {
		case *annotations.Entity:
			// Nothing to do here
		case *annotations.Table:
			e.TableName = t.Name
		default:
			return fmt.Errorf("Invalid annotation %T", t)
		}
	}
	for _, field := range e.Fields {
		if err := field.EvalAnnotations(); err != nil {
			return err
		}
	}
	return nil
}

func (e *Entity) ValidateTypes() error {
	for _, field := range e.Fields {
		if err := field.ValidateTypes(); err != nil {
			return err
		}
	}
	return nil
}

func (f *Field) EvalAnnotations() error {
	for _, annotation := range f.Annotations {
		switch t := annotation.(type) {
		case *annotations.Column:
			f.ColumnName = t.Name
		case *annotations.ID:
			f.IsID = true
		case annotations.RelationContract:
			if f.Relation != nil {
				if f.Relation.Contract == t {
					glog.V(3).Infof("Found relation %v on field %v", f.Relation, f)
					continue
				}
				return NewTooManyRelationsError(f.Field.Names[0].Name)
			}
			f.Relation = NewRelation(f, t)
			if f.Relation == nil {
				return NewError("Couldn't create relation")
			}
		case *annotations.OrderBy:
		default:
			return fmt.Errorf("Invalid annotation %T", t)
		}
	}
	return nil
}

// ValidateTypes check if the types of the field is valid (for example, a OneToMany needs to be a slice)
func (f *Field) ValidateTypes() error {
	if f.Relation == nil {
		return nil
	}
	switch f.Relation.Type {
	case annotations.OneToOneRelation:
		if !ExprIsPointer(f.Relation.FromField.Field.Type) {
			return NewInvalidRelationTypeError(f.Relation.Type, f.Name, "Pointer")
		}
	case annotations.OneToManyRelation:
		if !ExprIsPtrArray(f.Relation.FromField.Field.Type) {
			return NewInvalidRelationTypeError(f.Relation.Type, f.Name, "Array of Pointer")
		}

	case annotations.ManyToOneRelation:
		if !ExprIsPointer(f.Relation.FromField.Field.Type) {
			return NewInvalidRelationTypeError(f.Relation.Type, f.Name, "Pointer")
		}
	case annotations.ManyToManyRelation:
		if !ExprIsPtrArray(f.Relation.FromField.Field.Type) {
			return NewInvalidRelationTypeError(f.Relation.Type, f.Name, "Array of Pointer")
		}
	}
	return nil
}

// String ...
func (e *Entity) String() string {
	return fmt.Sprintf("Entity[%s]", e.Name)
}

// String ...
func (f *Field) String() string {
	return fmt.Sprintf("Field[%s]", f.Name)
}
