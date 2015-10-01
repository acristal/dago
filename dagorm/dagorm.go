// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package dagorm

import "fmt"

// Column represents the type storing column name
type Column string

// ColumnAlias represents the type storing aliased column name
type ColumnAlias Column

// IEntity represents the DB entity contract.
type IEntity interface {
	ID() uint
	SetID(id uint)

	IsDirty() bool
	SetDirty(dirty bool)
}

// IOperations represents the base contract and contains all
// basics database operations.
type IOperations interface {
	// CRUD
	LoadOne(id uint) IEntity
	FindOne(id uint) IEntity
	FindAll() []IEntity

	Create(entity IEntity)
	Update(entity IEntity) IEntity
	Delete(entity IEntity)
	DeleteById(entityID uint)

	// Initialize nested relations and lazy fields
	Initialize(entity IEntity, properties ...string) IEntity
	InitializeList(entities []IEntity, properties ...string) []IEntity
}

// IDao represents the DAO(repository) contract
type IDao interface {
	IOperations

	TableName() string
	FieldNames() []Column
	FieldAliases() []ColumnAlias
	FieldValues() []interface{}

	Query() IQueryBuilder
	ColumnOf(column Column) ColumnAlias
	ColumnsOf(columns ...Column) []ColumnAlias
}

// AbstractDao implements shared features for all DAOs
type AbstractDao struct {
	alias string
}

// ColumnOf ...
func (dao *AbstractDao) ColumnOf(column Column) ColumnAlias {
	return ColumnAlias(fmt.Sprintf("%s.%s", dao.alias, column))
}

// ColumnsOf ...
func (dao *AbstractDao) ColumnsOf(columns ...Column) []ColumnAlias {
	var columnAliases = make([]ColumnAlias, len(columns))

	for idx, column := range columns {
		columnAliases[idx] = dao.ColumnOf(column)
	}
	return columnAliases
}

// IService represents the Service contract
type IService interface {
	IOperations

	DAO() IOperations
}
