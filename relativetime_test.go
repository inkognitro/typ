package typ

import (
	"encoding/json"
	"testing"
)

func TestParseRelativeTimeCorrectly(t *testing.T) {
	strToParse := "\"11:39\""
	bytesToParse := []byte(strToParse)
	entry := RelativeTime("")
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

func TestCannotParseInvalidRelativeTimeFormat(t *testing.T) {
	strToParse := "\"2021-11-21T11:39\""
	bytesToParse := []byte(strToParse)
	entry := RelativeTime("")
	err := json.Unmarshal(bytesToParse, &entry)
	if err == nil {
		t.Errorf("parsed foo as valid json")
	}
}