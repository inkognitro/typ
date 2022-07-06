package typ

import (
	"encoding/json"
	"testing"
)

func TestParseMicroTimeCorrectly(t *testing.T) {
	strToParse := "\"2021-11-21T11:39:12.039Z\""
	bytesToParse := []byte(strToParse)
	entry := MicroTime{}
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

func TestParseMicroTimeCorrectlyWithoutMilliseconds(t *testing.T) {
	strToParse := "\"2021-11-21T11:39:12Z\""
	bytesToParse := []byte(strToParse)
	entry := MicroTime{}
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

func TestParseMicroTimeCorrectlyFromAnotherTimeZone(t *testing.T) {
	strToParse := "\"2021-11-21T13:39:12+02:00\""
	strToParseInUtc := "\"2021-11-21T11:39:12Z\""
	bytesToParse := []byte(strToParse)
	entry := MicroTime{}
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
	if actualStr != strToParseInUtc {
		t.Errorf(`did receive %s instead of %s`, actualStr, strToParseInUtc)
	}
}

func TestCannotParseInvalidMicroTimeFormat(t *testing.T) {
	strToParse := "\"foo\""
	bytesToParse := []byte(strToParse)
	entry := MicroTime{}
	err := json.Unmarshal(bytesToParse, &entry)
	if err == nil {
		t.Errorf("parsed foo as valid json")
	}
}