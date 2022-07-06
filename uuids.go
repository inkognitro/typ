package typ

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Uuids []Uuid

func (t Uuids) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToLowercaseUuidStrings())
}

func (t Uuids) Equals(compT Uuids) bool {
	if len(t) != len(compT) {
		return false
	}
	for index, value := range compT {
		if !t[index].Equals(value) {
			return false
		}
	}
	return true
}

func (t Uuids) Contains(uuidToFind Uuid) bool {
	for _, id := range t {
		if id.Equals(uuidToFind) {
			return true
		}
	}
	return false
}

func (t *Uuids) UnmarshalJSON(data []byte) error {
	var tempStrings []string
	if err := json.Unmarshal(data, &tempStrings); err != nil {
		return err
	}
	uuids := Uuids{}
	for _, tempString := range tempStrings {
		uuid, err := NewUuid(tempString)
		if err != nil {
			return fmt.Errorf("could not create Uuids struct: %w", err)
		}
		uuids = append(uuids, uuid)
	}
	*t = uuids
	return nil
}

func (t Uuids) ToUndefinableUuids() UndefinableUuids {
	return UndefinableUuids{NewUndefinable(t)}
}

func (t Uuids) ToLowercaseUuidStrings() []string {
	uuidStrings := make([]string, len(t))
	for index, uuid := range t {
		uuidStrings[index] = uuid.ToLowerCaseUuidString()
	}
	return uuidStrings
}

func (t Uuids) String() string {
	uuidStrings := t.ToLowercaseUuidStrings()
	return "[" + strings.Join(uuidStrings, ",") + "]"
}

func (t Uuids) Count() int {
	return len(t)
}

type UndefinableUuids struct {
	Undefinable[Uuids]
}
