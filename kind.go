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
	Method
)

var kindNames = map[Kind]string{
	String:     "String",
	Bool:       "Bool",
	Byte:       "Byte",
	Int:        "Int",
	Int8:       "Int8",
	Int16:      "Int16",
	Int32:      "Int32",
	Int64:      "Int64",
	Uint:       "Uint",
	Uint8:      "Uint8",
	Uint16:     "Uint16",
	Uint32:     "Uint32",
	Uint64:     "Uint64",
	Float32:    "Float32",
	Float64:    "Float64",
	Complex64:  "Complex64",
	Complex128: "Complex128",
	Reference:  "Reference",
	Pointer:    "Pointer",
	Slice:      "Slice",
	Map:        "Map",
	Interface:  "Interface",
	Chan:       "Chan",
	Method:     "Method",
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
}

func (kind Kind) String() string {
	return kindNames[kind]
}

func (kind Kind) IsFlat() bool {
	return flatKinds[kind]
}
