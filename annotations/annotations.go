// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package annotations

import (
	"go/ast"
	"regexp"

	"github.com/golang/glog"
)

var annotationRegexp = regexp.MustCompile(`\/\/\s*@(\w+)(\(((\s*\w+\s*=\s*\w+\s*)(\s*,\s*\w+\s*=\s*.+)*)\))?`)

// Type ...
type Type uint8

const (
	// TypeStruct ...
	TypeStruct Type = iota
	// TypeField ...
	TypeField
)

// Annotation ...
type Annotation interface {
	Parse(string) error
	IsValidFor(Type) bool
}

// Create is a factory-like function used to create annotation object from string
func Create(name string) Annotation {
	switch name {
	case "Column":
		return &Column{}
	case "Entity":
		return &Entity{}
	case "Id":
		return &ID{}
	case "ManyToMany":
		return &ManyToMany{}
	case "ManyToOne":
		return &ManyToOne{}
	case "OneToMany":
		return &OneToMany{}
	case "OneToOne":
		return &OneToOne{}
	case "OrderBy":
		return &OrderBy{}
	case "Table":
		return &Table{}
	case "Transient":
		return &Transient{}
	default:
		glog.Warningf("Unknown annotation: %s", name)
		return nil
	}
}

// Parse ...
func Parse(comments *ast.CommentGroup, annType Type) (annotations []Annotation) {
	if comments == nil {
		return
	}

	for _, comment := range comments.List {
		matches := annotationRegexp.FindStringSubmatch(comment.Text)
		if matches == nil {
			// glog.Infof("Comment is not an annotation: %v", comment.Text)
			continue
		}
		// glog.Infof("Match [%v] as annotation: %v[%v]", comment.Text, matches[1], matches[3])
		annotation := Create(matches[1])
		if annotation == nil {
			continue
		}
		if !annotation.IsValidFor(annType) {
			// glog.Warningf(" [%s] not valid for type: %d", matches[1], annType)
			continue
		}
		if err := annotation.Parse(matches[3]); err != nil {
			glog.Warningf("Error while parsing annotation: %v", err)
			continue
		}
		annotations = append(annotations, annotation)
	}
	return
}
