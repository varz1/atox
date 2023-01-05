package atox

import (
	"fmt"
	"reflect"
	"strconv"

	"golang.org/x/exp/constraints"
)

// MustN turn into anything from string or panic
func MustN[T constraints.Float | constraints.Integer](s string) T {
	t, err := N[T](s)
	if err != nil {
		panic(err)
	}
	return t
}

// N turn into anything from string
func N[T constraints.Float | constraints.Integer](s string) (T, error) {
	var z T
	rt := reflect.TypeOf(z)
	switch rt.Kind() {
	case reflect.Float32, reflect.Float64:
		t, err := strconv.ParseFloat(s, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		t, err := strconv.ParseInt(s, 10, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		t, err := strconv.ParseUint(s, 10, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	default:
		return z, fmt.Errorf("impossible?")
	}
}

// F turn anything into string
func F(a any) string {
	rt := reflect.TypeOf(a)
	switch rt.Kind() {
	case reflect.Int:
		return strconv.Itoa(a.(int))
	case reflect.Int8:
		return strconv.Itoa(int(a.(int8)))
	case reflect.Int16:
		return strconv.Itoa(int(a.(int16)))
	case reflect.Int32:
		return strconv.Itoa(int(a.(int32)))
	case reflect.Int64:
		return strconv.FormatInt(a.(int64), 10)
	case reflect.Uint:
		return strconv.Itoa(int(a.(uint)))
	case reflect.Uint8:
		return strconv.Itoa(int(a.(uint8)))
	case reflect.Uint16:
		return strconv.Itoa(int(a.(uint16)))
	case reflect.Uint32:
		return strconv.Itoa(int(a.(uint32)))
	case reflect.Uint64:
		return strconv.Itoa(int(a.(uint64)))
	case reflect.Uintptr:
		return strconv.Itoa(int(a.(uint8)))
	default:
		return fmt.Sprintf("%v", a)
	}
}
