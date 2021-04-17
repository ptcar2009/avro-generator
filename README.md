# Avro schema from struct generator

This is an avro-schema from struct generator. It ignores functions, channels and other non supported kinds, and creates a avroschema using names from the json tags for your structure.

## Usage

```bash
avro-generator [-o OUTPUT| -p PACKAGE] <typename>
```

This can be used together with gogen to generate avroschemas for structs. This usage requres go1.16

```go
//go:generate go run github.com/ptcar2009/avro-generator ExportedStruct
type ExportedStruct struct {
    Field int `json:"field"`
    OtherField struct {
        Enum int
    }
}
```

## Field nomeclature

Fields are named by the `json` tag value of the struct, or falls back to the struct name.

## Type translation

| `go` native type | `avro` type                                       |
| ---------------- | ------------------------------------------------- |
| `struct`         | `record`                                          |
| `int`            | `int`                                             |
| `float32`        | `float`                                           |
| `float64`        | `double`                                          |
| pointer types    | type union between the underlying type and `null` |
| `array`          | `array`                                           |
| `slice`          | `array`                                           |
| `char`           | `string`                                          |
| `rune`           | `string`                                          |
