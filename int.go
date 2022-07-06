package typ

import (
	"encoding/json"
	"strconv"
)

type Int int64

func (t Int) ToUndefinableInt() UndefinableInt {
	return UndefinableInt{NewUndefinable(t)}
}

func (t Int) ToNullableInt() NullableInt {
	return NullableInt{NewNullable(t)}
}

func (t Int) ToInt64() int64 {
	return int64(t)
}

func (t Int) Equals(compT Int) bool {
	return t == compT
}

func (t Int) String() string {
	return strconv.FormatInt(int64(t), 10)
}

func (t *Int) UnmarshalJSON(data []byte) error {
	var temp int64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*t = Int(temp)
	return nil
}

func (t Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(t))
}

type NullableInt struct {
	Nullable[Int]
}

func (t NullableInt) Equals(compT NullableInt) bool {
	return t.Nullable.Equals(compT.Nullable, func(t1 Int, t2 Int) bool {
		return t1.Equals(t2)
	})
}

func (t NullableInt) ToUndefinableNullableInt() UndefinableNullableInt {
	return UndefinableNullableInt{NewUndefinable(t)}
}

type UndefinableInt struct {
	Undefinable[Int]
}

func (t UndefinableInt) ToIntType() (Int, bool) {
	return t.Value()
}

func (t UndefinableInt) String() string {
	if tInt, ok := t.ToIntType(); ok {
		return tInt.String()
	}
	return ""
}

type UndefinableNullableInt struct {
	Undefinable[NullableInt]
}

func (t UndefinableNullableInt) Equals(compT UndefinableNullableInt) bool {
	return t.Undefinable.Equals(compT.Undefinable, func(t1 NullableInt, t2 NullableInt) bool {
		return t1.Equals(t2)
	})
}
