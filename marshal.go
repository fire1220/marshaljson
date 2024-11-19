package marshaljson

import (
	"encoding/json"
	"reflect"
	"time"
)

func verifyField(fieldType reflect.StructField, fieldVal reflect.Value, tabName string) (tabT, bool) {
	tm, ok := tabMap[tabName]
	if !ok {
		return tm, false
	}
	if _, ok := fieldType.Tag.Lookup(tabName); !ok {
		return tm, false
	}
	if tm.restrain != "" && fieldType.Type.String() != tm.restrain {
		return tm, false
	}
	if tabName == tabDateTime && fieldVal.IsZero() {
		return tm, true
	}
	if tabName == tabDefault && !fieldVal.IsZero() {
		return tm, false
	}
	if tabName != tabDefault && fieldVal.IsZero() {
		return tm, false
	}
	return tm, true
}

func MarshalFormat(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	typ := ref.Type()
	newField := make([]reflect.StructField, 0, ref.NumField())
	isNeedNewStruct := false
	for i := 0; i < ref.NumField(); i++ {
		field := typ.Field(i)
		fieldType := field.Type
		for _, tabName := range tabList {
			tm, ok := verifyField(field, ref.Field(i), tabName)
			if !ok {
				continue
			}
			fieldType = tm.refTypOf
			isNeedNewStruct = true
			break
		}
		newField = append(newField, reflect.StructField{
			Name: field.Name,
			Type: fieldType,
			Tag:  field.Tag,
		})
	}
	if !isNeedNewStruct {
		return json.Marshal(p)
	}

	newStruct := reflect.New(reflect.StructOf(newField)).Elem()
	for i := 0; i < newStruct.NumField(); i++ {
		oldFieldVal := ref.Field(i)
		oldFileType := typ.Field(i)
		var newFieldVal reflect.Value
		newFieldVal = oldFieldVal
		for _, tabName := range tabList {
			tm, ok := verifyField(oldFileType, oldFieldVal, tabName)
			if !ok {
				continue
			}
			if tm.fun == nil {
				continue
			}
			newVal, ok := tm.fun.typeConv(oldFieldVal, oldFileType)
			if ok {
				newFieldVal = newVal
			}
			break
		}
		newStruct.Field(i).Set(newFieldVal)
	}
	return json.Marshal(newStruct.Interface())
}

func MarshalFormat1(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	typ := ref.Type()

	newField := make([]reflect.StructField, 0, ref.NumField())
	dateTimeReflectType := reflect.TypeOf(dateTime{})
	for i := 0; i < ref.NumField(); i++ {
		field := typ.Field(i)
		fieldType := field.Type
		if field.Type.String() == "time.Time" {
			fieldType = dateTimeReflectType
		}
		newField = append(newField, reflect.StructField{
			Name: field.Name,
			Type: fieldType,
			Tag:  field.Tag,
		})
	}

	newStruct := reflect.New(reflect.StructOf(newField)).Elem()
	for i := 0; i < newStruct.NumField(); i++ {
		oldField := ref.Field(i)
		oldFieldType := typ.Field(i)
		if oldField.Type().String() != "time.Time" {
			newStruct.Field(i).Set(oldField)
			continue
		}
		if v, ok := oldField.Interface().(time.Time); ok {
			newStruct.Field(i).Set(reflect.ValueOf(dateTime{t: v, tag: oldFieldType.Tag}))
		}
	}

	return json.Marshal(newStruct.Interface())
}
