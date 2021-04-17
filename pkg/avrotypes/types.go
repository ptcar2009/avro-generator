package avrotypes

import (
	"bytes"
	"go/types"
	"text/template"

	"github.com/fatih/structtag"
)

func ASTNodeToAvro(name string, n types.Type) string {
	switch t := n.(type) {
	case *types.Struct:
		return structToAvro(name, t)
	case *types.Basic:
		return basicTypeToAvro(t)
	case *types.Named:
		return ASTNodeToAvro(t.Obj().Name(), t.Underlying())
	case *types.Pointer:
		return pointerToAvro(t)
	case *types.Array:
		return arrayToAvro(t)
	case *types.Slice:
		return sliceToAvro(t)
	}
	return ""
}

type structFields struct {
	Name   string
	Fields []struct {
		Name, Type string
	}
}

func pointerToAvro(p *types.Pointer) string {
	k := struct{ Type string }{
		ASTNodeToAvro("", p.Elem()),
	}
	tmp, err := template.New("pointer.tmpl").ParseGlob("*.tmpl")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBufferString("")
	err = tmp.Execute(buf, k)
	if err != nil {
		panic(err)
	}
	return buf.String()

}
func arrayToAvro(p *types.Array) string {
	k := struct{ Type string }{
		ASTNodeToAvro("", p.Elem()),
	}
	tmp, err := template.New("array.tmpl").ParseGlob("*.tmpl")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBufferString("")
	err = tmp.Execute(buf, k)
	if err != nil {
		panic(err)
	}
	return buf.String()

}

func sliceToAvro(p *types.Slice) string {
	k := struct{ Type string }{
		ASTNodeToAvro("", p.Elem()),
	}
	tmp, err := template.New("array.tmpl").ParseGlob("*.tmpl")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBufferString("")
	err = tmp.Execute(buf, k)
	if err != nil {
		panic(err)
	}
	return buf.String()

}
func structToAvro(name string, s *types.Struct) string {
	var results structFields
	results.Name = name
	for i := 0; i < s.NumFields(); i++ {
		field := s.Field(i)
		tag := s.Tag(i)
		tags, _ := structtag.Parse(tag)
		t, err := tags.Get("json")
		if err != nil {
			tag = field.Name()
		} else {
			tag = t.Name
		}
		n := ASTNodeToAvro(tag, field.Type())
		if n == "" {
			continue
		}
		results.Fields = append(results.Fields, struct {
			Name string
			Type string
		}{
			Name: tag,
			Type: n,
		})
	}
	tmp, err := template.New("struct.tmpl").ParseGlob("*.tmpl")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBufferString("")
	err = tmp.Execute(buf, results)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

// func astStarToAvro(s *ast.StarExpr) (string, error) {

// }

func basicTypeToAvro(i *types.Basic) string {
	switch i.String() {
	case "int":
		return "\"int\""
	case "float32":
		return "\"float\""
	case "float64":
		return "\"double\""
	case "byte", "char", "rune", "string":
		return "\"string\""
	case "bool":
		return "\"boolean\""
	default:
		panic("invalid type")
	}

}
