package marshaljson

import "reflect"

type typeConvT interface {
	typeConv(reflect.Value, reflect.StructField) (reflect.Value, bool)
}
