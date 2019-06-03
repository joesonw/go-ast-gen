package astgen

func (t Type) Map() MapType {
	return MapType(t)
}

type MapType Type

func (mt MapType) Key() Type {
	return mt.types[0]
}

func (mt MapType) Value() Type {
	return mt.types[1]
}
