package tmpl

import _ "embed"

//go:embed array.tmpl
var ArrayTemplate string

//go:embed struct.tmpl
var StructTemplate string

//go:embed pointer.tmpl
var PointerTemplate string
