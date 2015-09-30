// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package annotations

// AssociationType list all the association type handled by Dago
type RelationType uint8

const (
	OneToOneRelation RelationType = iota
	OneToManyRelation
	ManyToOneRelation
	ManyToManyRelation
)

var OppositeOfRelation = map[RelationType]RelationType{
	OneToOneRelation:   OneToOneRelation,
	OneToManyRelation:  ManyToOneRelation,
	ManyToOneRelation:  OneToManyRelation,
	ManyToManyRelation: ManyToManyRelation,
}

var RelationName = map[RelationType]string{
	OneToOneRelation:   "OneToOne",
	OneToManyRelation:  "OneToMany",
	ManyToOneRelation:  "ManyToOne",
	ManyToManyRelation: "ManyToMany",
}

// RelationType specifies the type of the relation (uni/bi directional)
type DirectionType uint8

const (
	Unidirectional DirectionType = iota
	Bidirectional
)

type RelationContract interface {
	BuildRelation() *Relation
}

// Relation defines a relation between two models
type Relation struct {
	Contract  RelationContract
	Type      RelationType
	Direction DirectionType
}
