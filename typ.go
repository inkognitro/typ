package typ

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type UndefinableSupporter interface {
	supportUndefinable()
}

type NullableSupporter interface {
	supportNullable()
}

func getStringFromAny(v any) string {
	if val, ok := v.(fmt.Stringer); ok {
		return fmt.Sprintf("%s", val)
	}
	return fmt.Sprintf("%v", v)
}

func isDefined(v any) bool {
	return v != nil
}

func extractUnderlyingValueFromPossiblePointer(ptrOrValue interface{}) interface{} {
	valueField := reflect.ValueOf(ptrOrValue)
	if valueField.Kind() != reflect.Ptr {
		return valueField.Interface()
	}
	if valueField.IsNil() {
		return nil
	}
	value := valueField.Elem().Interface()
	return extractUnderlyingValueFromPossiblePointer(value)
}

type compareFunc[T any] func(t1 T, t2 T) bool

type Undefinable[T any] struct {
	defined bool
	content T
}

func NewUndefinable[T any](content T) Undefinable[T] {
	return Undefinable[T]{defined: isDefined(content), content: content}
}

func (u Undefinable[T]) IsUndefined() bool {
	return !u.defined
}

func (u Undefinable[T]) supportUndefinable() {}

func (u Undefinable[T]) Equals(typeToCompare Undefinable[T], compare compareFunc[T]) bool {
	t1, okT1 := u.Value()
	t2, okT2 := typeToCompare.Value()
	if okT1 != okT2 {
		return false
	}
	if !okT1 {
		return true
	}
	return compare(t1, t2)
}

func (u *Undefinable[T]) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	u.defined = true
	return json.Unmarshal(data, &u.content)
}

func (u Undefinable[T]) MarshalJSON() ([]byte, error) {
	if !u.defined {
		return nil, nil
	}
	return json.Marshal(u.content)
}

func (u Undefinable[T]) String() string {
	if !u.defined {
		return "undefined"
	}
	return getStringFromAny(u.content)
}

func (u Undefinable[T]) Value() (T, bool) {
	return u.content, u.defined
}

type Nullable[T any] struct {
	defined bool
	content T
}

func NewNullable[T any](content T) Nullable[T] {
	return Nullable[T]{defined: isDefined(content), content: content}
}

func (u Nullable[T]) IsNull() bool {
	return !u.defined
}

func (u Nullable[T]) supportNullable() {}

func (u Nullable[T]) Equals(typeToCompare Nullable[T], compare compareFunc[T]) bool {
	t1, okT1 := u.Value()
	t2, okT2 := typeToCompare.Value()
	if okT1 != okT2 {
		return false
	}
	if !okT1 {
		return true
	}
	return compare(t1, t2)
}

func (u *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	u.defined = true
	return json.Unmarshal(data, &u.content)
}

func (u Nullable[T]) MarshalJSON() ([]byte, error) {
	if !u.defined {
		return []byte("null"), nil
	}
	return json.Marshal(u.content)
}

func (u Nullable[T]) String() string {
	if !u.defined {
		return "NULL"
	}
	return getStringFromAny(u.content)
}

func (u Nullable[T]) Value() (T, bool) {
	return u.content, u.defined
}
