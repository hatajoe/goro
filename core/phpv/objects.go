package phpv

type ZClassProp struct {
	VarName   ZString
	Default   Val
	Modifiers ZObjectAttr
}

type ZClassMethod struct {
	Name      ZString
	Modifiers ZObjectAttr
	Method    Callable
}

type ZClassHandlers struct {
	Constructor  *ZClassMethod
	HandleInvoke func(ctx Context, o ZObject, args []Runnable) (*ZVal, error)
}

type ZClass interface {
	GetName() ZString
	InstanceOf(subc ZClass) bool
	BaseName() ZString
	GetStaticProps(ctx Context) (*ZHashTable, error)
	GetMethod(name ZString) (*ZClassMethod, bool)
	Handlers() *ZClassHandlers
	GetParent() ZClass
}

type ZObjectAccess interface {
	ObjectGet(ctx Context, key Val) (*ZVal, error)
	ObjectSet(ctx Context, key Val, value *ZVal) error
}

type ZObject interface {
	ZObjectAccess
	Val

	GetOpaque(c ZClass) interface{}
	SetOpaque(c ZClass, v interface{})
	GetClass() ZClass
	NewIterator() ZIterator
	HashTable() *ZHashTable
}
