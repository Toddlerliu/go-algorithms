package weightedgraph

import (
	"reflect"
	"errors"
)

type Edge struct {
	a, b   int
	weight interface{} // weight
}

func (e Edge) V() int {
	return e.a
}

func (e Edge) W() int {
	return e.b
}

func (e Edge) Weight() interface{} {
	return e.weight
}

func (e Edge) Other(x int) int {
	if x == e.a || x == e.b {
		if x == e.a {
			return e.b
		}
		return e.a
	}
	return -1
}

func (e Edge) WTLtCompare(x Edge) (error, bool) {
	valueOf := reflect.ValueOf(e.weight)
	valueOf.Kind()
	switch wt := e.weight.(type) {
	case float64:
		if v, ok := x.Weight().(float64); ok {
			return nil, wt < v
		}
		return errors.New("error type"), false
	default:
		return errors.New("error type"), false
	}
}

func (e Edge) WTLeCompare(x Edge) (error, bool) {
	valueOf := reflect.ValueOf(e.weight)
	valueOf.Kind()
	switch wt := e.weight.(type) {
	case float64:
		if v, ok := x.Weight().(float64); ok {
			return nil, wt <= v
		}
		return errors.New("error type"), false
	default:
		return errors.New("error type"), false
	}
}
