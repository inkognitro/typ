package typ

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func NewUuid(value string) (Uuid, error) {
	t := Uuid(strings.ToLower(value))
	if !t.IsValid() {
		return t, fmt.Errorf("\"%s\" is not a valid uuid", value)
	}
	return t, nil
}

func NewUuidOrPanic(value string) Uuid {
	t, err := NewUuid(value)
	if err != nil {
		panic(err)
	}
	return t
}

func NewUuidV4OrPanic() Uuid {
	return NewUuidOrPanic(uuid.New().String())
}

func isValidUuidString(value string) bool {
	_, err := uuid.Parse(strings.ToLower(value))
	return err == nil
}

type Uuid string

func (t Uuid) ToUndefinableUuid() UndefinableUuid {
	return NewUndefinableUuidOrPanic(t)
}

func (t Uuid) ToNullableUuid() NullableUuid {
	return NewNullableUuid(t)
}

func (t Uuid) IsValid() bool {
	return isValidUuidString(string(t))
}

func (t Uuid) PanicWhenInvalid() {
	if !t.IsValid() {
		panic(fmt.Errorf(`"%s" is not a valid uuid string`, t))
	}
}

func (t Uuid) ToLowerCaseUuidString() string {
	if t.IsValid() {
		return strings.ToLower(string(t))
	}
	panic(fmt.Errorf(`"%s" is not a valid uuid string`, t))
}

func (t Uuid) Equals(compT Uuid) bool {
	return t.ToLowerCaseUuidString() == compT.ToLowerCaseUuidString()
}

func (t Uuid) ToUuids() Uuids {
	return Uuids{t}
}

func (t Uuid) String() string {
	return string(t)
}

func (t *Uuid) UnmarshalJSON(data []byte) error {
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if !isValidUuidString(temp) {
		return fmt.Errorf("%s is not a valid uuid string", temp)
	}
	*t = Uuid(temp)
	return nil
}

func (t Uuid) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t Uuid) ToUndefinableNullableUuid() UndefinableNullableUuid {
	return UndefinableNullableUuid{NewUndefinable(t.ToNullableUuid())}
}

func NewNullableUuid(uuid Uuid) NullableUuid {
	uuid.PanicWhenInvalid()
	return NullableUuid{NewNullable(uuid)}
}

type NullableUuid struct {
	Nullable[Uuid]
}

func (t NullableUuid) Equals(compT NullableUuid) bool {
	return t.Nullable.Equals(compT.Nullable, func(t1 Uuid, t2 Uuid) bool {
		return t1.Equals(t2)
	})
}

func (t NullableUuid) ToUndefinableNullableUuid() UndefinableNullableUuid {
	return UndefinableNullableUuid{NewUndefinable(t)}
}

func NewUndefinableUuidOrPanic(uuid Uuid) UndefinableUuid {
	val, err := NewUndefinableUuid(uuid)
	if err != nil {
		panic(err)
	}
	return val
}

func NewUndefinableUuid(uuid Uuid) (UndefinableUuid, error) {
	if !uuid.IsValid() {
		return UndefinableUuid{}, fmt.Errorf(`"%s" is an invalid uuid`, uuid)
	}
	return UndefinableUuid{NewUndefinable[Uuid](uuid)}, nil
}

type UndefinableUuid struct {
	Undefinable[Uuid]
}

func (t UndefinableUuid) Equals(compT UndefinableUuid) bool {
	return t.Undefinable.Equals(compT.Undefinable, func(t1 Uuid, t2 Uuid) bool {
		return t1.Equals(t2)
	})
}

type UndefinableNullableUuid struct {
	Undefinable[NullableUuid]
}

func (t UndefinableNullableUuid) Equals(compT UndefinableNullableUuid) bool {
	return t.Undefinable.Equals(compT.Undefinable, func(t1 NullableUuid, t2 NullableUuid) bool {
		return t1.Equals(t2)
	})
}
