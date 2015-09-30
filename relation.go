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
	"github.com/golang/glog"
)

// Relation defines a relation between two models
type Relation struct {
	*annotations.Relation

	FromEntity *Entity
	FromField  *Field

	ToEntity *Entity
	Opposite *Relation // null if Unidirectional
}

func (r *Relation) String() string {
	var direction = "===>"
	if r.Direction == annotations.Bidirectional {
		direction = "<==>"
	}
	var toField *Field
	if r.Opposite != nil {
		toField = r.Opposite.FromField
	}
	return fmt.Sprintf("Relation[%s]: %v:%v %s %v:%v", annotations.RelationName[r.Type], r.FromEntity, r.FromField, direction, r.ToEntity, toField)
}

func NewRelation(f *Field, rc annotations.RelationContract) *Relation {
	fieldName, err := TypeNameOfExpr(f.Field.Type)
	if err != nil {
		glog.Infof("%v", err)
		return nil
	}
	entity := f.Parent.Parent.FindEntityByName(fieldName)
	if entity == nil {
		glog.Errorf("Relation destination: %v: is not an entity", fieldName)
		return nil
	}
	destinationField, drc := FindInverse(rc, f, entity)
	newRelation := &Relation{rc.BuildRelation(), f.Parent, f, entity, nil}
	if destinationField != nil {
		newRelation.Direction = annotations.Bidirectional
		newRelation.Opposite = &Relation{drc.BuildRelation(), entity, destinationField, f.Parent, newRelation}
		destinationField.Relation = newRelation.Opposite
	}
	glog.Infof("New relation: %v", newRelation)
	return newRelation
}

func FindInverse(rc annotations.RelationContract, origin *Field, destination *Entity) (*Field, annotations.RelationContract) {
	oppositeRelationType := annotations.OppositeOfRelation[rc.BuildRelation().Type]
	for _, field := range destination.Fields {
		for _, annotation := range field.Annotations {
			relation, ok := annotation.(annotations.RelationContract)
			if !ok {
				continue
			}
			if relation.BuildRelation().Type != oppositeRelationType {
				continue
			}
			switch t := relation.(type) {
			case *annotations.OneToOne:
				if t.Inverse == origin.Name {
					return field, t
				}
			case *annotations.OneToMany:
				if t.MappedBy == origin.Name {
					return field, t
				}
			case *annotations.ManyToOne:
				// We should be sure that rc is a OneToMany because the opposite type of the ManyToOne is OneToMany
				// FIXME(ROMAIN): Maybe check the result of the cast?
				if field.Name == rc.(*annotations.OneToMany).MappedBy {
					return field, t
				}
			case *annotations.ManyToMany:
				// We should be sure that rc is a ManyToMany because the opposite type of the ManyToMany is ManyToMany
				// FIXME(ROMAIN): Maybe check the result of the cast?
				if t.MappedBy == origin.Name || field.Name == rc.(*annotations.ManyToMany).MappedBy {
					return field, t
				}
			default:
				glog.Warningf("Unknown relation type: %T", oppositeRelationType)
			}
		}
	}
	return nil, nil
}
