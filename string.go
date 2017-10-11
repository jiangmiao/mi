package mi

type String string

func (s String) Underscorize() String {
	return Bytes(s).Underscorize().String()
}

func (s String) Camelize(args ...interface{}) String {
	return Bytes(s).Camelize(args...).String()
}
