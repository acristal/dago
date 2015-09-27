// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package main

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

// Relation defines a relation between two models
type Relation struct {
	Association AssociationType
	Type        RelationType
}
