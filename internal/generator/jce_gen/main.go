package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"reflect"
	"strconv"
)

type JceField struct {
	Name string
	Type ast.Expr
	Id   int
}

type JceStruct struct {
	Name   string
	Fields []JceField
}

type Generator struct {
	JceStructs []JceStruct
}

func main() {
	file := flag.String("file", "structs.txt", "file to parse")
	output := flag.String("o", "", "output file")
	flag.Parse()
	f, err := parser.ParseFile(token.NewFileSet(), *file, nil, parser.ParseComments)
	assert(err == nil, err)
	g := Generator{}
	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.FuncDecl:
			return false
		case *ast.TypeSpec:
			x, ok := n.Type.(*ast.StructType)
			if !ok {
				return false
			}
			if x.Fields != nil {
				var jce JceStruct
				for _, f := range x.Fields.List {
					if f.Tag == nil {
						continue
					}
					unquoted, _ := strconv.Unquote(f.Tag.Value)
					tag, ok := reflect.StructTag(unquoted).Lookup("jceId")
					if !ok {
						continue
					}
					id, _ := strconv.Atoi(tag)
					if len(f.Names) != 1 {
						panic("unexpected name count")
					}
					jce.Fields = append(jce.Fields, JceField{
						Name: f.Names[0].Name,
						Type: f.Type,
						Id:   id,
					})
				}
				jce.Name = n.Name.Name
				g.JceStructs = append(g.JceStructs, jce)
			}
		}
		return true
	})
	buf := new(bytes.Buffer)
	assert(err == nil, err)
	g.Generate(buf)
	formated, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("%s\n", buf.Bytes())
		panic(err)
	}
	_ = os.WriteFile(*output, formated, 0644)
}

func (g Generator) Generate(w io.Writer) {
	io.WriteString(w, "// Code generated by internal/generator/jce_gen; DO NOT EDIT.\n\npackage jce\n\n")
	for _, jce := range g.JceStructs {
		if jce.Name == "" || len(jce.Fields) == 0 {
			continue
		}

		var generate func(typ ast.Expr, i int)
		generate = func(typ ast.Expr, i int) {
			field := jce.Fields[i]
			switch typ := typ.(type) {
			case *ast.Ident:
				switch typ.Name {
				case "uint8", "uint16", "uint32", "uint64":
					typename := []byte(typ.Name)[1:]
					casted := fmt.Sprintf("%s(pkt.%s)", typename, field.Name)
					typename[0] ^= ' '
					fmt.Fprintf(w, "w.Write%s(%s, %d)\n", typename, casted, field.Id)

				case "int8", "int16", "int32", "int64",
					"byte", "string":
					typename := []byte(typ.Name)
					typename[0] ^= ' '
					fmt.Fprintf(w, "w.Write%s(pkt.%s, %d)\n", typename, field.Name, field.Id)

				case "int":
					fmt.Fprintf(w, "w.WriteInt64(int64(pkt.%s), %d)\n", field.Name, field.Id)

				default:
					fmt.Fprintf(w, `{ // write %s tag=%d
					w.writeHead(10, %d)
					w.buf.Write(pkt.%s.ToBytes())
					w.writeHead(11, 0)}`+"\n", field.Name, field.Id, field.Id, field.Name)
				}
			case *ast.StarExpr:
				_ = typ.X.(*ast.Ident) // assert
				generate(typ.X, i)
			case *ast.ArrayType:
				assert(typ.Len == nil, "unexpected array len")
				var method string
				switch typeName(typ) {
				case "[]byte":
					method = "WriteBytes"
				case "[]int64":
					method = "WriteInt64Slice"
				case "[][]byte":
					method = "WriteBytesSlice"

				default:
					fmt.Fprintf(w,
						"\t"+`{ // write pkt.%s tag=%d 
	w.writeHead(9, %d)
	if len(pkt.%s) == 0 {
		w.writeHead(12, 0) // w.WriteInt32(0, 0)
	} else {
		w.WriteInt32(int32(len(pkt.%s)), 0)
		for _, i := range pkt.%s {
			w.writeHead(10, 0)
			w.buf.Write(i.ToBytes())
			w.writeHead(11, 0)
		}
	}}`+"\n", field.Name, field.Id, field.Id, field.Name, field.Name, field.Name)
					return
				}
				assert(method != "", typeName(typ))
				fmt.Fprintf(w, "w.%s(pkt.%s, %d)\n", method, field.Name, field.Id)
			case *ast.MapType:
				var method string
				typName := typeName(typ)
				switch typName {
				case "map[string]string":
					method = "writeMapStrStr"
				case "map[string][]byte":
					method = "writeMapStrBytes"
				case "map[string]map[string][]byte":
					method = "writeMapStrMapStrBytes"
				}
				assert(method != "", typName)
				fmt.Fprintf(w, "w.%s(pkt.%s, %d)\n", method, field.Name, field.Id)
			}
		}

		fmt.Fprintf(w, "func (pkt *%s) ToBytes() []byte {\nw := NewJceWriter()\n", jce.Name)
		for i := range jce.Fields {
			generate(jce.Fields[i].Type, i)
		}
		fmt.Fprintf(w, "return w.Bytes()\n}\n\n")
	}
}

func assert(cond bool, val interface{}) {
	if !cond {
		panic("assertion failed: " + fmt.Sprint(val))
	}
}

func typeName(typ ast.Expr) string {
	switch typ := typ.(type) {
	case *ast.Ident:
		return typ.Name
	case *ast.StarExpr:
		return "*" + typeName(typ.X)
	case *ast.ArrayType:
		if typ.Len != nil {
			panic("unexpected array type")
		}
		return "[]" + typeName(typ.Elt)
	case *ast.MapType:
		return "map[" + typeName(typ.Key) + "]" + typeName(typ.Value)
	}
	panic("unexpected type")
}
