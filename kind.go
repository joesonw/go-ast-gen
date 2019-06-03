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

func (kind Kind) String() string {
	return kindNames[kind]
}
