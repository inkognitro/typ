package typ

import (
	"encoding/json"
	"fmt"
)

type Boolean bool

func (t Boolean) ToBool() bool {
	return bool(t)
}

func (t Boolean) ToUndefinableBoolean() UndefinableBoolean {
	return UndefinableBoolean{NewUndefinable(t)}
}

func (t Boolean) String() string {
	return fmt.Sprintf("%v", t.ToBool())
}

func (t Boolean) IsSame(compT Boolean) bool {
	return t.ToBool() == compT.ToBool()
}

func NewBoolean(content bool) Boolean {
	return Boolean(content)
}

func (t *Boolean) UnmarshalJSON(data []byte) error {
	var temp bool
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*t = Boolean(temp)
	return nil
}

func (t Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(t))
}

type UndefinableBoolean struct {
	Undefinable[Boolean]
}

func (t UndefinableBoolean) ToBool() (bool, bool) {
	if val, ok := t.Value(); ok {
		return bool(val), true
	}
	return false, false
}
