// Copyright 2015 Romain Gros
//
// Dago, Relational Persistence for Golang
//
// License: GNU Lesser General Public License (LGPL), version 2.1 or later.
// See the LICENSE file in the root directory or <http://www.gnu.org/licenses/lgpl-2.1.html>.
//

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/acrisal/dago/annotations"
	"github.com/golang/glog"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	for _, path := range os.Args[1:] {
		processFile(path)
	}
}

func processFile(file string) {
	glog.Info("Processing ", file)

	var g Generator

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		glog.Fatalf("Could not parse file: %s", err)
	}

	g.parse(fset, f)
	src := g.format()

	outputName := fmt.Sprintf("%s_dago.go", strings.TrimSuffix(file, filepath.Ext(file)))

	err = ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		glog.Fatalf("Writing output failed: %s", err)
	}
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf         bytes.Buffer // Accumulated output.
	packageName string
	Entities    []*Entity // List of entity we found
}

func (g *Generator) parse(fs *token.FileSet, f *ast.File) {
	g.packageName = f.Name.Name
	glog.Infof("Package: %v", g.packageName)

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			// We only care about TYPE declaration
			continue
		}
		for _, spec := range genDecl.Specs {
			// We look for the structs
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			entity := NewEntity(typeSpec)
			// glog.Infof("Found struct %v", entity.TableName)
			entity.Annotations = annotations.Parse(genDecl.Doc, annotations.TypeStruct)
			// glog.Infof("Found %d annotations", len(entity.Annotations))
			if entity.Annotations.AnyOf(&annotations.Entity{}) == false {
				continue
			}
			glog.Infof("%v: %d annotations", entity, len(entity.Annotations))
			for _, field := range structType.Fields.List {
				if field.Names == nil {
					continue
				}
				entityField := NewField(field)
				if field.Doc != nil {
					// glog.Infof("Found comment [%v] for field %v", field.Doc.Text(), field.Names[0].Name)
					entityField.Annotations = annotations.Parse(field.Doc, annotations.TypeField)
					if entityField.Annotations.AnyOf(&annotations.Transient{}) {
						// A @Transient field means that we ignore this field
						continue
					}
				}
				entity.Fields = append(entity.Fields, entityField)
				glog.Infof("%v: %d annotations", entityField, len(entityField.Annotations))
			}
			g.Entities = append(g.Entities, entity)
		}
	}
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		glog.Errorf("internal error: invalid Go generated: %s", err)
		glog.Errorf("compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}
