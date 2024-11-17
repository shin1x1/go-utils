package assert

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func EqualStructFields(t *testing.T, actual any, expected map[string]any) {
	r := reflect.ValueOf(actual)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}

	if r.Kind() != reflect.Struct {
		t.Fatalf("expected struct, but got: %s", r.Kind())
	}

	for k, v := range expected {
		field := r.FieldByName(k)
		if !field.IsValid() {
			t.Fatalf("field not found in struct: %s", k)
		}

		fieldType := field.Type()
		valueType := reflect.TypeOf(v)
		if field.Type() != reflect.TypeOf(v) {
			t.Fatalf("field %s expected %s, but got: %s", k, fieldType, valueType)
		}

		var actualValue any
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			actualValue = r.FieldByName(k).String()
		case reflect.Int:
			actualValue = r.FieldByName(k).Int()
		default:
			t.Fatalf("unknown type for field %s: %s", k, reflect.TypeOf(v).Kind())
		}

		assert.EqualValues(t, v, actualValue)
	}
}
