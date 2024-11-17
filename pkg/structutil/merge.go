package structutil

import "reflect"

func MergeStruct[T any](a T, b map[string]any) T {
	org := reflect.ValueOf(a)

	var r reflect.Value
	if org.Kind() == reflect.Ptr {
		r = reflect.New(org.Elem().Type()).Elem()
		r.Set(org.Elem())
	} else {
		r = reflect.New(org.Type()).Elem()
		r.Set(org)
	}

	for k, v := range b {
		field := r.FieldByName(k)
		if !field.IsValid() {
			panic("no such field:" + k)
		}

		if !field.CanSet() {
			panic("cannot set field:" + k)
		}

		vv := reflect.ValueOf(v)
		if field.Type() != vv.Type() {
			panic("field type mismatch:" + k)
		}

		field.Set(vv)
	}

	if org.Kind() == reflect.Ptr {
		return r.Addr().Interface().(T)
	}

	return r.Interface().(T)
}
