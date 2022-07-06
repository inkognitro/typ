package typ

import (
	"encoding/json"
	"strings"
)

type Strings []String

func (t Strings) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToStrings())
}

func (t Strings) EqualsCaseSensitive(compT Strings) bool {
	if len(t) != len(compT) {
		return false
	}
	for index, value := range compT {
		if !t[index].EqualsCaseSensitive(value) {
			return false
		}
	}
	return true
}

func (t *Strings) UnmarshalJSON(data []byte) error {
	var tempStrings []string
	if err := json.Unmarshal(data, &tempStrings); err != nil {
		return err
	}
	tStrings := Strings{}
	for _, tempString := range tempStrings {
		tStr := NewString(tempString)
		tStrings = append(tStrings, tStr)
	}
	*t = tStrings
	return nil
}

func (t Strings) ToUndefinableStrings() UndefinableStrings {
	return UndefinableStrings{NewUndefinable(t)}
}

func (t Strings) Add(content String) Strings {
	newT := make(Strings, len(t)+1)
	for index, val := range t {
		newT[index] = val
	}
	newT[len(t)] = content
	return newT
}

func (t Strings) ToStrings() []string {
	primitiveStrings := make([]string, len(t))
	for index, tStrings := range t {
		primitiveStrings[index] = tStrings.ToString()
	}
	return primitiveStrings
}

func (t Strings) String() string {
	primitiveStrings := t.ToStrings()
	return "[" + strings.Join(primitiveStrings, ",") + "]"
}

func (t Strings) Count() int {
	return len(t)
}

type UndefinableStrings struct {
	Undefinable[Strings]
}
