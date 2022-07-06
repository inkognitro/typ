package typ

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

const defaultFloatPrecision = 10

func NewFloat(value float64) Float {
	return Float{value: value}
}

type Float struct {
	precision                 int
	precisionWasSetExplicitly bool
	value                     float64
}

func (t Float) ToUndefinableFloat() UndefinableFloat {
	return UndefinableFloat{NewUndefinable(t)}
}

func (t Float) ToNullableFloat() NullableFloat {
	return NullableFloat{NewNullable(t)}
}

func (t Float) ToFloat64() float64 {
	return t.value
}

func (t Float) SetPrecision(precision int) Float {
	if precision < 0 || precision > 10 {
		panic(fmt.Errorf("float precision requirements not met by precision of %d: 0 <= precision <= 10", precision))
	}
	t.precisionWasSetExplicitly = true
	t.precision = precision
	return t
}

func (t Float) Precision() int {
	if t.precisionWasSetExplicitly {
		return t.precision
	}
	return defaultFloatPrecision
}

func (t Float) Equals(compT Float) bool {
	precision := t.Precision()
	compTPrecision := compT.Precision()
	if compTPrecision > precision {
		precision = compTPrecision
	}
	float64EqualityThreshold := 1 / math.Pow(10, float64(precision+1))
	return math.Abs(t.ToFloat64()-compT.ToFloat64()) <= float64EqualityThreshold
}

func (t Float) String() string {
	return strconv.FormatFloat(t.value, 'f', t.Precision(), 64)
}

func (t *Float) UnmarshalJSON(data []byte) error {
	var tempVal float64
	if err := json.Unmarshal(data, &tempVal); err != nil {
		return err
	}
	t.value = tempVal
	return nil
}

func (t Float) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.*f", t.Precision(), t.value)
	return []byte(s), nil
}

type NullableFloat struct {
	Nullable[Float]
}

func (t NullableFloat) ToFloatType() (Float, bool) {
	return t.Value()
}

func (t NullableFloat) Equals(compT NullableFloat) bool {
	return t.Nullable.Equals(compT.Nullable, func(t1 Float, t2 Float) bool {
		return t1.Equals(t2)
	})
}

func (t NullableFloat) ToUndefinableNullableFloat() UndefinableNullableFloat {
	return UndefinableNullableFloat{NewUndefinable(t)}
}

type UndefinableFloat struct {
	Undefinable[Float]
}

func (t UndefinableFloat) String() string {
	if tFloat, ok := t.Value(); ok {
		return tFloat.String()
	}
	return ""
}

type UndefinableNullableFloat struct {
	Undefinable[NullableFloat]
}

func (t UndefinableNullableFloat) ToNullableFloat() (NullableFloat, bool) {
	return t.Value()
}

func (t UndefinableNullableFloat) Equals(compT UndefinableNullableFloat) bool {
	return t.Undefinable.Equals(compT.Undefinable, func(t1 NullableFloat, t2 NullableFloat) bool {
		return t1.Equals(t2)
	})
}
