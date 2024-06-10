package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

const addrTemplate = `func (i *%[1]s) Get%[2]s() %[3]v {
	if i != nil && i.%[2]s != nil {
		return %[5]si.%[2]s
	}
	return %[4]s
}
`

const baseTemplate = `func (i *%[1]s) Get%[2]s() %[3]v {
	if i != nil {
		return i.%[2]s
	}
	return %[4]s
}
`

const arrTemplate = `func (i *%[1]s) Get%[2]s() []%[3]v {
	if i != nil {
		return i.%[2]s
	}
	return nil
}
`

// uintptr ignored
var (
	TypeDefaultValueMap = map[string]string{
		"string":     "\"\"",
		"bool":       "false",
		"int":        "0",
		"int8":       "0",
		"int16":      "0",
		"int32":      "0",
		"int64":      "0",
		"uint":       "0",
		"uint8":      "0",
		"uint16":     "0",
		"uint32":     "0",
		"uint64":     "0",
		"byte":       "0",
		"rune":       "0",
		"float32":    "0",
		"float64":    "0",
		"complex64":  "0",
		"complex128": "0",
		"array":      "[]interface{}",
	}
)

func (g *Generator) ptrTypeProcessing(field *ast.Field, structName string) {

	fieldType := field.Type

	star, ok := fieldType.(*ast.StarExpr)

	if ok {
		fieldType = star.X
	} else {
		fmt.Println("this is not a star type. caller should check")
		return
	}

	fieldString := fmt.Sprintf("%s", fieldType)

	deRef := "*"

	defaultVal, ok := TypeDefaultValueMap[fieldString]
	if !ok {
		fieldString = "*" + fieldString
		defaultVal = "nil"
		deRef = ""
	}
	g.Printf(addrTemplate, structName, field.Names[0], fieldString, defaultVal, deRef)
}

func (g *Generator) baseTypeProcessing(field *ast.Field, structName string) {

	fieldType := field.Type

	fieldString := fmt.Sprintf("%s", fieldType)

	defaultVal, ok := TypeDefaultValueMap[fieldString]
	if !ok {
		// loll this should not be the case ah...
		defaultVal = fieldString + "{}"
	}
	g.Printf(baseTemplate, structName, field.Names[0], fieldString, defaultVal)
}

func (g *Generator) arrTypeProcessing(field *ast.Field, structName string) {
	arr, ok := field.Type.(*ast.ArrayType)

	if !ok {
		fmt.Println("this is not a arr type. caller should check")
		return
	}

	elementType := arr.Elt

	_, ok = elementType.(*ast.MapType)
	if ok {
		fmt.Println("map type in array not supported")
		return
	}

	isStar := false
	star, ok := elementType.(*ast.StarExpr)
	if ok {
		elementType = star.X
		isStar = true
	}

	fieldString := fmt.Sprintf("%s", elementType)

	if isStar {
		fieldString = "*" + fieldString
	}
	g.Printf(arrTemplate, structName, field.Names[0], fieldString)
}

func (g *Generator) structCheck(node ast.Node) bool {
	t, ok := node.(*ast.TypeSpec)
	if !ok {
		return true
	}

	if t.Type == nil {
		return true
	}

	structName := t.Name.Name

	x, ok := t.Type.(*ast.StructType)
	if !ok {
		return true
	}

	for _, field := range x.Fields.List {

		_, ok := field.Type.(*ast.MapType)
		if ok {
			fmt.Println("map type not supported")
			continue
		}

		_, ok = field.Type.(*ast.ArrayType)

		if ok {
			g.arrTypeProcessing(field, structName)
			continue
		}

		_, ok = field.Type.(*ast.StarExpr)

		if ok {
			g.ptrTypeProcessing(field, structName)
			continue
		}

		g.baseTypeProcessing(field, structName)
	}

	return true
}

type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package

	trimPrefix  string
	lineComment bool
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) parsePackage() {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
	}

	// only load current directory
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		log.Fatal(err)
	}

	g.addPackage(pkgs[0])
}

func (g *Generator) addPackage(pkg *packages.Package) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &File{
			file:        file,
			pkg:         g.pkg,
			trimPrefix:  g.trimPrefix,
			lineComment: g.lineComment,
		}
	}
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	// typeName string // Name of the constant type.
	// values   []Value // Accumulator for constant values of that type.

	trimPrefix  string
	lineComment bool
}

func (g *Generator) generate() {
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		// file.typeName = typeName
		// file.values = nil
		if file.file != nil {
			ast.Inspect(file.file, g.structCheck)
		}
	}
}

func main() {
	fmt.Println("hello world")
	fmt.Println("generating getters for your structs")

	g := Generator{}

	g.parsePackage()

	g.Printf("package %s\n", g.pkg.name)
	g.Printf("// Code generated by getterGen. DO NOT EDIT\n\n")

	g.generate()

	src := g.format()
	// Write to file.
	outputName := ""
	if outputName == "" {
		baseName := fmt.Sprintf("%s_getters.go", g.pkg.name)
		outputName = filepath.Join(".", strings.ToLower(baseName))
	}
	err := os.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
