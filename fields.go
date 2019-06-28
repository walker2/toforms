package toforms

import (
	"reflect"
)

type field struct {
	Label       string
	Name        string
	Type        string
	Placeholder string
	Value       interface{}
}

func valueOf(v interface{}) reflect.Value {
	var rv reflect.Value
	switch value := v.(type) {
	case reflect.Value:
		rv = value
	default:
		rv = reflect.ValueOf(v)
	}

	if rv.Kind() == reflect.Ptr {
		// If pointer is nil we should create new copy of underlying type
		if rv.IsNil() {
			rv = reflect.New(rv.Type().Elem())
		}
		// Dereference pointer
		rv = rv.Elem()
	}
	return rv
}

func fields(strct interface{}) []field {
	rv := valueOf(strct)
	if rv.Kind() != reflect.Struct {
		panic("Form: only structs are supported")
	}

	t := rv.Type()
	var ret []field

	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		rvf := valueOf(rv.Field(i))
		if !rvf.CanInterface() {
			continue
		}
		if rvf.Kind() == reflect.Struct {
			nestedFields := fields(rvf.Interface())
			for i, nf := range nestedFields {
				nestedFields[i].Name = tf.Name + "." + nf.Name
			}
			ret = append(ret, nestedFields...)
			continue
		}
		f := field{
			Label:       tf.Name,
			Name:        tf.Name,
			Type:        "text",
			Placeholder: tf.Name,
			Value:       rvf.Interface(),
		}
		ret = append(ret, f)
	}
	return ret
}

// func HTML(strct interface{}, tpl *template.Template) template.HTML {
// }
