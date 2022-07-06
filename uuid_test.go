package typ

import (
	"encoding/json"
	"testing"
)

func TestParseUuidCorrectly(t *testing.T) {
	strToParse := "\"8126d6A7-172f-42b5-bb77-7214ddf5a0fc\""
	bytesToParse := []byte(strToParse)
	entry := Uuid("")
	err := json.Unmarshal(bytesToParse, &entry)
	if err != nil {
		t.Errorf("could not execute json.Unmarshal: %v", err)
		return
	}
	actualBytes, err := entry.MarshalJSON()
	actualStr := string(actualBytes)
	if err != nil {
		panic("could not json marshal entry")
	}
	if actualStr != strToParse {
		t.Errorf(`did receive %s instead of %s`, actualStr, strToParse)
	}
}

func TestCannotParseInvalidUuidFormat(t *testing.T) {
	strToParse := "\"foo\""
	bytesToParse := []byte(strToParse)
	entry := Uuid("")
	err := json.Unmarshal(bytesToParse, &entry)
	if err == nil {
		t.Errorf("parsed foo as valid uuid")
	}
}