package yamlutil

import (
	"encoding/json"
	"reflect"
)

// IsEmpty returns true if it is nil, "", 0, or its lenghth is zero.
func IsEmpty(it interface{}) bool {
	if it == nil {
		return true
	}
	t := reflect.TypeOf(it)
	if t.Kind() == reflect.Slice ||
		t.Kind() == reflect.Array ||
		t.Kind() == reflect.Map ||
		t.Kind() == reflect.Chan ||
		t.Kind() == reflect.String {
		return reflect.ValueOf(it).Len() == 0
	} else if reflect.DeepEqual(it, reflect.Zero(t).Interface()) {
		return true
	}
	return false
}

// RemoveEmptyFromMap recursively removes empty values from a map.
func RemoveEmptyFromMap(it interface{}) interface{} {
	mt, mv := reflect.TypeOf(it), reflect.ValueOf(it)
	if mt.Kind() != reflect.Map || mv.Len() == 0 {
		// If it is not a map, do nothing
		return it
	}

	iter := mv.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value().Interface()
		if IsEmpty(v) {
			mv.SetMapIndex(k, reflect.Value{})
		} else {
			vt, vv := reflect.TypeOf(v), reflect.ValueOf(v)
			if IsEmpty(v) {
				mv.SetMapIndex(k, reflect.Value{})
			}
			if vt.Kind() == reflect.Map {
				RemoveEmptyFromMap(v)
				if vv.Len() == 0 {
					mv.SetMapIndex(k, reflect.Value{})
				}
			} else if vt.Kind() == reflect.Ptr {
				// is there a nil pointer in Go?
				v = vv.Elem().Interface()
				RemoveEmptyFromMap(v)
				if IsEmpty(v) {
					mv.SetMapIndex(k, reflect.Value{})
				}
			}
		}
	}
	return it
}

// RemoveEmptyAny marshals a given e into a string, unmarshal back to a map and then remove empty objects.
func RemoveEmptyAny(e interface{}) (map[string]interface{}, error) {
	x, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	z := map[string]interface{}{}
	err = json.Unmarshal(x, &z)
	if err != nil {
		return nil, err
	}

	RemoveEmptyFromMap(z)
	return z, nil
}

// MarshalJSONOmitEmptyIndented mashals a given value as a JSON omitting all empty values
func MarshalJSONOmitEmptyIndented(e interface{}) ([]byte, error) {
	z, err := RemoveEmptyAny(e)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(z, "", "    ")
}

// MarshalJSONOmitEmpty mashals a given value as a JSON omitting all empty values
func MarshalJSONOmitEmpty(e interface{}) ([]byte, error) {
	z, err := RemoveEmptyAny(e)
	if err != nil {
		return json.Marshal(e)
	}
	return json.Marshal(z)
}

// MustMarshalJSONOmitEmpty is a sugar version of MustMarshalJSONOmitEmpty.
func MustMarshalJSONOmitEmpty(e interface{}) string {
	var v interface{}
	if s, ok := e.(string); ok {
		if err := json.Unmarshal([]byte(s), &v); err != nil {
			e = v
		}
	}

	z, _ := MarshalJSONOmitEmptyIndented(e)
	return string(z)
}
