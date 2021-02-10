package demo

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	getIndex() int
}

type Iterable interface {
	Iterator() Iterator
}
