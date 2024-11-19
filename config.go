package marshaljson

import "reflect"

const (
	tabDefault  = "default"
	tabDateTime = "datetime"
)

type tabT struct {
	refTypOf reflect.Type
	restrain string
	fun      typeConvT
}

var (
	tabList = []string{tabDefault, tabDateTime}
	tabMap  = map[string]tabT{
		tabDefault: {
			refTypOf: reflect.TypeOf(defaultT{}),
			restrain: "",
			fun:      defaultT{},
		},
		tabDateTime: {
			refTypOf: reflect.TypeOf(dateTime{}),
			restrain: "time.Time",
			fun:      dateTime{},
		},
	}
)
