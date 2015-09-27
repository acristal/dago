// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package database

//go:generate dago $GOFILE

// @Entity
// @Table(name = persons)
// @Id
type Person struct {
	// @Id
	// @Column(name = id)
	ID uint

	Name string
	// @Transient
	Age uint16

	// @OneToOne(mappedBy = Person)
	Address *Address
}

// @Entity
type Address struct {
	// @Id
	ID uint

	Street string

	// @OneToOne(inverse = Address)
	Person *Person
}

// This struct is not an entity, because it has no @Entity annotation in the comments
// @Table(name = titi)
type NotAnEntity struct {
}
