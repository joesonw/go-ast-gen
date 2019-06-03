package astgen

type Kind uint

const (
	Invalid Kind = iota
	String
	Bool
	Byte
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	Complex64
	Complex128
	Reference
	Pointer
	Slice
	Map
	Interface
	Chan
	Imported
)

var kindNames = map[Kind]string{
	String:     "string",
	Bool:       "bool",
	Byte:       "byte",
	Int:        "int",
	Int8:       "int8",
	Int16:      "int16",
	Int32:      "int32",
	Int64:      "int64",
	Uint:       "uint",
	Uint8:      "uint8",
	Uint16:     "uint16",
	Uint32:     "uint32",
	Uint64:     "uint64",
	Float32:    "float32",
	Float64:    "float64",
	Complex64:  "complex64",
	Complex128: "complex128",
	Reference:  "Reference",
	Pointer:    "Pointer",
	Slice:      "Slice",
	Map:        "Map",
	Interface:  "interface{}",
	Chan:       "Chan",
	Imported:   "Imported",
}

var flatKinds = map[Kind]bool{
	String:     true,
	Bool:       true,
	Byte:       true,
	Int:        true,
	Int8:       true,
	Int16:      true,
	Int32:      true,
	Int64:      true,
	Uint:       true,
	Uint8:      true,
	Uint16:     true,
	Uint32:     true,
	Uint64:     true,
	Float32:    true,
	Float64:    true,
	Complex64:  true,
	Complex128: true,
	Interface:  true,
}

func (kind Kind) String() string {
	return kindNames[kind]
}

func (kind Kind) IsFlat() bool {
	return flatKinds[kind]
}
