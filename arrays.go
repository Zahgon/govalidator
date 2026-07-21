package govalidator

type Iterator func(interface{}, int)

type ResultIterator func(interface{}, int) interface{}

type ConditionIterator func(interface{}, int) bool

type ReduceIterator func(interface{}, interface{}) interface{}

func Some(array []interface{}, iterator ConditionIterator) bool {
	_ = "STUB: not implemented"
	return false
}

func Every(array []interface{}, iterator ConditionIterator) bool {
	_ = "STUB: not implemented"
	return false
}

func Reduce(array []interface{}, iterator ReduceIterator, initialValue interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func Each(array []interface{}, iterator Iterator) { _ = "STUB: not implemented"; return }

func Map(array []interface{}, iterator ResultIterator) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func Find(array []interface{}, iterator ConditionIterator) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func Filter(array []interface{}, iterator ConditionIterator) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func Count(array []interface{}, iterator ConditionIterator) int {
	_ = "STUB: not implemented"
	return 0
}
